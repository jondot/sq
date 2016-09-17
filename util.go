package main

import "strings"

// GroupPathToName tbd
func GroupPathToName(g string) string {
	return strings.Replace(g, "/", "_", -1)
}
