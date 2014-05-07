package graphs

type DirectedGraph struct {
	vertices VertexMap
}

func (graph *DirectedGraph) Vertices() VertexMap {
	return graph.vertices
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{vertices: make(VertexMap)}
}

func (graph *DirectedGraph) Append(key interface{}) {
	appendVertex(graph, key)
}

func (graph *DirectedGraph) Connect(fromKey, toKey interface{}, weight int) {
	connectVertices(graph, fromKey, toKey, weight)
}

func (graph *DirectedGraph) DepthFirstSearch(startKey interface{}, callback GraphSearchCallback) {
	clearVisitedFlags(graph)
	depthFirstSearch(graph, startKey, callback)
}
