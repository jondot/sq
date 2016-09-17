package main

import "fmt"

// GroupPrinter tbd
type GroupPrinter struct {
}

// VisitHeader tbd
func (g *GroupPrinter) VisitHeader(group interface{}) {
	groupName := fmt.Sprintf("group_%s", group.(string))
	fmt.Printf("%s\n", groupName)
}

// VisitNode tbd
func (g *GroupPrinter) VisitNode(i int, sz int, header interface{}, file interface{}) {
	fi := file.(string)
	if i == sz-1 {
		fmt.Printf("└── ")
		fmt.Printf("%s\n%v item(s)\n\n", fi, sz)
	} else {
		fmt.Printf("├── ")
		fmt.Printf("%s\n", fi)
	}
}
