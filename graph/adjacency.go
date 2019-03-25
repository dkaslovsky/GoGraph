package graph

import (
	"fmt"
)

type dirAdj map[string]map[string]float64

func (a dirAdj) Print() {
	for node := range a {
		fmt.Printf("%s:\n", node)
		nbrs := a[node]
		for n, wgt := range nbrs {
			fmt.Printf(" -->  %s: %f\n", n, wgt)
		}
	}
}

func (a dirAdj) AddEdge(from string, to string, wgt float64) {
	nbrs, ok := a[from]
	if !ok {
		a[from] = map[string]float64{to: wgt}
	} else {
		nbrs[to] = wgt
	}
}

func (a dirAdj) RemoveEdge(from string, to string) {
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

// RemoveNode removes a node from dirAdj
// Note that this is an O(n) operation, so if an inverted index is available it will be
// more efficient (i.e. O(k) where k is th number nodes for which node is a to-neighbor)
// to use that structure to remove a node
func (a dirAdj) RemoveNode(node string) {
	delete(a, node)
	for n := range a {
		a.RemoveEdge(n, node)
	}
}

func (a dirAdj) GetNeighbors(node string) (nbrs []string, found bool) {
	adj, ok := a[node]
	if !ok {
		return nbrs, false
	}
	for n := range adj {
		nbrs = append(nbrs, n)
	}
	return nbrs, true
}

func (a dirAdj) HasEdge(from string, to string) bool {
	nbrs, ok := a[from]
	if !ok {
		return false
	}
	_, ok = nbrs[to]
	if !ok {
		return false
	}
	return true
}

func (a dirAdj) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	if a.HasEdge(from, to) {
		w := a[from][to]
		return w, true
	}
	return weight, false
}
