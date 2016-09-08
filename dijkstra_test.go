package dijkstra

import "testing"

func TestEmptyGraph(t *testing.T) {
	graph := make(Graph)

	_, _, err := graph.Path("a", "z")
	if err == nil {
		t.Error("expected empty graph to return an error")
	}
}

func TestGraphErrors(t *testing.T) {
	graph := Graph{
		"a": Neighbors{
			"b": 20,
			"c": 80,
		},
		"b": Neighbors{
			"a": 20,
			"c": 20,
		},
		"c": Neighbors{
			"a": 80,
			"b": 20,
		},
	}

	_, _, err := graph.Path("a", "z")
	if err == nil {
		t.Error("expected empty graph to return an error")
	}

	_, _, err = graph.Path("z", "c")
	if err == nil {
		t.Error("expected empty graph to return an error")
	}
}

func TestPath1(t *testing.T) {
	graph := Graph{
		"a": Neighbors{
			"b": 20,
			"c": 80,
		},
		"b": Neighbors{
			"a": 20,
			"c": 20,
		},
		"c": Neighbors{
			"a": 80,
			"b": 20,
		},
	}

	// The shortest path is correct
	path, cost, err := graph.Path("a", "c")
	if err != nil {
		t.Errorf("could not get path from a to c: %v", err)
	}

	expectedPath := []string{"a", "b", "c"}

	if len(path) != len(expectedPath) {
		t.Errorf("expected path %v to match %v", path, expectedPath)
	}
	for i, key := range path {
		if key != expectedPath[i] {
			t.Errorf("expected path %v to match %v", path, expectedPath)
		}
	}

	expectedCost := 40
	if cost != expectedCost {
		t.Errorf("expected cost %v to match %v", cost, expectedCost)
	}
}

func TestPath2(t *testing.T) {
	graph := Graph{
		"a": Neighbors{
			"b": 7,
			"c": 9,
			"f": 14,
		},
		"b": Neighbors{
			"c": 10,
			"d": 15,
		},
		"c": Neighbors{
			"d": 11,
			"f": 2,
		},
		"d": Neighbors{"e": 6},
		"e": Neighbors{"f": 9},
	}

	// The shortest path is correct
	path, _, err := graph.Path("a", "e")
	if err != nil {
		t.Errorf("could not get path from a to e: %v", err)
	}

	expectedPath := []string{"a", "c", "d", "e"}

	if len(path) != len(expectedPath) {
		t.Errorf("expected path %v to match %v", path, expectedPath)
	}
	for i, key := range path {
		if key != expectedPath[i] {
			t.Errorf("expected path %v to match %v", path, expectedPath)
		}
	}
}

func BenchmarkPath(b *testing.B) {
	g := Graph{
		"a": Neighbors{
			"b": 20,
			"c": 80,
		},
		"b": Neighbors{
			"a": 20,
			"c": 20,
		},
		"c": Neighbors{
			"a": 80,
			"b": 20,
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Path("a", "c")
	}
}
