package graph

type DirGraph struct {
	Name   string
	outAdj dirAdj
	inAdj  dirAdj // inverted index of outAdj
}

func NewDirGraph(name string) *DirGraph {
	return &DirGraph{
		Name:   name,
		outAdj: dirAdj{},
		inAdj:  dirAdj{},
	}
}

func (dg *DirGraph) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}
	dg.outAdj.addEdge(from, to, wgt)
	dg.inAdj.addEdge(to, from, wgt)
}

func (dg *DirGraph) RemoveEdge(from string, to string) {
	dg.outAdj.removeEdge(from, to)
	dg.inAdj.removeEdge(to, from)
}

func (dg *DirGraph) RemoveNode(node string) {
	if nbrs, ok := dg.GetInNeighbors(node); ok {
		for _, n := range nbrs {
			dg.outAdj.removeEdge(n, node)
		}
	}
	if nbrs, ok := dg.GetOutNeighbors(node); ok {
		for _, n := range nbrs {
			dg.inAdj.removeEdge(n, node)
		}
	}
	delete(dg.outAdj, node)
	delete(dg.inAdj, node)
}

func (dg *DirGraph) PrintAdj() {
	dg.PrintOutAdj()
}

func (dg *DirGraph) PrintOutAdj() {
	dg.outAdj.print()
}

func (dg *DirGraph) PrintInAdj() {
	dg.inAdj.print()
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
	return dg.outAdj.getNeighbors(node)
}

func (dg *DirGraph) GetInNeighbors(node string) (nbrs []string, found bool) {
	return dg.inAdj.getNeighbors(node)
}

func (dg *DirGraph) GetTotalDegree(node string) (deg float64, found bool) {
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

func (dg *DirGraph) HasEdge(from string, to string) bool {
	return dg.outAdj.hasEdge(from, to)
}

func (dg *DirGraph) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	return dg.outAdj.getEdgeWeight(from, to)
}
