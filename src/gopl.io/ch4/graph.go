package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {

	addEdge("a", "c")
	addEdge("b", "c")
	addEdge("a", "b")
	addEdge("a", "d")

	for d, e := range graph {
		fmt.Printf("from %s connect to ", d)
		for f, t := range e {
			if t {
				fmt.Printf("%s,", f)
			}
		}
		fmt.Println()
	}
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

