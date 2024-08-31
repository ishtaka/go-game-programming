package search

import (
	"cmp"
	"slices"
)

type GBFSScratch struct {
	ParentEdge  *WeightedEdge
	Heuristic   float32
	InOpenSet   bool
	InClosedSet bool
}

type GBFSMap = map[*WeightedGraphNode]*GBFSScratch

func GBFS(g *WeightedGraph, start, goal *WeightedGraphNode, outMap GBFSMap) bool {
	openSet := make([]*WeightedGraphNode, 0, len(g.nodes))

	// Set current node to start, and mark in closed set
	current := start
	outMap[current].InClosedSet = true

	for {
		// Add adjacent nodes to open set
		for _, edge := range current.edges {
			// Get scratch data for this node
			data := outMap[edge.to]
			// Add it only if it's not in the closed set
			if !data.InClosedSet {
				// Set the adjacent node's parent edge
				data.ParentEdge = edge
				if !data.InOpenSet {
					// Compute the heuristic for this node, and add to open set
					data.Heuristic = edge.to.Heuristic(goal)
					data.InOpenSet = true
					openSet = append(openSet, edge.to)
				}
			}
		}

		// If open set is empty, all possible paths are exhausted
		if len(openSet) == 0 {
			break
		}

		// Find the lowest cost node in open set
		lowest := slices.MinFunc(openSet, func(a, b *WeightedGraphNode) int {
			return cmp.Compare(outMap[a].Heuristic, outMap[b].Heuristic)
		})

		// Set to current and move from open to closed
		current = lowest
		openSet = slices.DeleteFunc(openSet, func(n *WeightedGraphNode) bool {
			return n == lowest
		})
		outMap[current].InOpenSet = false
		outMap[current].InClosedSet = true

		if current == goal {
			break
		}
	}

	// Did we find a path?
	return current == goal
}
