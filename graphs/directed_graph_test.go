package graphs_test

import (
	. "github.com/modocache/cargo/graphs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DirectedGraph", func() {
	var graph *DirectedGraph
	BeforeEach(func() { graph = NewDirectedGraph() })
	Describe(".Append()", func() {
		It("adds a vertex to the graph", func() {
			graph.Append("A")
			Expect(graph.Vertices()["A"]).ToNot(BeNil())
		})
	})

	Describe(".Connect()", func() {
		It("connects two vertices in the graph", func() {
			graph.Append("goto")
			graph.Append("fail")
			graph.Connect("goto", "fail", 0)

			edge := graph.Vertices()["goto"].Edges[0]
			Expect(edge.From.Value).To(Equal("goto"))
			Expect(edge.To.Value).To(Equal("fail"))
		})
	})

	Describe(".DepthFirstSearch()", func() {
		It("traverses each vertex", func() {
			graph.AppendAdjacencyList(AdjacencyList{
				"A": Connections{{"B", 0}, {"C", 0}},
				"C": Connections{{"D", 0}, {"F", 0}},
				"D": Connections{{"E", 0}},
				"F": Connections{{"G", 0}},
				"G": Connections{{"A", 0}},
			})

			keys := []string{}
			graph.DepthFirstSearch("A", func(vertex *Vertex) bool {
				keys = append(keys, vertex.Value.(string))
				return false
			})
			Expect(keys).To(Equal([]string{"A", "B", "C", "D", "E", "F", "G"}))
		})
	})
})
