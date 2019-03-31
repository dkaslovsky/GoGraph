package main

import (
	"fmt"
	"os"

	"github.com/dkaslovsky/GoGraph/graph"
)

func main() {

	filepath := "graph.golden"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	defer file.Close()

	// undirected graph
	g, err := graph.NewGraph("Graph", file)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Print(g.Name)
	fmt.Println("\nAdjacency:")
	g.Print()
	fmt.Println()

	// // directed graph
	// dg, err := graph.NewDirGraph("DirGraph", file)
	// if err != nil {
	// 	fmt.Printf("Error %v\n", err)
	// }

	// fmt.Print(dg.Name)
	// fmt.Println("\nOut adjacency:")
	// dg.Print()
	// fmt.Println("\nIn adjacency:")
	// dg.PrintInv()
}
