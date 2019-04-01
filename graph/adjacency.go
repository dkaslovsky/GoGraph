package graph

import (
	"fmt"
)

type dirAdj map[string]map[string]float64

// Print prints the adjacency structure
func (a dirAdj) Print() {
	for node := range a {
		fmt.Printf("%s:\n", node)
		nbrs := a[node]
		for n, wgt := range nbrs {
			fmt.Printf(" -->  %s: %f\n", n, wgt)
		}
	}
}

func (a dirAdj) addDirectedEdge(from string, to string, wgt float64) {
	if nbrs, ok := a[from]; ok {
		nbrs[to] = wgt
		return
	}
	a[from] = map[string]float64{to: wgt}
}

func (a dirAdj) removeDirectedEdge(from string, to string) {
	nbrs, ok := a[from]
	if !ok {
		return
	}
	delete(nbrs, to)
	// delete from node if it no longer has neighbors
	if len(nbrs) == 0 {
		delete(a, from)
	}
}

func (a dirAdj) getFromNodes() (nodes []string) {
	for node := range a {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetNeighbors gets the nodes that a specified node connects to with an edge
func (a dirAdj) GetNeighbors(node string) (nbrs map[string]float64, found bool) {
	adj, ok := a[node]
	return adj, ok
}

// GetOutDegree calculates the sum of weights of all edges from a node
func (a dirAdj) GetOutDegree(node string) (deg float64, found bool) {
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
func (a dirAdj) HasEdge(from string, to string) bool {
	nbrs, ok := a.GetNeighbors(from)
	if !ok {
		return false
	}
	_, ok = nbrs[to]
	if !ok {
		return false
	}
	return true
}

// GetEdgeWeight returns the weight of the edge from a node to another node if it exists
func (a dirAdj) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	if a.HasEdge(from, to) {
		w := a[from][to]
		return w, true
	}
	return weight, false
}
