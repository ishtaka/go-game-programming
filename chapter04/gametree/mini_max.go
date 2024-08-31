package gametree

import (
	"math"
)

func MaxPlayer(node *Node) float64 {
	// If this is a leaf, return score
	if len(node.Children) == 0 {
		return node.State.GetScore()
	}

	maxValue := math.Inf(-1)
	// Find the subtree with the maximum value
	for _, child := range node.Children {
		maxValue = math.Max(maxValue, MinPlayer(child))
	}

	return maxValue
}

func MinPlayer(node *Node) float64 {
	// If this is a leaf, return score
	if len(node.Children) == 0 {
		return node.State.GetScore()
	}

	minValue := math.Inf(1)
	// Find the subtree with the minimum value
	for _, child := range node.Children {
		minValue = math.Min(minValue, MaxPlayer(child))
	}

	return minValue
}

func MinimaxDecide(root *Node) *Node {
	// Find the subtree with the maximum value, and save the choice
	var choice *Node
	maxValue := math.Inf(-1)
	for _, child := range root.Children {
		v := MinPlayer(child)
		if v > maxValue {
			maxValue = v
			choice = child
		}
	}

	return choice
}
