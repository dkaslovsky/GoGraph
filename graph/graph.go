package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type adjacency map[string]map[string]float64

type DirGraph struct {
	Name   string
	outAdj adjacency
	inAdj  adjacency //inverted index of outAdj
}

func NewDirGraph(name string) *DirGraph {
	return &DirGraph{
		Name:   name,
		outAdj: adjacency{},
		inAdj:  adjacency{},
	}
}

func (dg *DirGraph) PrintAdj() {
	dg.PrintOutAdj()
}

func (dg *DirGraph) PrintOutAdj() {
	for node := range dg.outAdj {
		fmt.Printf("%s:\n", node)
		adj := dg.outAdj[node]
		for n := range adj {
			fmt.Printf("  %s: %f\n", n, adj[n])
		}
	}
}

func (dg *DirGraph) PrintInAdj() {
	for node := range dg.inAdj {
		fmt.Printf("%s:\n", node)
		adj := dg.inAdj[node]
		for n := range adj {
			fmt.Printf("  %s: %f\n", n, adj[n])
		}
	}
}

func (dg *DirGraph) FromFile(filepath string) error {
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
			dg.AddEdge(from, to)
		} else {
			weight, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return err
			}
			dg.AddEdge(from, to, weight)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (dg *DirGraph) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}

	toNodes, ok := dg.outAdj[from]
	if !ok {
		dg.outAdj[from] = map[string]float64{to: wgt}
	} else {
		toNodes[to] = wgt
	}

	fromNodes, ok := dg.inAdj[to]
	if !ok {
		dg.inAdj[to] = map[string]float64{from: wgt}
	} else {
		fromNodes[from] = wgt
	}
}

func (dg *DirGraph) GetNodes() (nodes []string) {
	for node := range dg.outAdj {
		nodes = append(nodes, node)
	}
	return nodes
}

func (dg *DirGraph) GetNeighbors(node string) (nbrs []string, found bool) {
	return dg.GetOutNeighbors(node)
}

func (dg *DirGraph) GetOutNeighbors(node string) (nbrs []string, found bool) {
	adj, ok := dg.outAdj[node]
	if !ok {
		return nbrs, false
	}
	for n := range adj {
		nbrs = append(nbrs, n)
	}
	return nbrs, true
}

func (dg *DirGraph) GetInNeighbors(node string) (nbrs []string, found bool) {
	adj, ok := dg.inAdj[node]
	if !ok {
		return nbrs, false
	}
	for n := range adj {
		nbrs = append(nbrs, n)
	}
	return nbrs, true
}

func (dg *DirGraph) GetDegree(node string) (deg float64, found bool) {
	outDeg, ok := dg.GetOutDegree(node)
	if !ok {
		return deg, false
	}
	inDeg, ok := dg.GetInDegree(node)
	if !ok {
		return deg, false
	}
	return inDeg + outDeg, true
}

func (dg *DirGraph) GetOutDegree(node string) (deg float64, found bool) {
	adj, ok := dg.outAdj[node]
	if !ok {
		return deg, false
	}
	for n := range adj {
		deg += adj[n]
	}
	return deg, true
}

func (dg *DirGraph) GetInDegree(node string) (deg float64, found bool) {
	adj, ok := dg.inAdj[node]
	if !ok {
		return deg, false
	}
	for n := range adj {
		deg += adj[n]
	}
	return deg, true
}
