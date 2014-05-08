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
				root.Insert("Bart")
				root.Insert("Lisa")
				searchable = root.Insert("Maggie")
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
				root.Insert("Tina")
				root.Insert("Gene")
				root.Insert("Louise")
				searchable = root.Find("Louise")
			})
			It("returns the number of nodes to the root", func() {
				Expect(Depth(searchable)).To(Equal(3))
			})
		})
	})
})
