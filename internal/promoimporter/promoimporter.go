package promoimporter

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/core/logger"
	promocodestore "github.com/arjunksofficial/kart-challenge/internal/entities/promocode/store"
)

const chunkSize = 5_000_000

type PromoImporter struct {
	RedisCli promocodestore.Cache
	Logger   *logger.CustomLogger
}

func New() *PromoImporter {
	return &PromoImporter{
		RedisCli: promocodestore.Get(),
		Logger:   logger.GetLogger(),
	}
}

func (p *PromoImporter) Run() error {
	ctx := context.Background()
	var sortedFiles []string
	for i, src := range config.GetPromoImporterConfig().FileSources {
		p.Logger.Infof("Sorting file source: %s tag: source%d", src, i+1)
		tag := fmt.Sprintf("source%d", i+1)
		sorted, err := externalSortAuto(src, tag)
		if err != nil {
			return fmt.Errorf("failed to sort %s: %v", src, err)
		}
		sortedFiles = append(sortedFiles, sorted)
		p.Logger.Infof("Sorted file source: %s tag: %s", src, tag)
	}
	return p.ValidateAndInsertToDB(ctx, sortedFiles[0], sortedFiles[1], sortedFiles[2])
}

func externalSortAuto(pathOrURL, tag string) (string, error) {
	var reader io.ReadCloser
	var err error

	if strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://") {
		resp, err := http.Get(pathOrURL)
		if err != nil {
			return "", fmt.Errorf("failed to GET url: %v", err)
		}
		reader = resp.Body
	} else {
		reader, err = os.Open(pathOrURL)
		if err != nil {
			return "", fmt.Errorf("failed to open file: %v", err)
		}
	}
	defer reader.Close()

	if strings.HasSuffix(pathOrURL, ".gz") {
		gz, err := gzip.NewReader(reader)
		if err != nil {
			return "", fmt.Errorf("failed to init gzip reader: %v", err)
		}
		defer gz.Close()
		reader = gz
	}

	scanner := bufio.NewScanner(reader)
	chunkFiles := []string{}
	chunk := []string{}
	index := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		chunk = append(chunk, line)
		if len(chunk) >= chunkSize {
			tmpFile := filepath.Join("./.sorted", fmt.Sprintf("%s_chunk_%d.txt", tag, index))
			if err := writeSortedChunk(chunk, tmpFile); err != nil {
				return "", err
			}
			chunkFiles = append(chunkFiles, tmpFile)
			chunk = nil
			index++
		}
	}
	if len(chunk) > 0 {
		tmpFile := filepath.Join("./.sorted", fmt.Sprintf("%s_chunk_%d.txt", tag, index))
		if err := writeSortedChunk(chunk, tmpFile); err != nil {
			return "", err
		}
		chunkFiles = append(chunkFiles, tmpFile)
	}

	final := filepath.Join("./.sorted", fmt.Sprintf("%s_sorted.txt", tag))
	out, _ := os.Create(final)
	defer out.Close()
	writer := bufio.NewWriter(out)
	mergeSortedFiles(chunkFiles, writer)
	writer.Flush()

	for _, f := range chunkFiles {
		_ = os.Remove(f)
	}
	return final, nil
}

func writeSortedChunk(lines []string, filePath string) error {
	sort.Strings(lines)
	out, _ := os.Create(filePath)
	defer out.Close()
	writer := bufio.NewWriter(out)
	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}
	return writer.Flush()
}

type scannerState struct {
	scanner *bufio.Scanner
	current string
	eof     bool
}

func newScanner(path string) *scannerState {
	f, _ := os.Open(path)
	s := bufio.NewScanner(f)
	ss := &scannerState{scanner: s}
	ss.advance()
	return ss
}

func (s *scannerState) advance() {
	if s.scanner.Scan() {
		s.current = s.scanner.Text()
	} else {
		s.eof = true
		s.current = ""
	}
}

func mergeSortedFiles(files []string, writer *bufio.Writer) {
	scanners := make([]*scannerState, len(files))
	for i, f := range files {
		scanners[i] = newScanner(f)
	}

	for {
		min := ""
		minIndex := -1
		for i, s := range scanners {
			if !s.eof && (min == "" || s.current < min) {
				min = s.current
				minIndex = i
			}
		}
		if minIndex == -1 {
			break // All files exhausted
		}

		count := 0
		for _, s := range scanners {
			if !s.eof && s.current == min {
				count++
				s.advance()
			}
		}
		if count >= 2 {
			_, _ = writer.WriteString(min + "\n")
		}
	}
}
