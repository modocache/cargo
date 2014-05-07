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
			graph.Append("A")
			graph.Append("B")
			graph.Append("C")
			graph.Append("D")
			graph.Append("E")
			graph.Append("F")
			graph.Append("G")
			graph.Connect("A", "B", 0)
			graph.Connect("A", "C", 0)
			graph.Connect("C", "D", 0)
			graph.Connect("D", "E", 0)
			graph.Connect("C", "F", 0)
			graph.Connect("F", "G", 0)
			graph.Connect("G", "A", 0)

			keys := []string{}
			graph.DepthFirstSearch("A", func(vertex *Vertex) bool {
				keys = append(keys, vertex.Value.(string))
				return false
			})
			Expect(keys).To(Equal([]string{"A", "B", "C", "D", "E", "F", "G"}))
		})
	})
})
