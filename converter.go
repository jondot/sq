package main

import (
	"log"
	"regexp"
)

/*
fgroup --glob . --exifdate
fgroup --glob . --filedate
fgroup --glob . --format "(.*)/DSC_(.*).jpg --> $1$2"
fgroup --glob . --format "(.*)/DSC_(.*).jpg --> $1-$2" --timestamp "DD-MM"
*/

// Converter t
type Converter interface {
	Convert(string) int64
	String(string) string
}

// ResolveConverter t
func ResolveConverter(format string, tsLayout string, isExif bool, isFiledate bool) Converter {
	if isExif {
		return NewExifConverter()
	}
	if isFiledate {
		return NewFiledateConverter()
	}
	if format == "" && tsLayout == "" {
		// if user didn't pick anything, use filedate
		return NewFiledateConverter()
	}

	re := ""
	subs := ""
	if format != "" {
		res := regexp.MustCompile("\\s*-->\\s*").Split(format, -1)
		if len(res) != 2 {
			log.Fatalf("Cannot parse format string. Did you try '[regex] --> [capture groups]'?\n")
		}
		re = res[0]
		subs = res[1]
	}
	fc := NewFormatConverter(re, subs)
	if tsLayout != "" {
		fc.TimeLayout = tsLayout
	}

	return fc
}
