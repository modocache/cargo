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

type Connection struct {
	Key    interface{}
	Weight int
}
type Connections []Connection
type AdjacencyList map[interface{}]Connections

func (graph *DirectedGraph) AppendAdjacencyList(adjacencies AdjacencyList) {
	for key, connections := range adjacencies {
		graph.Append(key)
		for _, connection := range connections {
			graph.Append(connection.Key)
		}
	}

	for key, connections := range adjacencies {
		for _, connection := range connections {
			graph.Connect(key, connection.Key, connection.Weight)
		}
	}
}

func (graph *DirectedGraph) DepthFirstSearch(startKey interface{}, callback GraphSearchCallback) {
	clearVisitedFlags(graph)
	depthFirstSearch(graph, startKey, callback)
}
