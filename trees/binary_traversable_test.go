package trees_test

import (
	"github.com/modocache/cargo"
	"github.com/modocache/cargo/comparators"
	. "github.com/modocache/cargo/trees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinaryTraversable", func() {
	var traversable BinaryTraversable

	Describe("Root()", func() {
		Context("when the traversable argument is nil", func() {
			BeforeEach(func() { traversable = nil })
			It("panics", func() {
				Expect(func() { Root(traversable) }).To(Panic())
			})
		})

		Context("when the traversable argument is an orphan", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree("Bruce Wayne", comparators.StringLess)
			})
			It("returns the orphan, since it's the root", func() {
				Expect(Root(traversable)).To(Equal(traversable))
			})
		})

		Context("when the traversable argument is a child", func() {
			var root BinaryTraversable
			BeforeEach(func() {
				root = NewBinarySearchTree("Homer", comparators.StringLess)
				root.(*BinarySearchTree).InsertAll("Bart", "Lisa", "Maggie")
				traversable = root.(*BinarySearchTree).Find("Maggie")
			})
			It("returns the orphan, since it's the root", func() {
				Expect(Root(traversable)).To(Equal(root))
			})
		})
	})

	Describe("Depth()", func() {
		Context("when the traversable is an orphan", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree("Dick Grayson", comparators.StringLess)
			})
			It("returns 0", func() {
				Expect(Depth(traversable)).To(Equal(0))
			})
		})

		Context("when the traversable is not an orphan", func() {
			BeforeEach(func() {
				root := NewBinarySearchTree("Bob", comparators.StringLess)
				root.InsertAll("Tina", "Gene", "Louise")
				traversable = root.Find("Louise")
			})
			It("returns the number of nodes to the root", func() {
				Expect(Depth(traversable)).To(Equal(3))
			})
		})
	})

	Describe("Height()", func() {
		Context("when the traversable is nil", func() {
			BeforeEach(func() { traversable = nil })
			It("returns 0", func() {
				Expect(Height(traversable)).To(Equal(0))
			})
		})

		Context("when the traversable is a leaf", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree("konoha", comparators.StringLess)
			})
			It("returns 0", func() {
				Expect(Height(traversable)).To(Equal(0))
			})
		})

		Context("when the traversable has subtrees", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree(0, comparators.IntLess)
				traversable.(*BinarySearchTree).InsertAll(-1, -3, -2, -4, 1, 2, 3)
			})
			It("returns the height of the traversable", func() {
				Expect(Height(traversable)).To(Equal(3))
			})
		})
	})

	Describe("IsBalanced()", func() {
		Context("when traversable is nil", func() {
			BeforeEach(func() { traversable = nil })
			It("returns true (it is balanced in its nothingness)", func() {
				Expect(IsBalanced(traversable)).To(BeTrue())
			})
		})

		Context("when traversable is a leaf", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree("House of Leaves", comparators.StringLess)
			})
			It("returns true (its two subtrees have equal heights of 0)", func() {
				Expect(IsBalanced(traversable)).To(BeTrue())
			})
		})

		Context("when traversable has two subtrees whose height differ by more than 1", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree(0, comparators.IntLess)
				traversable.(*BinarySearchTree).InsertAll(-10, -5, -20, 10, 20, 30, 40)
			})
			It("returns false", func() {
				Expect(IsBalanced(traversable)).To(BeFalse())
			})
		})

		Context("when traversable has two subtrees whose heights are within 1 of one another", func() {
			BeforeEach(func() {
				traversable = NewBinarySearchTree(0, comparators.IntLess)
				traversable.(*BinarySearchTree).InsertAll(-10, -20, -30, 10, 20, 30, 40)
			})
			It("returns true", func() {
				Expect(IsBalanced(traversable)).To(BeTrue())
			})
		})
	})

	Describe(".BreadthFirstSearch()", func() {
		var callback TraversalCallback

		Context("when the traversable is nil", func() {
			var executed bool
			BeforeEach(func() {
				traversable = nil
				executed = false
				callback = func(traversable BinaryTraversable) bool {
					executed = true
					return executed
				}
			})
			It("never executes the callback", func() {
				BreadthFirstSearch(traversable, callback)
				Expect(traversable).To(BeNil())
			})
		})

		Context("when the traversable is not nil", func() {
			var visited []int
			BeforeEach(func() {
				traversable = NewBinarySearchTree(100, comparators.IntLess)
				traversable.(*BinarySearchTree).InsertAll(0, 200, -50, 50, 150, 250)

				visited = []int{}
				callback = func(traversable BinaryTraversable) bool {
					visited = append(visited, traversable.Value().(int))
					return false
				}
			})

			It("traverses the subtrees breadth-first", func() {
				BreadthFirstSearch(traversable, callback)
				Expect(visited).To(Equal([]int{100, 0, 200, -50, 50, 150, 250}))
			})
		})
	})

	Describe(".DepthFirstSearch()", func() {
		var order cargo.TraversalOrder
		var callback TraversalCallback

		Context("when the traversable is nil", func() {
			order = cargo.PreOrder
			var executed bool
			BeforeEach(func() {
				traversable = nil
				executed = false
				callback = func(traversable BinaryTraversable) bool {
					executed = true
					return executed
				}
			})
			It("never executes the callback", func() {
				DepthFirstSearch(traversable, order, callback)
				Expect(traversable).To(BeNil())
			})
		})

		Context("when the traversable is not nil", func() {
			var visited []int
			BeforeEach(func() {
				traversable = NewBinarySearchTree(100, comparators.IntLess)
				traversable.(*BinarySearchTree).InsertAll(0, 200, -50, 50, 150, 250)

				visited = []int{}
				callback = func(traversable BinaryTraversable) bool {
					visited = append(visited, traversable.Value().(int))
					return false
				}
			})

			Context("using pre-order", func() {
				BeforeEach(func() { order = cargo.PreOrder })
				It("traverses the subtrees depth-first using pre-order", func() {
					DepthFirstSearch(traversable, order, callback)
					Expect(visited).To(Equal([]int{100, 0, -50, 50, 200, 150, 250}))
				})
			})

			Context("using in-order", func() {
				BeforeEach(func() { order = cargo.InOrder })
				It("traverses the subtrees depth-first using in-order", func() {
					DepthFirstSearch(traversable, order, callback)
					Expect(visited).To(Equal([]int{-50, 0, 50, 100, 150, 200, 250}))
				})
			})

			Context("using post-order", func() {
				BeforeEach(func() { order = cargo.PostOrder })
				It("traverses the subtrees depth-first using post-order", func() {
					DepthFirstSearch(traversable, order, callback)
					Expect(visited).To(Equal([]int{-50, 50, 0, 150, 250, 200, 100}))
				})
			})
		})
	})

	Describe("IsBinarySearchTree()", func() {
		var comparator comparators.Less
		BeforeEach(func() {
			comparator = comparators.IntLess
			traversable = NewBinarySearchTree(100, comparator)
			traversable.(*BinarySearchTree).InsertAll(50, 150, 25, 75, 125, 175)
		})

		Context("when left < current < right", func() {
			It("returns true", func() {
				Expect(IsBinarySearchTree(traversable, comparator)).To(BeTrue())
			})
		})

		Context("when the left < current < right condition is not maintained", func() {
			BeforeEach(func() {
				child := traversable.(*BinarySearchTree).Find(75)
				child.SetLeft(NewBinarySearchTree(100, comparator))
			})
			It("returns false", func() {
				Expect(IsBinarySearchTree(traversable, comparator)).To(BeFalse())
			})
		})
	})
})
