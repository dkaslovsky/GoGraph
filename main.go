package main

import (
	"fmt"

	"github.com/dkaslovsky/GoGraph/graph"
	"github.com/dkaslovsky/GoGraph/io"
)

func main() {

	// undirected graph
	g := graph.NewGraph("Graph")
	err := io.LoadGraphFromFile("graph.golden", g)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	fmt.Print(g.Name)
	fmt.Println("\nAdjacency:")
	g.PrintAdj()
	fmt.Println()

	// directed graph
	dg := graph.NewDirGraph("DirGraph")
	err = io.LoadGraphFromFile("graph.golden", dg)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	fmt.Print(dg.Name)
	fmt.Println("\nOut adjacency:")
	dg.PrintOutAdj()
	fmt.Println("\nIn adjacency:")
	dg.PrintInAdj()
}
