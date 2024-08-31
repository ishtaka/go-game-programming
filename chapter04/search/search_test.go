package search

import "testing"

func createWeightedGraph(t *testing.T) *WeightedGraph {
	t.Helper()

	l := 5
	g := NewWeightedGraph(l * l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			node := NewWeightedGraphNode(4)
			g.nodes = append(g.nodes, node)
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			node := g.nodes[i*l+j]
			if i > 0 {
				from := node
				to := g.nodes[(i-1)*l+j]
				edge := NewWeightedEdge(from, to, 1.0)
				node.edges = append(node.edges, edge)
			}
			if i < l-1 {
				from := node
				to := g.nodes[(i+1)*l+j]
				edge := NewWeightedEdge(from, to, 1.0)
				node.edges = append(node.edges, edge)
			}
			if j > 0 {
				from := node
				to := g.nodes[(i*l)+j-1]
				edge := NewWeightedEdge(from, to, 1.0)
				node.edges = append(node.edges, edge)
			}
			if j < l-1 {
				from := node
				to := g.nodes[(i*l)+j+1]
				edge := NewWeightedEdge(from, to, 1.0)
				node.edges = append(node.edges, edge)
			}
		}
	}

	return g
}
