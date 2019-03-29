package graph

// DirGraph is an adjacency map representation of a directed graph
type DirGraph struct {
	Name   string
	outAdj dirAdj
	inAdj  dirAdj // inverted index of outAdj
}

// NewDirGraph creates a new directed graph
func NewDirGraph(name string) *DirGraph {
	return &DirGraph{
		Name:   name,
		outAdj: dirAdj{},
		inAdj:  dirAdj{},
	}
}

// AddEdge adds an edge from a node to another node with an optional weight that defaults to 1.0
func (dg *DirGraph) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}
	dg.outAdj.addEdge(from, to, wgt)
	dg.inAdj.addEdge(to, from, wgt)
}

// RemoveEdge removes an edge that exists from a node to another node
func (dg *DirGraph) RemoveEdge(from string, to string) {
	dg.outAdj.removeEdge(from, to)
	dg.inAdj.removeEdge(to, from)
}

// RemoveNode removes a node entirely from a DirGraph such that
// no edges exist between it an any other node
func (dg *DirGraph) RemoveNode(node string) {
	if nbrs, ok := dg.GetInNeighbors(node); ok {
		for _, n := range nbrs {
			dg.RemoveEdge(n, node)
		}
	}
	if nbrs, ok := dg.GetOutNeighbors(node); ok {
		for _, n := range nbrs {
			dg.RemoveEdge(node, n)
		}
	}
}

// PrintAdj displays a DirGraph's adjacency structure as a map of source nodes to target nodes
func (dg *DirGraph) PrintAdj() {
	dg.PrintOutAdj()
}

// PrintOutAdj displays a DirGraph's outgoing adjacency structure
func (dg *DirGraph) PrintOutAdj() {
	dg.outAdj.print()
}

// PrintInAdj displays a DirGraph's incoming adjacency structure
func (dg *DirGraph) PrintInAdj() {
	dg.inAdj.print()
}

// GetNodes gets a slice of all nodes in a DirGraph
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

// GetNeighbors gets a slice of nodes that have an edge to them from a specified node
func (dg *DirGraph) GetNeighbors(node string) (nbrs []string, found bool) {
	return dg.GetOutNeighbors(node)
}

// GetOutNeighbors gets a slice of nodes that have an edge to them from a specified node
func (dg *DirGraph) GetOutNeighbors(node string) (nbrs []string, found bool) {
	return dg.outAdj.getNeighbors(node)
}

// GetInNeighbors gets a slice of nodes that have an edge from them to a specified node
func (dg *DirGraph) GetInNeighbors(node string) (nbrs []string, found bool) {
	return dg.inAdj.getNeighbors(node)
}

// GetTotalDegree calculates the sum of weights of all edges from and to a node
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

// GetOutDegree calculates the sum of weights of all edges from a node
func (dg *DirGraph) GetOutDegree(node string) (deg float64, found bool) {
	nbrs, ok := dg.GetOutNeighbors(node)
	if !ok {
		return deg, false
	}
	for _, n := range nbrs {
		if w, ok := dg.GetEdgeWeight(node, n); ok {
			deg += w
		}
	}
	return deg, true
}

// GetInDegree calculates the sum of weights of all edges to a node
func (dg *DirGraph) GetInDegree(node string) (deg float64, found bool) {
	nbrs, ok := dg.GetInNeighbors(node)
	if !ok {
		return deg, false
	}
	for _, n := range nbrs {
		if w, ok := dg.GetEdgeWeight(n, node); ok {
			deg += w
		}
	}
	return deg, true
}

// HasEdge returns true if an edge exists from a node to another node, false otherwise
func (dg *DirGraph) HasEdge(from string, to string) bool {
	return dg.outAdj.hasEdge(from, to)
}

// GetEdgeWeight gets the weight of an edge from a node to another node if the edge exists
func (dg *DirGraph) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	return dg.outAdj.getEdgeWeight(from, to)
}
