package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type DirGraph struct {
	Name   string
	outAdj DirAdj
	inAdj  DirAdj // inverted index of outAdj
}

func NewDirGraph(name string) *DirGraph {
	return &DirGraph{
		Name:   name,
		outAdj: DirAdj{},
		inAdj:  DirAdj{},
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
	dg.outAdj.AddEdge(from, to, wgt)
	dg.inAdj.AddEdge(to, from, wgt)
}

func (dg *DirGraph) RemoveEdge(from string, to string) {
	dg.outAdj.RemoveEdge(from, to)
	dg.inAdj.RemoveEdge(to, from)
}

func (dg *DirGraph) RemoveNode(node string) {
	dg.outAdj.RemoveNode(node)
	dg.inAdj.RemoveNode(node)
}

func (dg *DirGraph) PrintAdj() {
	dg.PrintOutAdj()
}

func (dg *DirGraph) PrintOutAdj() {
	dg.outAdj.Print()
}

func (dg *DirGraph) PrintInAdj() {
	dg.inAdj.Print()
}

func (dg *DirGraph) GetNodes() (nodes []string) {
	set := map[string]struct{}{}
	for node := range dg.outAdj {
		if _, ok := set[node]; !ok {
			set[node] = struct{}{}
			nodes = append(nodes, node)
		}
	}
	for node := range dg.inAdj {
		if _, ok := set[node]; !ok {
			set[node] = struct{}{}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (dg *DirGraph) GetNeighbors(node string) (nbrs []string, found bool) {
	return dg.GetOutNeighbors(node)
}

func (dg *DirGraph) GetOutNeighbors(node string) (nbrs []string, found bool) {
	return dg.outAdj.GetNeighbors(node)
}

func (dg *DirGraph) GetInNeighbors(node string) (nbrs []string, found bool) {
	return dg.inAdj.GetNeighbors(node)
}

// func (dg *DirGraph) GetDegree(node string) (deg float64, found bool) {
// 	outDeg, ok := dg.GetOutDegree(node)
// 	if !ok {
// 		return deg, false
// 	}
// 	inDeg, ok := dg.GetInDegree(node)
// 	if !ok {
// 		return deg, false
// 	}
// 	return inDeg + outDeg, true
// }

// func (dg *DirGraph) GetOutDegree(node string) (deg float64, found bool) {
// 	adj, ok := dg.outAdj[node]
// 	if !ok {
// 		return deg, false
// 	}
// 	for n := range adj {
// 		deg += adj[n]
// 	}
// 	return deg, true
// }

// func (dg *DirGraph) GetInDegree(node string) (deg float64, found bool) {
// 	adj, ok := dg.inAdj[node]
// 	if !ok {
// 		return deg, false
// 	}
// 	for n := range adj {
// 		deg += adj[n]
// 	}
// 	return deg, true
// }

func (dg *DirGraph) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	return dg.outAdj.GetEdgeWeight(from, to)
}
