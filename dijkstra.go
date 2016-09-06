package dijkstra

import "fmt"

// Neighbors is a map of adjacent nodes and the cost to reac them
type Neighbors map[string]int

// Graph is a rappresentation of our graph
type Graph map[string]*Neighbors

// Path find the shortest path between start and target
func Path(graph Graph, start, target string) (path []string, cost int, err error) {
	if len(graph) == 0 {
		err = fmt.Errorf("cannot find path in empty map")
		return
	}

	// Initialize the values
	// explored := make(map[string]bool)
	// frontier := new(Queue)

	return
}
