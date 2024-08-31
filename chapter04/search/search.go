package search

type WeightedEdge struct {
	// Which nodes are connected by this edge?
	from *WeightedGraphNode
	to   *WeightedGraphNode
	// Weight of this edge
	weight float32
}

func NewWeightedEdge(from, to *WeightedGraphNode, weight float32) *WeightedEdge {
	return &WeightedEdge{
		from:   from,
		to:     to,
		weight: weight,
	}
}

type WeightedGraphNode struct {
	edges []*WeightedEdge
}

func NewWeightedGraphNode(size int) *WeightedGraphNode {
	return &WeightedGraphNode{
		edges: make([]*WeightedEdge, 0, size),
	}
}

func (w *WeightedGraphNode) Heuristic(_ *WeightedGraphNode) float32 {
	return 0.0

}

type WeightedGraph struct {
	nodes []*WeightedGraphNode
}

func NewWeightedGraph(size int) *WeightedGraph {
	return &WeightedGraph{
		nodes: make([]*WeightedGraphNode, 0, size),
	}
}
