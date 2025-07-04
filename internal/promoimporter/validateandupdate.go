package promoimporter

import (
	"bufio"
	"context"
	"os"
)

type UniqueScannerState struct {
	scanner     *bufio.Scanner
	current     string
	eof         bool
	lastScanned string
}

func newUniqueScanner(path string) *UniqueScannerState {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	ss := &UniqueScannerState{scanner: s}
	ss.advance()
	return ss
}

func (s *UniqueScannerState) advance() {
	for s.scanner.Scan() {
		line := s.scanner.Text()
		if line != s.lastScanned {
			s.current = line
			s.lastScanned = line
			return
		}
	}
	s.eof = true
	s.current = ""
}

func (p *PromoImporter) ValidateAndInsertToDB(ctx context.Context, f1, f2, f3 string) error {
	s1 := newUniqueScanner(f1)
	s2 := newUniqueScanner(f2)
	s3 := newUniqueScanner(f3)
	totalValidCodes := 0
	defer func() {
		if err := s1.scanner.Err(); err != nil {
			p.Logger.Errorf("Error reading file 1: %v", err)
		}
		if err := s2.scanner.Err(); err != nil {
			p.Logger.Errorf("Error reading file 2: %v", err)
		}
		if err := s3.scanner.Err(); err != nil {
			p.Logger.Errorf("Error reading file 3: %v", err)
		}
	}()
	for !s1.eof || !s2.eof || !s3.eof {
		min := minNonEmpty(s1.current, s2.current, s3.current)
		presentIn := 0
		if s1.current == min {
			presentIn++
			s1.advance()
		}
		if s2.current == min {
			presentIn++
			s2.advance()
		}
		if s3.current == min {
			presentIn++
			s3.advance()
		}
		if presentIn >= 2 {
			totalValidCodes++
			p.RedisCli.SAdd(ctx, "valid_promo_codes", min)
		}
	}
	p.Logger.Infof("Total valid promo codes inserted: %d", totalValidCodes)
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
