package graphs_test

import (
	. "github.com/modocache/cargo/graphs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UndirectedGraph", func() {
	var graph *UndirectedGraph
	BeforeEach(func() {
		graph = NewUndirectedGraph()
	})

	Describe(".Connect()", func() {
		It("connects two vertices to one another", func() {
			graph.Append("both")
			graph.Append("ways")
			graph.Connect("both", "ways", 0)

			there := graph.Vertices()["both"].Edges[0]
			Expect(there.To.Value).To(Equal("ways"))
			back := graph.Vertices()["ways"].Edges[0]
			Expect(back.To.Value).To(Equal("both"))
		})
	})
})
