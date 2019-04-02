package graph

import (
	"io"
)

// Graph is a symmetric adjacency map representation of an undirected graph
type Graph struct {
	*dirAdj
	Name   string
	invAdj *dirAdj
}

// NewGraph creates a new undirected graph
func NewGraph(name string, readers ...io.ReadCloser) (*Graph, error) {
	g := &Graph{
		dirAdj: &dirAdj{},
		Name:   name,
	}
	// undireted graph has a symmetric adjacency structure so the inverse adjacency
	// (inverted index) is just a pointer to the adjacency map
	g.invAdj = g.dirAdj

	for _, r := range readers {
		err := g.addFromReader(r)
		if err != nil {
			return g, err
		}
	}
	return g, nil
}

// AddEdge adds an edge between two nodes with an optional weight that defaults to 1.0
func (g *Graph) AddEdge(src string, tgt string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}
	g.addDirectedEdge(src, tgt, wgt)
	g.invAdj.addDirectedEdge(tgt, src, wgt)
}

// RemoveEdge removes an edge between two nodes
func (g *Graph) RemoveEdge(src string, tgt string) {
	g.removeDirectedEdge(src, tgt)
	g.invAdj.removeDirectedEdge(tgt, src)
}

// RemoveNode removes a node entirely from a Graph such that
// no edges exist between it an any other node
func (g *Graph) RemoveNode(node string) {
	if nbrs, ok := g.GetNeighbors(node); ok {
		for n := range nbrs {
			g.RemoveEdge(node, n)
		}
	}
}

// PrintInv displays a DirGraph's incoming adjacency structure
func (g *Graph) PrintInv() {
	g.invAdj.Print()
}

// GetNodes gets a slice of all nodes in a Graph
func (g *Graph) GetNodes() (nodes []string) {
	return g.getSrcNodes()
}

// GetInvNeighbors gets a slice of nodes that have an edge from them to a specified node
func (g *Graph) GetInvNeighbors(node string) (nbrs map[string]float64, found bool) {
	return g.invAdj.GetNeighbors(node)
}

// GetTotalDegree calculates the sum of weights of all edges with node as the source node
func (g *Graph) GetTotalDegree(node string) (deg float64, found bool) {
	return g.GetOutDegree(node)
}

// GetDegree calculates the sum of weights of all edges with node as the source node
func (g *Graph) GetDegree(node string) (deg float64, found bool) {
	return g.GetOutDegree(node)
}

// GetInDegree calculates the sum of weights of all edges with node as the target node
func (g *Graph) GetInDegree(node string) (deg float64, found bool) {
	return g.invAdj.GetOutDegree(node)
}
