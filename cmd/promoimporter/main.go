package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/arjunksofficial/kart-challenge/internal/config"
)

const chunkSize = 5_000_000

func main() {
	log.Println("🚀 Starting Promo Importer...")

	config.LoadPromoImporterConfig()

	var sortedFiles []string
	for i, src := range config.PromoImporterCfg.FileSources {
		sorted, err := externalSortAuto(src, fmt.Sprintf("source%d", i+1))
		if err != nil {
			log.Fatalf("Failed to sort %s: %v", src, err)
		}
		sortedFiles = append(sortedFiles, sorted)
	}
	log.Println("✅ All files sorted successfully.")
	err := mergeAndWriteValid(sortedFiles[0], sortedFiles[1], sortedFiles[2], "valid_codes.txt")
	if err != nil {
		log.Fatalf("Merge failed: %v", err)
	}

	fmt.Println("✅ Done. Valid codes written to valid_codes.txt")
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

	// Gzip detection
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
			tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("%s_chunk_%d.txt", tag, index))
			if err := writeSortedChunk(chunk, tmpFile); err != nil {
				return "", err
			}
			chunkFiles = append(chunkFiles, tmpFile)
			chunk = nil
			index++
		}
	}
	if len(chunk) > 0 {
		tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("%s_chunk_%d.txt", tag, index))
		if err := writeSortedChunk(chunk, tmpFile); err != nil {
			return "", err
		}
		chunkFiles = append(chunkFiles, tmpFile)
	}

	// Merge sorted chunks
	final := filepath.Join(os.TempDir(), fmt.Sprintf("%s_sorted.txt", tag))
	out, _ := os.Create(final)
	defer out.Close()
	writer := bufio.NewWriter(out)
	mergeSortedFiles(chunkFiles, writer)
	writer.Flush()

	// Cleanup
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

func mergeAndWriteValid(f1, f2, f3, outputPath string) error {
	s1 := newScanner(f1)
	s2 := newScanner(f2)
	s3 := newScanner(f3)

	outFile, _ := os.Create(outputPath)
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)

	for !s1.eof || !s2.eof || !s3.eof {
		min := minNonEmpty(s1.current, s2.current, s3.current)
		count := 0
		if s1.current == min {
			count++
			s1.advance()
		}
		if s2.current == min {
			count++
			s2.advance()
		}
		if s3.current == min {
			count++
			s3.advance()
		}
		if count >= 2 {
			_, _ = writer.WriteString(min + "\n")
		}
	}
	writer.Flush()
	return nil
}

func minNonEmpty(a, b, c string) string {
	min := ""
	for _, v := range []string{a, b, c} {
		if v == "" {
			continue
		}
		if min == "" || v < min {
			min = v
		}
	}
	return min
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
