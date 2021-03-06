package search

import (
	n "github.com/dkaslovsky/GoGraph/node"
)

type hasNodeNeighborGetter interface {
	HasNode(n.Node) bool
	GetNeighbors(n.Node) (map[n.Node]float64, bool)
}

// DFS performs a depth first search starting at a specified node
func DFS(g hasNodeNeighborGetter, node n.Node) []n.Node {

	if !g.HasNode(node) {
		return []n.Node{}
	}

	visited := n.NewSet()

	s := n.NewStack()
	s.Push(node)

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

// BFS performs a breadth first search starting at a specified node
func BFS(g hasNodeNeighborGetter, node n.Node) []n.Node {
	if !g.HasNode(node) {
		return []n.Node{}
	}

	visited := n.NewSet()

	q := n.NewQueue()
	q.Push(node)

	for q.Len() > 0 {
		curNode, _ := q.Pop() // no need to check error since the queue cannot be empty here
		if visited.Contains(curNode) {
			continue
		}
		visited.Add(curNode)

		nbrs, ok := g.GetNeighbors(curNode)
		if !ok {
			continue
		}
		for nbr := range nbrs {
			q.Push(nbr)
		}
	}

	return visited.ToSlice()
}
