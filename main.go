package main

import (
	"fmt"
)

type Node struct {
	ID        string
	EdgesFrom []*Node
	EdgesTo   []*Node
}

type Edge struct {
	FromNode string
	ToNode   string
}

type Graph struct {
	ID    string
	Nodes map[string]*Node
}

func newGraph(id string) *Graph {
	return &Graph{
		ID:    id,
		Nodes: map[string]*Node{},
	}
}

func newNode(id string) *Node {
	return &Node{
		ID:        id,
		EdgesFrom: []*Node{},
		EdgesTo:   []*Node{},
	}
}

func newEdge(from, to *Node) *Edge {
	return &Edge{
		FromNode: from.ID,
		ToNode:   to.ID,
	}
}

func (g Graph) AddEdge(from, to *Node) {
	if from.ID == to.ID {
		return
	}
	g.addNode(from)
	g.addNode(to)
	from.EdgesTo = append(from.EdgesTo, to)
	to.EdgesFrom = append(to.EdgesFrom, from)
}

func (g Graph) addNode(node *Node) {
	_, ok := g.Nodes[node.ID]
	if !ok {
		g.Nodes[node.ID] = node
	}
}

func (g Graph) GetEdges() []*Edge {
	var edges []*Edge
	for _, node := range g.Nodes {
		for _, fromNode := range node.EdgesFrom {
			edges = append(edges, newEdge(fromNode, node))
		}
	}
	return edges
}

func (e Edge) print() {
	fmt.Println(e.FromNode, "-->", e.ToNode)
}

func (n Node) getOutDegree() int {
	return len(n.EdgesTo)
}

func (n Node) getInDegree() int {
	return len(n.EdgesFrom)
}

func main() {

	n1 := newNode("A")
	n2 := newNode("B")
	n3 := newNode("C")

	g := newGraph("myGraph")
	g.AddEdge(n1, n2)
	g.AddEdge(n3, n1)
	g.AddEdge(n1, n3)
	g.AddEdge(n2, n2)
	fmt.Println("Graph", g.ID, "has", len(g.Nodes), "nodes")

	fmt.Println("Node", n1.ID, "has in-degree", n1.getInDegree())
	fmt.Println("Node", n1.ID, "has out-degree", n1.getOutDegree())

	for _, node := range g.Nodes {
		fmt.Println("Node", node.ID, "has edges to...")
		for _, n := range node.EdgesTo {
			fmt.Println("\t", n.ID)
		}
		fmt.Println("Node", node.ID, "has edges from...")
		for _, n := range node.EdgesFrom {
			fmt.Println("\t", n.ID)
		}
	}

	for _, edge := range g.GetEdges() {
		edge.print()
	}

}
