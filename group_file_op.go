package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	cp "github.com/cleversoap/go-cp"
)

// Op tbd
type Op int

const (
	//OpCopy tbd
	OpCopy Op = iota
	//OpMove tbd
	OpMove
)

// GroupFileOp tbd
type GroupFileOp struct {
	GroupPrefix    string
	OutputDir      string
	ShouldSequence bool
	FileOp         Op
}

// VisitHeader tbd
func (g *GroupFileOp) VisitHeader(group interface{}) {
	groupName := fmt.Sprintf("%s/%s%s", g.OutputDir, g.GroupPrefix, GroupPathToName(group.(string)))
	err := os.MkdirAll(groupName, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating a group: %v", err)
	}
	log.Printf("created: %s", groupName)
}

// VisitNode tbd
func (g *GroupFileOp) VisitNode(i int, sz int, header interface{}, file interface{}) {
	fi := file.(string)
	hd := GroupPathToName(header.(string))
	groupName := fmt.Sprintf("%s/%s%s", g.OutputDir, g.GroupPrefix, hd)
	targetName := fmt.Sprintf("%s/%s", groupName, fi)
	if g.ShouldSequence {
		targetName = fmt.Sprintf("%s/%d%s", groupName, i, filepath.Ext(fi))
	}

	if g.FileOp == OpCopy {
		err := cp.Copy(fi, targetName)
		if err != nil {
			log.Fatalf("Cannot copy: %v", err)
		}
	} else {
		err := os.Rename(fi, targetName)
		if err != nil {
			log.Fatalf("Cannot move: %v", err)
		}
	}
	log.Printf("ok: %s -> %s", fi, targetName)
}
