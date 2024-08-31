package search

import "testing"

func TestGBFS(t *testing.T) {
	g := createWeightedGraph(t)

	m := make(GBFSMap, len(g.nodes))
	for _, node := range g.nodes {
		m[node] = &GBFSScratch{}
	}

	if found := GBFS(g, g.nodes[0], g.nodes[9], m); !found {
		t.Error("path not found")
	}
}
