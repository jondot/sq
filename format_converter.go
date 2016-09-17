package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

// FormatConverter tbd
type FormatConverter struct {
	expr       *regexp.Regexp
	subs       string
	TimeLayout string
}

// NewFormatConverter tbd
func NewFormatConverter(re string, subs string) *FormatConverter {
	tre := ".*"
	if re != "" {
		tre = re
	}
	tsubs := "$0"
	if subs != "" {
		tsubs = subs
	}
	return &FormatConverter{expr: regexp.MustCompile(tre), subs: tsubs}
}

// Convert tbd
func (f *FormatConverter) Convert(text string) int64 {
	if f.TimeLayout != "" {
		return f.time(text, f.TimeLayout).Unix()
	}

	return f.int(text)
}

// String t
func (f *FormatConverter) String(text string) string {
	return f.expr.ReplaceAllString(text, f.subs)
}

// Int t
func (f *FormatConverter) int(text string) int64 {
	s := f.String(text)
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Cannot parse integer value from [%s]\n", s)
	}
	return int64(res)
}

// Time t
func (f *FormatConverter) time(text string, layout string) time.Time {
	s := f.String(text)
	res, err := time.Parse(layout, s)
	if err != nil {
		fmt.Errorf("Cannot parse a date out of [%s] with layout [%s]]\n", s, layout)
	}
	return res
}
