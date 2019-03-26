package graph

type Graph struct {
	Name string
	adj  dirAdj
}

func NewGraph(name string) *Graph {
	return &Graph{
		Name: name,
		adj:  dirAdj{},
	}
}

func (g *Graph) AddEdge(from string, to string, weight ...float64) {
	wgt := 1.0
	if len(weight) > 0 {
		wgt = weight[0]
	}
	g.adj.addEdge(from, to, wgt)
	g.adj.addEdge(to, from, wgt)
}

func (g *Graph) RemoveEdge(from string, to string) {
	g.adj.removeEdge(from, to)
	g.adj.removeEdge(to, from)
}

func (g *Graph) RemoveNode(node string) {
	if nbrs, ok := g.GetNeighbors(node); ok {
		for _, n := range nbrs {
			g.RemoveEdge(node, n)
		}
	}
}

func (g *Graph) PrintAdj() {
	g.adj.print()
}

func (g *Graph) GetNodes() (nodes []string) {
	for node := range g.adj {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *Graph) GetNeighbors(node string) (nbrs []string, found bool) {
	return g.adj.getNeighbors(node)
}

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

func (g *Graph) HasEdge(from string, to string) bool {
	return g.adj.hasEdge(from, to)
}

func (g *Graph) GetEdgeWeight(from string, to string) (weight float64, found bool) {
	return g.adj.getEdgeWeight(from, to)
}
