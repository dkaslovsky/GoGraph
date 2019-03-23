package main

import (
	"fmt"

	"github.com/dkaslovsky/GoGraph/graph"
)

func printNbrs(g *graph.DirGraph, node string) {
	nbrs, ok := g.GetNeighbors(node)
	if !ok {
		fmt.Printf("%s not found in %s\n", node, g.Name)
		return
	}
	fmt.Printf("%s's neighbors: %v\n", node, nbrs)
}

func printOutDegree(g *graph.DirGraph, node string) {
	d, ok := g.GetOutDegree(node)
	if !ok {
		fmt.Printf("%s not found in %s\n", node, g.Name)
		return
	}
	fmt.Printf("%s's out degree: %f\n", node, d)
}

func printInDegree(g *graph.DirGraph, node string) {
	d, ok := g.GetInDegree(node)
	if !ok {
		fmt.Printf("%s not found in %s\n", node, g.Name)
		return
	}
	fmt.Printf("%s's in degree: %f\n", node, d)
}

func main() {
	g := graph.NewDirGraph("myGraph")
	err := g.FromFile("graph.txt")
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	fmt.Printf("Nodes in %s: %v\n", g.Name, g.GetNodes())
	g.PrintAdj()
	for _, node := range []string{"a", "c", "x"} {
		printNbrs(g, node)
		printOutDegree(g, node)
		printInDegree(g, node)
	}

	fmt.Println()

	n1 := "c"
	n2 := "d"
	fmt.Printf("Removing edge (%s, %s)", n1, n2)
	g.RemoveEdge(n1, n2)
	fmt.Printf("Nodes in %s: %v\n", g.Name, g.GetNodes())
	g.PrintAdj()
	for _, node := range []string{"a", "c", "x"} {
		printNbrs(g, node)
		printOutDegree(g, node)
		printInDegree(g, node)
	}
}
