package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

// ExifConverter t
type ExifConverter struct {
}

// NewExifConverter t
func NewExifConverter() *ExifConverter {
	return &ExifConverter{}
}

// Convert t
func (e *ExifConverter) String(text string) string {
	return fmt.Sprintf("%v", e.Convert(text))
}

// Convert t
func (e *ExifConverter) Convert(text string) uint64 {
	exif.RegisterParsers(mknote.All...)
	f, err := os.Open(text)
	defer f.Close()
	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	tm, _ := x.DateTime()
	return uint64(tm.Unix())
}
