package graph

import (
	"fmt"
)

type Node string

type dirAdj map[Node]map[Node]float64

// Print prints the adjacency structure
func (a dirAdj) Print() {
	for node, nbrs := range a {
		fmt.Printf("%s:\n", node)
		for n, wgt := range nbrs {
			fmt.Printf(" -->  %s: %f\n", n, wgt)
		}
	}
}

func (a dirAdj) addDirectedEdge(src Node, tgt Node, wgt float64) {
	if nbrs, ok := a[src]; ok {
		nbrs[tgt] = wgt
		return
	}
	a[src] = map[Node]float64{tgt: wgt}
}

func (a dirAdj) removeDirectedEdge(src Node, tgt Node) {
	nbrs, ok := a[src]
	if !ok {
		return
	}
	delete(nbrs, tgt)
	// delete src node if it no longer has neighbors
	if len(nbrs) == 0 {
		delete(a, src)
	}
}

func (a dirAdj) getSrcNodes() (nodes []Node) {
	for node := range a {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetNeighbors gets the nodes that a specified node connects to with an edge
func (a dirAdj) GetNeighbors(node Node) (map[Node]float64, bool) {
	nbrs, ok := a[node]
	return nbrs, ok
}

// GetOutDegree calculates the sum of weights of all edges with node as the source node
func (a dirAdj) GetOutDegree(node Node) (deg float64, found bool) {
	nbrs, ok := a.GetNeighbors(node)
	if !ok {
		return deg, false
	}
	for _, w := range nbrs {
		deg += w
	}
	return deg, true
}

// HasEdge returns true if an edge exists from a node to another node, false otherwise
func (a dirAdj) HasEdge(src Node, tgt Node) bool {
	nbrs, ok := a.GetNeighbors(src)
	if !ok {
		return false
	}
	_, ok = nbrs[tgt]
	return ok
}

// GetEdgeWeight returns the weight of the edge from a node to another node if it exists
func (a dirAdj) GetEdgeWeight(src Node, tgt Node) (weight float64, found bool) {
	if !a.HasEdge(src, tgt) {
		return weight, false
	}
	weight = a[src][tgt]
	return weight, true
}
