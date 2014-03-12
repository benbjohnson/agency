package agency

import (
	"bytes"
	"errors"
	"unicode"
	"unicode/utf8"
)

// errEOF is an internal EOF marker.
var errEOF = errors.New("eof")

// Scanner is a user agent tokenizer.
type Scanner struct {
	c         rune
	buf       []byte
	buflen    int
	idx       int
	size      int
	prevstart int
	mobile    bool
	browsers  []*browser
	devices   []*device
	oses      []*os
}

// NewScanner creates a new user agent scanner.
func NewScanner() *Scanner {
	return &Scanner{
		browsers: make([]*browser, maxRank),
		devices:  make([]*device, maxRank),
		oses:     make([]*os, maxRank),
	}
}

// ScanBytes scans a user agent byte slice for device information.
func (s *Scanner) ScanBytes(b []byte) (*UserAgent, error) {
	var ua = new(UserAgent)
	s.buf = b
	s.buflen = len(b)
	s.reset()

	// Iterate over each word in the byte slice.
	for {
		unigram, bigram, err := s.readNgrams()
		if err == errEOF {
			break
		} else if err != nil {
			return nil, err
		}

		// Set the mobile flag.
		if !s.mobile {
			s.matchMobile(unigram)
		}

		s.matchBrowser(unigram, bigram)
		s.matchDevice(unigram, bigram)
		s.matchOS(unigram, bigram)
	}

	// Find browser by rank level.
	for _, browser := range s.browsers {
		if browser != nil {
			ua.Browser.Type = browser.typ
			ua.Browser.Name = browser.name
			break
		}
	}

	// Find device by rank level.
	for _, device := range s.devices {
		if device != nil {
			ua.Device.Type = device.typ
			break
		}
	}

	// Find OS by rank level.
	for _, os := range s.oses {
		if os != nil {
			ua.OS.Name = os.name
			ua.OS.Version = os.version
			break
		}
	}

	// Special mobile cases.
	if s.mobile {
		if ua.Browser.Name == "Firefox" {
			ua.Browser.Type = "Mobile Browser"
			ua.Browser.Name = "Mobile Firefox"
		} else if ua.Browser.Name == "Safari" {
			ua.Browser.Type = "Mobile Browser"
			ua.Browser.Name = "Mobile Safari"
		} else if ua.Browser.Name == "Opera" {
			ua.Browser.Type = "Mobile Browser"
			ua.Browser.Name = "Opera Mobile"
		} else if ua.Browser.Name == "Yandex.Browser" {
			ua.Browser.Type = "Mobile Browser"
			ua.Browser.Name = "Yandex.Browser mobile"
		}
	}

	return ua, nil
}

// Scan scans a user agent string for device information.
func (s *Scanner) Scan(str string) (*UserAgent, error) {
	return s.ScanBytes([]byte(str))
}

// read retrieves the next rune from the string.
func (s *Scanner) read() error {
	if s.idx >= s.buflen {
		return errEOF
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
		if err := s.read(); err == errEOF {
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
		return nil, nil, errEOF
	}

	unigram := s.buf[start:s.idx]
	bigram := s.buf[s.prevstart:s.idx]
	s.prevstart = start
	return unigram, bigram, nil
}

// match checks a unigram against the list of mobile tokens.
func (s *Scanner) matchMobile(unigram []byte) {
	for _, mobile := range mobiles {
		if bytes.Equal(unigram, mobile.token) {
			s.mobile = true
		}
	}
}

// matchBrowser checks a unigram and bigram against the list of browser tokens.
func (s *Scanner) matchBrowser(unigram []byte, bigram []byte) {
	for _, browser := range browsers {
		if bytes.Equal(unigram, browser.token) || bytes.Equal(bigram, browser.token) {
			s.browsers[browser.rank] = browser
		}
	}
}

// matchDevice checks a unigram and bigram against the list of device tokens.
func (s *Scanner) matchDevice(unigram []byte, bigram []byte) {
	for _, device := range devices {
		if bytes.Equal(unigram, device.token) || bytes.Equal(bigram, device.token) {
			s.devices[device.rank] = device
		}
	}
}

// matchOS checks a unigram and bigram against the list of OS tokens.
func (s *Scanner) matchOS(unigram []byte, bigram []byte) {
	for _, os := range oses {
		if bytes.Equal(unigram, os.token) || bytes.Equal(bigram, os.token) {
			s.oses[os.rank] = os
		}
	}
}

// reset re-initializes the state of the scanner.
func (s *Scanner) reset() {
	s.idx = 0
	s.size = 0
	s.prevstart = 0
	s.mobile = false

	for i := range s.browsers {
		s.browsers[i] = nil
	}
}

// ScanBytes extracts properties from a user agent byte slice.
func ScanBytes(b []byte) (*UserAgent, error) {
	return NewScanner().ScanBytes(b)
}

// Scan extracts properties from a user agent string.
func Scan(str string) (*UserAgent, error) {
	return NewScanner().Scan(str)
}
