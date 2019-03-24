package graph

import (
	"fmt"
)

type DirAdj map[string]map[string]float64

func (a DirAdj) Print() {
	for node := range a {
		fmt.Printf("%s:\n", node)
		nbrs := a[node]
		for n, wgt := range nbrs {
			fmt.Printf(" -->  %s: %f\n", n, wgt)
		}
	}
}

func (a DirAdj) AddEdge(from string, to string, wgt float64) {
	nbrs, ok := a[from]
	if !ok {
		a[from] = map[string]float64{to: wgt}
	} else {
		nbrs[to] = wgt
	}
}

func (a DirAdj) RemoveEdge(from string, to string) {
	nbrs, ok := a[from]
	if !ok {
		return
	}
	delete(nbrs, to)
	// delete from node if it no longer has neighbors
	if len(nbrs) == 0 {
		a.RemoveNode(from)
	}
}

func (a DirAdj) RemoveNode(node string) {
	delete(a, node)
	for n := range a {
		a.RemoveEdge(n, node)
	}
}

func (a DirAdj) GetNeighbors(node string) (nbrs []string, found bool) {
	adj, ok := a[node]
	if !ok {
		return nbrs, false
	}
	for n := range adj {
		nbrs = append(nbrs, n)
	}
	return nbrs, true
}

func (a DirAdj) HasEdge(from string, to string) bool {
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

func (a DirAdj) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	if a.HasEdge(from, to) {
		w := a[from][to]
		return w, true
	}
	return weight, false
}
