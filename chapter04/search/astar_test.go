package search

import "testing"

func TestAStar(t *testing.T) {
	g := createWeightedGraph(t)

	m := make(AStarMap, len(g.nodes))
	for _, node := range g.nodes {
		m[node] = &AStarScratch{}
	}

	if found := AStar(g, g.nodes[0], g.nodes[9], m); !found {
		t.Error("path not found")
	}
}
