package main

import (
	"fmt"
	"log"
	"os"
)

// FiledateConverter t
type FiledateConverter struct {
}

// NewFiledateConverter t
func NewFiledateConverter() *FiledateConverter {
	return &FiledateConverter{}
}

// String t
func (e *FiledateConverter) String(text string) string {
	return fmt.Sprintf("%v", e.Convert(text))
}

// Convert t
func (e *FiledateConverter) Convert(text string) uint64 {
	fi, err := os.Stat(text)
	if err != nil {
		log.Fatalf("Cannot stat file [%s]: %v\n", text, err)
	}
	return uint64(fi.ModTime().Unix())
}
