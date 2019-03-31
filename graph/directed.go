package graph

import (
	"io"
)

// DirGraph is an adjacency map representation of a directed graph
type DirGraph struct {
	dirAdj
	Name   string
	invAdj dirAdj // inverted adjacency for faster lookups
}

// NewDirGraph creates a new directed graph
func NewDirGraph(name string, readers ...io.ReadCloser) (*DirGraph, error) {
	dg := &DirGraph{
		dirAdj: dirAdj{},
		Name:   name,
		invAdj: dirAdj{},
	}
	for _, r := range readers {
		err := dg.addFromReader(r)
		if err != nil {
			return dg, err
		}
	}
	return dg, nil
}

// AddEdge adds an edge from a node to another node with an optional weight that defaults to 1.0
func (dg *DirGraph) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}
	dg.addDirectedEdge(from, to, wgt)
	dg.invAdj.addDirectedEdge(to, from, wgt)
}

// RemoveEdge removes an edge that exists from a node to another node
func (dg *DirGraph) RemoveEdge(from string, to string) {
	dg.removeDirectedEdge(from, to)
	dg.invAdj.removeDirectedEdge(to, from)
}

// RemoveNode removes a node entirely from a DirGraph such that
// no edges exist between it an any other node
func (dg *DirGraph) RemoveNode(node string) {
	if nbrs, ok := dg.GetInvNeighbors(node); ok {
		for _, n := range nbrs {
			dg.RemoveEdge(n, node)
		}
	}
	if nbrs, ok := dg.GetNeighbors(node); ok {
		for _, n := range nbrs {
			dg.RemoveEdge(node, n)
		}
	}
}

// PrintInv displays a DirGraph's incoming adjacency structure
func (dg *DirGraph) PrintInv() {
	dg.invAdj.Print()
}

// GetNodes gets a slice of all nodes in a DirGraph
func (dg *DirGraph) GetNodes() (nodes []string) {
	set := map[string]struct{}{}
	for _, node := range dg.getFromNodes() {
		if _, ok := set[node]; !ok {
			set[node] = struct{}{}
			nodes = append(nodes, node)
		}
	}
	for _, node := range dg.invAdj.getFromNodes() {
		if _, ok := set[node]; !ok {
			set[node] = struct{}{}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

// GetInvNeighbors gets a slice of nodes that have an edge from them to a specified node
func (dg *DirGraph) GetInvNeighbors(node string) (nbrs []string, found bool) {
	return dg.invAdj.GetNeighbors(node)
}

// GetTotalDegree calculates the sum of weights of all edges from and to a node
func (dg *DirGraph) GetTotalDegree(node string) (deg float64, found bool) {
	outDeg, ok := dg.GetOutDegree(node)
	if !ok {
		return deg, false
	}
	inDeg, ok := dg.GetInDegree(node)
	if !ok {
		return deg, false
	}
	return inDeg + outDeg, true
}

// GetDegree calculates the sum of weights of all edges from a node
func (dg *DirGraph) GetDegree(node string) (deg float64, found bool) {
	return dg.GetOutDegree(node)
}

// GetOutDegree calculates the sum of weights of all edges from a node
func (dg *DirGraph) GetOutDegree(node string) (deg float64, found bool) {
	nbrs, ok := dg.GetNeighbors(node)
	if !ok {
		return deg, false
	}
	for _, n := range nbrs {
		if w, ok := dg.GetEdgeWeight(node, n); ok {
			deg += w
		}
	}
	return deg, true
}

// GetInDegree calculates the sum of weights of all edges to a node
func (dg *DirGraph) GetInDegree(node string) (deg float64, found bool) {
	nbrs, ok := dg.GetInvNeighbors(node)
	if !ok {
		return deg, false
	}
	for _, n := range nbrs {
		if w, ok := dg.GetEdgeWeight(n, node); ok {
			deg += w
		}
	}
	return deg, true
}
