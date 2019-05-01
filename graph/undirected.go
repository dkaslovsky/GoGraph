package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// defaultWgt is the weight applied to an edge when a weight is not specified
const defaultWgt float64 = 1.0

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

func (g *Graph) addFromReader(r io.ReadCloser) error {
	defer r.Close()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		src := Node(parts[0])
		tgt := Node(parts[1])
		if src == "" || tgt == "" {
			continue
		}

		if len(parts) == 2 {
			g.AddEdge(src, tgt)
			continue
		}

		weight, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return err
		}
		g.AddEdge(src, tgt, weight)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// AddEdge adds an edge between two nodes with an optional weight that defaults to 1.0
func (g *Graph) AddEdge(src Node, tgt Node, weight ...float64) {
	wgt := defaultWgt
	if len(weight) > 0 {
		wgt = weight[0]
	}
	g.addDirectedEdge(src, tgt, wgt)
	g.invAdj.addDirectedEdge(tgt, src, wgt)
}

// RemoveEdge removes an edge between two nodes
func (g *Graph) RemoveEdge(src Node, tgt Node) {
	g.removeDirectedEdge(src, tgt)
	g.invAdj.removeDirectedEdge(tgt, src)
}

// RemoveNode removes a node entirely from a Graph such that
// no edges exist between it an any other node
func (g *Graph) RemoveNode(node Node) {
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
func (g *Graph) GetNodes() []Node {
	return g.getSrcNodes()
}

// GetInvNeighbors gets a slice of nodes that have an edge from them to a specified node
func (g *Graph) GetInvNeighbors(node Node) (map[Node]float64, bool) {
	return g.invAdj.GetNeighbors(node)
}

// GetTotalDegree calculates the sum of weights of all edges with node as the source node
func (g *Graph) GetTotalDegree(node Node) (float64, bool) {
	return g.GetOutDegree(node)
}

// GetDegree calculates the sum of weights of all edges with node as the source node
func (g *Graph) GetDegree(node Node) (float64, bool) {
	return g.GetOutDegree(node)
}

// GetInDegree calculates the sum of weights of all edges with node as the target node
func (g *Graph) GetInDegree(node Node) (float64, bool) {
	return g.invAdj.GetOutDegree(node)
}
