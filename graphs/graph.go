package graphs

import (
	"fmt"
	"github.com/modocache/cargo/queues"
)

type visitedFlag int

const (
	unvisited visitedFlag = 1
	visiting  visitedFlag = 2
	visited   visitedFlag = 3
)

type Vertex struct {
	Edges []*Edge
	Value interface{}
	flag  visitedFlag
}

type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
}

func (edge *Edge) Reverse() *Edge {
	return &Edge{From: edge.To, To: edge.From, Weight: edge.Weight}
}

type VertexMap map[interface{}]*Vertex

type GraphSearchCallback func(vertex *Vertex) bool

func (vertices VertexMap) getOrPanic(key interface{}) *Vertex {
	vertex, exists := vertices[key]
	if !exists {
		message := fmt.Sprintf("Vertex", key, "is not included in the graph")
		panic(message)
	}

	return vertex
}

type Graph interface {
	Vertices() VertexMap
}

func appendVertex(graph Graph, key interface{}) {
	graph.Vertices()[key] = &Vertex{Value: key, flag: unvisited}
}

func connectVertices(graph Graph, fromKey, toKey interface{}, weight int) {
	from := graph.Vertices().getOrPanic(fromKey)
	to := graph.Vertices().getOrPanic(toKey)
	edge := &Edge{From: from, To: to, Weight: weight}
	from.Edges = append(from.Edges, edge)
}

func clearVisitedFlags(graph Graph) {
	for _, vertex := range graph.Vertices() {
		vertex.flag = unvisited
	}
}

func depthFirstSearch(graph Graph, startKey interface{}, callback GraphSearchCallback) {
	vertex := graph.Vertices().getOrPanic(startKey)
	vertex.flag = visiting
	if callback(vertex) {
		return
	}

	for _, edge := range vertex.Edges {
		if next := edge.To; next.flag == unvisited {
			depthFirstSearch(graph, edge.To.Value, callback)
		}
	}

	vertex.flag = visited
}

func breadthFirstSearch(graph Graph, queue *queues.Queue, callback GraphSearchCallback) {
	vertex := graph.Vertices().getOrPanic(queue.Pop())
	if callback(vertex) {
		return
	}
	vertex.flag = visited

	for _, edge := range vertex.Edges {
		if to := edge.To; to.flag == unvisited {
			to.flag = visiting
			queue.Push(to.Value)
		}
	}

	if !queue.IsEmpty() {
		breadthFirstSearch(graph, queue, callback)
	}
}
