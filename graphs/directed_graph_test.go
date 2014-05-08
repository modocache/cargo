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

	Describe(".RouteExists()", func() {
		Context("when a route exists between the start and end keys", func() {
			BeforeEach(func() {
				graph.AppendAdjacencyList(AdjacencyList{
					"A": Connections{{"A", 0}, {"B", 0}},
					"B": Connections{{"C", 0}, {"D", 0}},
					"C": Connections{{"D", 0}, {"E", 0}, {"F", 0}, {"G", 0}},
					"F": Connections{{"H", 0}, {"I", 0}},
					"I": Connections{{"J", 0}},
				})
			})
			It("returns true", func() {
				Expect(graph.RouteExists("A", "J")).To(BeTrue())
			})
		})

		Context("when a route does not exist between the start and end keys", func() {
			BeforeEach(func() {
				graph.AppendAdjacencyList(AdjacencyList{
					"A": Connections{{"A", 0}, {"B", 0}},
					"B": Connections{{"C", 0}, {"D", 0}},
					"C": Connections{{"D", 0}, {"E", 0}, {"G", 0}},
					"F": Connections{{"H", 0}, {"I", 0}},
					"I": Connections{{"J", 0}},
				})
			})
			It("returns false", func() {
				Expect(graph.RouteExists("A", "J")).To(BeFalse())
			})
		})
	})

	Describe("searching", func() {
		var visited []string
		var callback func(vertex *Vertex) bool

		BeforeEach(func() {
			graph.AppendAdjacencyList(AdjacencyList{
				"A": Connections{{"B", 0}, {"C", 0}},
				"C": Connections{{"D", 0}, {"F", 0}},
				"D": Connections{{"E", 0}},
				"F": Connections{{"G", 0}},
				"G": Connections{{"A", 0}},
			})

			visited = []string{}
			callback = func(vertex *Vertex) bool {
				visited = append(visited, vertex.Value.(string))
				return false
			}
		})

		Describe(".DepthFirstSearch()", func() {
			It("traverses each vertex, depth-first", func() {
				graph.DepthFirstSearch("A", callback)
				Expect(visited).To(Equal([]string{"A", "B", "C", "D", "E", "F", "G"}))
			})
		})

		Describe(".BreadthFirstSearch()", func() {
			It("traverses each vertex, breadth-first", func() {
				graph.BreadthFirstSearch("A", callback)
				Expect(visited).To(Equal([]string{"A", "B", "C", "D", "F", "E", "G"}))
			})
		})
	})
})
