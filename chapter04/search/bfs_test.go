package search

import "testing"

func TestBFS(t *testing.T) {
	l := 5
	g := NewGraph(l * l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			node := NewGraphNode(4)
			g.nodes = append(g.nodes, node)
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			node := g.nodes[i*l+j]
			if i > 0 {
				node.adjacent = append(node.adjacent, g.nodes[(i-1)*l+j])
			}
			if i < l-1 {
				node.adjacent = append(node.adjacent, g.nodes[(i+1)*l+j])
			}
			if j > 0 {
				node.adjacent = append(node.adjacent, g.nodes[i*l+j-1])
			}
			if j < l-1 {
				node.adjacent = append(node.adjacent, g.nodes[i*l+j+1])
			}
		}
	}

	m := make(NodeToParentMap, l)
	if found := BFS(g, g.nodes[0], g.nodes[9], m); !found {
		t.Error("path not found")
	}
}
