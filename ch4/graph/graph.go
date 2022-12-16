package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("a", "e")
	fmt.Println(hasEdge("a", "e"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
