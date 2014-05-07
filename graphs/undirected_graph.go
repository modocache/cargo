package graphs

type UndirectedGraph struct {
	*DirectedGraph
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{NewDirectedGraph()}
}

func (graph *UndirectedGraph) Connect(fromKey, toKey interface{}, weight int) {
	connectVertices(graph, fromKey, toKey, weight)
	connectVertices(graph, toKey, fromKey, weight)
}
