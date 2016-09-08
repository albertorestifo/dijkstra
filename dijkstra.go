// Package dijkstra implements the Dijkstra shortest path algorithm
package dijkstra

import "fmt"

// Node is the rapresentation fo node in the graph with a key and a cost
type Node struct {
	key  string
	cost int
}

// Neighbors is a map of adjacent nodes and the cost to reac them
type Neighbors map[string]int

// Graph is a rappresentation of our graph
type Graph map[string]Neighbors

// Path find the shortest path between start and target
func (graph Graph) Path(start, target string) (path []string, cost int, err error) {
	if len(graph) == 0 {
		err = fmt.Errorf("cannot find path in empty map")
		return
	}

	// ensure start and target are part of the graph
	if _, ok := graph[start]; !ok {
		err = fmt.Errorf("cannot find start %v in graph", start)
		return
	}
	if _, ok := graph[target]; !ok {
		err = fmt.Errorf("cannot find target %v in graph", target)
		return
	}

	explored := make(map[string]bool)   // set of nodes we already explored
	frontier := NewQueue()              // queue of the nodes to explore
	previous := make(map[string]string) // previously visited node

	// add starting point to the frontier as it'll be the first node visited
	frontier.Set(start, 0)

	// run until we visited every node in the frontier
	for !frontier.IsEmpty() {
		// get the node in the frontier with the lowest cost (or priority)
		aKey, aPriority := frontier.Next()
		node := Node{aKey, aPriority}

		// when the node with the lowest cost in the frontier is target, we can
		// compute the cost and path and exit the loop
		if node.key == target {
			cost = node.cost

			nKey := node.key
			for nKey != start {
				path = append(path, nKey)
				nKey = previous[nKey]
			}

			break
		}

		// add the current node to the explored set
		explored[node.key] = true

		// loop all the neighboring nodes
		for nKey, nCost := range graph[node.key] {
			// skip alreadt-explored nodes
			if explored[nKey] {
				continue
			}

			// if the node is not yet in the frontier add it with the cost
			if _, ok := frontier.Get(nKey); !ok {
				previous[nKey] = node.key
				frontier.Set(nKey, node.cost+nCost)
				continue
			}

			frontierCost, _ := frontier.Get(nKey)
			nodeCost := node.cost + nCost

			// only update the cost of this node in the frontier when
			// it's below what's currently set
			if nodeCost < frontierCost {
				previous[nKey] = node.key
				frontier.Set(nKey, nodeCost)
			}
		}
	}

	// add the origin at the end of the path
	path = append(path, start)

	// reverse the path because it was popilated
	// in reverse, form target to start
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return
}
