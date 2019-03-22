package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type graph map[string]map[string]float64

func (g graph) fromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		from := parts[0]
		to := parts[1]
		if from == "" || to == "" {
			continue
		}

		if len(parts) == 2 {
			g.addEdge(from, to)
		} else {
			weight, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return err
			}
			g.addEdge(from, to, weight)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (g graph) addEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}

	toNodes, ok := g[from]
	if !ok {
		g[from] = map[string]float64{to: wgt}
		return
	}
	_, ok = toNodes[to]
	if ok {
		// Should return error
		fmt.Printf("Error: edge (%s, %s) already in graph", from, to)
		return
	}
	toNodes[to] = wgt
}

func (g graph) getNodes() (nodes []string) {
	for node := range g {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g graph) getNeighbors(node string) (nodes []string) {
	nbrs, ok := g[node]
	if !ok {
		// Error if node not in graph?
		return nodes
	}
	for n := range nbrs {
		nodes = append(nodes, n)
	}
	return nodes
}

func (g graph) getDegree(node string) (degree float64) {
	nbrs, ok := g[node]
	if !ok {
		// Error if node not in graph?
		return degree
	}
	for n := range nbrs {
		degree += nbrs[n]
	}
	return degree
}

func main() {
	g := graph{}
	fmt.Printf("Nodes in graph: %v\n", g.getNodes())

	err := g.fromFile("graph.txt")
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	fmt.Printf("Graph: %v\n", g)
	fmt.Printf("Nodes in graph: %v\n", g.getNodes())

	fmt.Printf("a's neighbors: %v\n", g.getNeighbors("a"))
	fmt.Printf("x's neighbors: %v\n", g.getNeighbors("x"))

	fmt.Printf("a's degree: %v\n", g.getDegree("a"))
	fmt.Printf("x's degree: %v\n", g.getDegree("x"))
}
