package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"alogrithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {

	var ordered []string
	for k := range m {
		ordered = append(ordered, k)
		sort.Strings(ordered)
	}

	//keep track of what has been processed
	seen := make(map[string]bool)

	var result []string
	//receives the course to check and the result to append to
	var f func(k string)
	f = func(course string) {
		if !seen[course] {
			seen[course] = true
			for _, dep := range m[course] {
				f(dep)
			}
			result = append(result, course)
		}
	}

	for _, k := range ordered {
		f(k)
	}
	return result
}

func main() {
	fmt.Printf("%#v\n", prereqs)
	sorted := topoSort(prereqs)
	for i, course := range sorted {
		fmt.Printf("%d\t%s\n", (i + 1), course)
	}
}
