package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// GroupCommandRunner tbd
type GroupCommandRunner struct {
	Command     string
	OutputDir   string
	GroupPrefix string
}

// VisitHeader tbd
func (g *GroupCommandRunner) VisitHeader(group interface{}) {
	groupName := fmt.Sprintf("%s/%s%s", g.OutputDir, g.GroupPrefix, GroupPathToName(group.(string)))
	cmd := strings.Replace(g.Command, "{{GROUP}}", groupName, -1)
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Fatalf("Cannot run format for group [%s]: %v", groupName, err)
	}
	log.Printf("Command [%s] result for [%s]: %s", cmd, groupName, string(out))
}

// VisitNode tbd
func (g *GroupCommandRunner) VisitNode(i int, sz int, header interface{}, file interface{}) {
}
