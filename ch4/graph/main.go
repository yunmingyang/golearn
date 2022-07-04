package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("test", "test1")
	fmt.Println(graph)

	fmt.Println(hasEdge("test", "test"))
	fmt.Println(hasEdge("test", "test1"))
}


func addEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool{
	return graph[from][to]
}