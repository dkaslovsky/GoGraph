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

func (a DirAdj) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}

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

func (a DirAdj) GetEdge() {

}
