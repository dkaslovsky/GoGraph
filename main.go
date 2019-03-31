package main

import (
	"fmt"
	"os"

	"github.com/dkaslovsky/GoGraph/graph"
)

func main() {

	filepath := "graph.txt"

	// undirected graph
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	g, err := graph.NewGraph("Graph", file)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Print(g.Name)
	fmt.Println("\nAdjacency:")
	g.Print()
	fmt.Println()

	// directed graph
	file, err = os.Open(filepath)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	dg, err := graph.NewDirGraph("DirGraph", file)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Print(dg.Name)
	fmt.Println("\nOut adjacency:")
	dg.Print()
	fmt.Println("\nIn adjacency:")
	dg.PrintInv()
}
