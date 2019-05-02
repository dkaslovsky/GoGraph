package search

import (
	n "github.com/dkaslovsky/GoGraph/node"
)

type neighborGetter interface {
	GetNeighbors(n.Node) (map[n.Node]float64, bool)
}

// DFS performs a depth first search starting at a specified node
func DFS(g neighborGetter, node n.Node) (nodes []n.Node) {

	// add HasNode to graph(s) with tests
	// add HasNode to interface and rename
	// return error here?
	if !g.HasNode(node) {
		return nodes
	}

	s := n.NewNodeStack()
	s.Push(node)

	visited := n.NewNodeSet()

	for s.Len() > 0 {
		curNode, _ := s.Pop() // no need to check error since the stack cannot be empty here
		if visited.Contains(curNode) {
			continue
		}
		visited.Add(curNode)

		nbrs, ok := g.GetNeighbors(curNode)
		if !ok {
			continue
		}
		for nbr := range nbrs {
			s.Push(nbr)
		}
	}

	return visited.ToSlice()
}
