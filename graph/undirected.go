package graph

// Graph is a symmetric adjacency map representation of an undirected graph
type Graph struct {
	Name string
	adj  dirAdj
}

// NewGraph creates a new undirected graph
func NewGraph(name string) *Graph {
	return &Graph{
		Name: name,
		adj:  dirAdj{},
	}
}

// AddEdge adds an edge between two nodes with an optional weight that defaults to 1.0
func (g *Graph) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}
	g.adj.addEdge(from, to, wgt)
	g.adj.addEdge(to, from, wgt)
}

// RemoveEdge removes an edge between two nodes
func (g *Graph) RemoveEdge(from string, to string) {
	g.adj.removeEdge(from, to)
	g.adj.removeEdge(to, from)
}

// RemoveNode removes a node entirely from a Graph such that
// no edges exist between it an any other node
func (g *Graph) RemoveNode(node string) {
	if nbrs, ok := g.GetNeighbors(node); ok {
		for _, n := range nbrs {
			g.RemoveEdge(node, n)
		}
	}
}

// PrintAdj displays a Graph's adjacency structure
func (g *Graph) PrintAdj() {
	g.adj.print()
}

// GetNodes gets a slice of all nodes in a Graph
func (g *Graph) GetNodes() (nodes []string) {
	for node := range g.adj {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetNeighbors gets a slice of nodes that are adjacent to a specified node
func (g *Graph) GetNeighbors(node string) (nbrs []string, found bool) {
	return g.adj.getNeighbors(node)
}

// GetDegree calculates the sum of weights of all edges of a node
func (g *Graph) GetDegree(node string) (deg float64, found bool) {
	nbrs, ok := g.GetNeighbors(node)
	if !ok {
		return deg, false
	}
	for _, n := range nbrs {
		if w, ok := g.GetEdgeWeight(node, n); ok {
			deg += w
		}
	}
	return deg, true
}

// HasEdge returns true if an edge exists between two nodes, false otherwise
func (g *Graph) HasEdge(from string, to string) bool {
	return g.adj.hasEdge(from, to)
}

// GetEdgeWeight gets the weight of an edge between two nodes if the edge exists
func (g *Graph) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	return g.adj.getEdgeWeight(from, to)
}
