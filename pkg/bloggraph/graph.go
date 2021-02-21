package bloggraph

// Graph represents the complete graph spanning the whole blog having basically
// an entrypoint with every single blog post.
type Graph struct {
	nodes []*Node
	edges []*Edge
}

// NewGraph generates a new empty Graph instance.
func NewGraph() *Graph {
	return &Graph{
		nodes: make([]*Node, 0, 10),
		edges: make([]*Edge, 0, 10),
	}
}

func (g *Graph) NumEdges() int {
	return len(g.edges)
}

func (g *Graph) NumNodes() int {
	return len(g.nodes)
}

type WalkerFunc func(graph *Graph, node *Node, depth int)

// WalkDown the graph.
func (g *Graph) WalkDown(startID string, kind string, degree int, walker WalkerFunc) {
	for _, e := range g.edges {
		if e.Source.ContentID == startID && e.Kind == kind {
			walker(g, e.Target, degree)
			g.WalkDown(e.Target.ContentID, kind, degree+1, walker)
		}
	}
}

// WalkUp the graph.
func (g *Graph) WalkUp(startID string, kind string, depth int, walker WalkerFunc) {
	for _, e := range g.edges {
		if e.Target.ContentID == startID && e.Kind == kind {
			walker(g, e.Source, depth)
			g.WalkUp(e.Source.ContentID, kind, depth+1, walker)
		}
	}
}

func (g *Graph) GetOrCreateNode(n Node) *Node {
	for _, existing := range g.nodes {
		if existing.ContentID == n.ContentID {
			return existing
		}
	}
	g.nodes = append(g.nodes, &n)
	return &n
}

func (g *Graph) CreateEdge(source *Node, target *Node, kind string) *Edge {
	for _, e := range g.edges {
		if e.Kind == kind && e.Source == source && e.Target == target {
			return e
		}
	}
	e := Edge{
		Source: source,
		Target: target,
		Kind:   kind,
	}
	g.edges = append(g.edges, &e)
	return &e
}

// Node represents a single blog post inside the overall graph.
type Node struct {
	ContentID string
	Title     string
}

// Edge combines two nodes.
type Edge struct {
	Source *Node
	Target *Node
	Kind   string
}
