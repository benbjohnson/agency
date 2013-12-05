package agency

import (
	"bytes"
	"errors"
	"unicode"
	"unicode/utf8"
)

var eof = errors.New("eof")

// Scanner is a user agent tokenizer.
type Scanner struct {
	c      rune
	buf    []byte
	buflen int
	idx    int
	size   int
	prevstart int
}

// NewScanner creates a new user agent scanner.
func NewScanner() *Scanner {
	return &Scanner{}
}

// Scan scans a user agent string for category, type and browser.
func (s *Scanner) ScanBytes(b []byte) (*UserAgent, error) {
	var ua = new(UserAgent)
	s.buf = b
	s.buflen = len(b)
	s.idx = 0

	// Iterate over each word in the byte slice.
	for {
		unigram, bigram, err := s.readNgrams()
		if err == eof {
			break
		} else if err != nil {
			return nil, err
		}

		if ua.Type == "" {
			ua.Type = s.match(types, unigram, bigram)
		}
		if ua.Category == "" {
			ua.Category = s.match(categories, unigram, bigram)
		}
		if ua.Browser == "" {
			ua.Browser = s.match(browsers, unigram, bigram)
		}
	}

	// Default to desktop if nothing else was found.
	if ua.Category == "" {
		ua.Category = "Desktop"
	}

	return ua, nil
}

// Scan scans a user agent string for category, type and browser.
func (s *Scanner) Scan(str string) (*UserAgent, error) {
	return s.ScanBytes([]byte(str))
}

// read retrieves the next rune from the string.
func (s *Scanner) read() error {
	if s.idx >= s.buflen {
		return eof
	}

	// Read a single byte and then determine if utf8 decoding is needed.
	b := s.buf[s.idx]
	if b < utf8.RuneSelf {
		s.c = rune(b)
		s.size = 1
	} else {
		s.c, s.size = utf8.DecodeRune(s.buf[s.idx:])
	}
	s.idx += s.size
	return nil
}

// unread moves back one rune. Only works once.
func (s *Scanner) unread() {
	s.idx -= s.size
	s.size = 0
}

// readWord reads a word and previous bigram from the string.
func (s *Scanner) readNgrams() ([]byte, []byte, error) {
	var index int
	start := s.idx
	for {
		if err := s.read(); err == eof {
			break
		}

		// Only read in letters, numbers and some punctuation.
		if unicode.IsLetter(s.c) || unicode.IsDigit(s.c) || s.c == '-' || s.c == '.' {
			index++
		} else if index == 0 {
			// This section skips over initial non-word characters.
			start = s.idx
		} else {
			s.unread()
			break
		}
	}

	// If nothing was read then it's EOF.
	if s.idx == start {
		return nil, nil, eof
	}

	unigram := s.buf[start:s.idx]
	bigram := s.buf[s.prevstart:s.idx]
	s.prevstart = start
	return unigram, bigram, nil
}

// match checks a unigram and bigram against a list of values.
func (s *Scanner) match(m map[string][][]byte, unigram []byte, bigram []byte) string {
	for key, values := range m {
		for _, value := range values {
			if bytes.Equal(unigram, value) || bytes.Equal(bigram, value) {
				return key
			}
		}
	}
	return ""
}

// Scan extracts properties from a user agent byte slice.
func ScanBytes(b []byte) (*UserAgent, error) {
	return NewScanner().ScanBytes(b)
}

// ScanString extracts properties from a user agent string.
func Scan(str string) (*UserAgent, error) {
	return NewScanner().Scan(str)
}
