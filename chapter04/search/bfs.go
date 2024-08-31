package search

type GraphNode struct {
	// Adjacency list
	adjacent []*GraphNode
}

func NewGraphNode(size int) *GraphNode {
	return &GraphNode{
		adjacent: make([]*GraphNode, 0, size),
	}
}

type Graph struct {
	// A graph contains nodes
	nodes []*GraphNode
}

func NewGraph(size int) *Graph {
	return &Graph{
		nodes: make([]*GraphNode, 0, size),
	}
}

type NodeToParentMap = map[*GraphNode]*GraphNode

func BFS(graph *Graph, start, goal *GraphNode, outMap NodeToParentMap) bool {
	// Whether we found a path
	pathFound := false
	// Nodes to consider
	queue := NewQueue[*GraphNode](len(graph.nodes))
	// Enqueue the first node
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		// Dequeue a node
		current := queue.Dequeue()
		if current == goal {
			pathFound = true
			break
		}

		// Enqueue adjacent nodes that aren't already in the queue
		for _, node := range current.adjacent {
			// If node hasn't been enqueued
			// (except for the start node)
			if _, ok := outMap[node]; !ok && node != start {
				// Enqueue this node, setting its parent
				outMap[node] = current
				queue.Enqueue(node)
			}
		}
	}

	return pathFound
}
