package dijkstra

import (
	"fmt"
	"testing"
)

func TestEmptyGraph(t *testing.T) {
	g := make(Graph)

	_, _, err := g.Path("a", "z")
	if err == nil {
		t.Error("Error nil; want error message")
	}
}

func TestGraphErrors(t *testing.T) {
	g := Graph{
		"a": {"b": 20, "c": 80},
		"b": {"a": 20, "c": 20},
		"c": {"a": 80, "b": 20},
	}

	_, _, err := g.Path("a", "z")
	if err == nil {
		t.Error("err = nil; want not in graph error")
	}

	_, _, err = g.Path("z", "c")
	if err == nil {
		t.Error("err = nil; want not in graph error")
	}
}

func TestPath1(t *testing.T) {
	g := Graph{
		"a": {"b": 20, "c": 80},
		"b": {"a": 20, "c": 20},
		"c": {"a": 80, "b": 20},
	}

	// The shortest path is correct
	path, cost, err := g.Path("a", "c")
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}

	expectedPath := []string{"a", "b", "c"}

	if len(path) != len(expectedPath) {
		t.Errorf("path = %v; want %v", path, expectedPath)
	}
	for i, key := range path {
		if key != expectedPath[i] {
			t.Errorf("path = %v; want %v", path, expectedPath)
		}
	}

	expectedCost := 40
	if cost != expectedCost {
		t.Errorf("cost = %v; want %v", cost, expectedCost)
	}
}

func TestPath2(t *testing.T) {
	g := Graph{
		"a": {"b": 7, "c": 9, "f": 14},
		"b": {"c": 10, "d": 15},
		"c": {"d": 11, "f": 2},
		"d": {"e": 6},
		"e": {"f": 9},
	}

	// The shortest path is correct
	path, _, err := g.Path("a", "e")
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}

	expectedPath := []string{"a", "c", "d", "e"}

	if len(path) != len(expectedPath) {
		t.Errorf("path = %v; want %v", path, expectedPath)
	}
	for i, key := range path {
		if key != expectedPath[i] {
			t.Errorf("path = %v; want %v", path, expectedPath)
		}
	}
}

func BenchmarkPath(b *testing.B) {
	g := Graph{
		"a": {"b": 20, "c": 80},
		"b": {"a": 20, "c": 20},
		"c": {"a": 80, "b": 20},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Path("a", "c")
	}
}

func ExampleGraph_Path() {
	g := Graph{
		"a": {"b": 20, "c": 80},
		"b": {"a": 20, "c": 20},
		"c": {"a": 80, "b": 20},
	}

	path, cost, _ := g.Path("a", "c") // skipping error handling

	fmt.Printf("path: %v, cost: %v", path, cost)
	// Output: path: [a b c], cost: 40
}
