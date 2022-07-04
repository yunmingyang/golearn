package main

import (
	"fmt"
	"sort"
)



var prereqs = map[string][]string {
	"algorithms": {"data structures"},
	"calculus": {"line algebra"},
	"compilers": {
		"data structures",
		"formal  languages",
		"computer organization",
	},
	"data structures": {"discreate math"},
	"databases": {"data structures"},
	"discreate math": {"intro to programming"},
	"formal  languages": {"discreate math"},
	"networks": {"opreating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}


func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	visitAll(keys)

	return order
}