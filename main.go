package main

import (
	"fmt"

	"github.com/dkaslovsky/GoGraph/graph"
	"github.com/dkaslovsky/GoGraph/io"
)

func main() {
	g := graph.NewDirGraph("myGraph")
	err := io.ReadFromFile("graph.txt", g)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	fmt.Println("\nOut adjacency:")
	g.PrintOutAdj()
	fmt.Println("\nIn adjacency:")
	g.PrintInAdj()
}
