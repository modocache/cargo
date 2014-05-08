package trees_test

import (
	"github.com/modocache/cargo/comparators"
	. "github.com/modocache/cargo/trees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearchable", func() {
	var searchable BinarySearchable

	Describe("Root()", func() {
		Context("when the searchable argument is nil", func() {
			BeforeEach(func() { searchable = nil })
			It("panics", func() {
				Expect(func() { Root(searchable) }).To(Panic())
			})
		})

		Context("when the searchable argument is an orphan", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree("Bruce Wayne", comparators.StringLess)
			})
			It("returns the orphan, since it's the root", func() {
				Expect(Root(searchable)).To(Equal(searchable))
			})
		})

		Context("when the searchable argument is a child", func() {
			var root BinarySearchable
			BeforeEach(func() {
				root = NewBinarySearchTree("Homer", comparators.StringLess)
				root.(*BinarySearchTree).InsertAll("Bart", "Lisa", "Maggie")
				searchable = root.Find("Maggie")
			})
			It("returns the orphan, since it's the root", func() {
				Expect(Root(searchable)).To(Equal(root))
			})
		})
	})

	Describe("Depth()", func() {
		Context("when the searchable is an orphan", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree("Dick Grayson", comparators.StringLess)
			})
			It("returns 0", func() {
				Expect(Depth(searchable)).To(Equal(0))
			})
		})

		Context("when the searchable is not an orphan", func() {
			BeforeEach(func() {
				root := NewBinarySearchTree("Bob", comparators.StringLess)
				root.InsertAll("Tina", "Gene", "Louise")
				searchable = root.Find("Louise")
			})
			It("returns the number of nodes to the root", func() {
				Expect(Depth(searchable)).To(Equal(3))
			})
		})
	})

	Describe("Height()", func() {
		Context("when the searchable is nil", func() {
			BeforeEach(func() { searchable = nil })
			It("returns 0", func() {
				Expect(Height(searchable)).To(Equal(0))
			})
		})

		Context("when the searchable is a leaf", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree("konoha", comparators.StringLess)
			})
			It("returns 0", func() {
				Expect(Height(searchable)).To(Equal(0))
			})
		})

		Context("when the searchable has subtrees", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree(0, comparators.IntLess)
				searchable.(*BinarySearchTree).InsertAll(-1, -3, -2, -4, 1, 2, 3)
			})
			It("returns the height of the searchable", func() {
				Expect(Height(searchable)).To(Equal(3))
			})
		})
	})

	Describe("IsBalanced()", func() {
		Context("when searchable is nil", func() {
			BeforeEach(func() { searchable = nil })
			It("returns true (it is balanced in its nothingness)", func() {
				Expect(IsBalanced(searchable)).To(BeTrue())
			})
		})

		Context("when searchable is a leaf", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree("House of Leaves", comparators.StringLess)
			})
			It("returns true (its two subtrees have equal heights of 0)", func() {
				Expect(IsBalanced(searchable)).To(BeTrue())
			})
		})

		Context("when searchable has two subtrees whose height differ by more than 1", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree(0, comparators.IntLess)
				searchable.(*BinarySearchTree).InsertAll(-10, -5, -20, 10, 20, 30, 40)
			})
			It("returns false", func() {
				Expect(IsBalanced(searchable)).To(BeFalse())
			})
		})

		Context("when searchable has two subtrees whose heights are within 1 of one another", func() {
			BeforeEach(func() {
				searchable = NewBinarySearchTree(0, comparators.IntLess)
				searchable.(*BinarySearchTree).InsertAll(-10, -20, -30, 10, 20, 30, 40)
			})
			It("returns true", func() {
				Expect(IsBalanced(searchable)).To(BeTrue())
			})
		})
	})
})
