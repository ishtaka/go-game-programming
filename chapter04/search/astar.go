package search

import (
	"cmp"
	"slices"
)

type AStarScratch struct {
	ParentEdge      *WeightedEdge
	Heuristic       float32
	ActualFromStart float32
	InOpenSet       bool
	InClosedSet     bool
}

type AStarMap = map[*WeightedGraphNode]*AStarScratch

func AStar(g *WeightedGraph, start, goal *WeightedGraphNode, outMap AStarMap) bool {
	openSet := make([]*WeightedGraphNode, 0, len(g.nodes))

	// Set current node to start, and mark in closed set
	current := start
	outMap[current].InClosedSet = true

	for {
		// Add adjacent nodes to open set
		for _, edge := range current.edges {
			neighbor := edge.to
			// Get scratch data for this node
			data := outMap[neighbor]
			// Only check nodes that aren't in the closed set
			if !data.InClosedSet {
				// Not in the open set, so parent must be current
				if !data.InOpenSet {
					data.ParentEdge = edge
					data.Heuristic = neighbor.Heuristic(goal)
					// Actual cost is the parent's plus cost of traversing edge
					data.ActualFromStart = outMap[current].ActualFromStart + edge.weight
					data.InOpenSet = true
					openSet = append(openSet, neighbor)
				} else {
					// Compute what new actual cost is if current becomes parent
					newG := outMap[current].ActualFromStart + edge.weight
					if newG < data.ActualFromStart {
						// Current should adopt this node
						data.ParentEdge = edge
						data.ActualFromStart = newG
					}
				}
			}
		}

		// If open set is empty, all possible paths are exhausted
		if len(openSet) == 0 {
			break
		}

		// Find the lowest cost node in open set
		lowest := slices.MinFunc(openSet, func(a, b *WeightedGraphNode) int {
			fOfA := outMap[a].Heuristic + outMap[a].ActualFromStart
			fOfB := outMap[b].Heuristic + outMap[b].ActualFromStart
			return cmp.Compare(fOfA, fOfB)
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
