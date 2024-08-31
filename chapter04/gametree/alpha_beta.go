package gametree

import "math"

func AlphaBetaMax(node *Node, alpha, beta float64) float64 {
	// If this is a leaf, return score
	if len(node.Children) == 0 {
		return node.State.GetScore()
	}

	maxValue := math.Inf(-1)
	// Find the subtree with the maximum value
	for _, child := range node.Children {
		maxValue = math.Max(maxValue, AlphaBetaMin(child, alpha, beta))
		if maxValue >= beta {
			return maxValue // Beta prune
		}
		alpha = math.Max(maxValue, alpha)
	}
	return maxValue
}

func AlphaBetaMin(node *Node, alpha, beta float64) float64 {
	// If this is a leaf, return score
	if len(node.Children) == 0 {
		return node.State.GetScore()
	}

	minValue := math.Inf(1)
	// Find the subtree with the minimum value
	for _, child := range node.Children {
		minValue = math.Min(minValue, AlphaBetaMax(child, alpha, beta))
		if minValue <= alpha {
			return minValue // Alpha prune
		}
		beta = math.Min(minValue, beta)
	}
	return minValue
}

func AlphaBetaDecide(root *Node) *Node {
	// Find the subtree with the maximum value, and save the choice
	var choice *Node
	maxValue := math.Inf(-1)
	beta := math.Inf(1)
	for _, child := range root.Children {
		v := AlphaBetaMin(child, maxValue, beta)
		if v > maxValue {
			maxValue = v
			choice = child
		}
	}

	return choice
}
