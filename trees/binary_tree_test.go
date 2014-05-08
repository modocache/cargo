package trees_test

import (
	. "github.com/modocache/cargo/trees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trees/BinaryTree", func() {
	var tree *BinaryTree
	BeforeEach(func() {
		tree = NewBinaryTree("tree")
	})

	Describe(".SetParent()", func() {
		It("sets the parent of the tree", func() {
			Expect(tree.Parent()).To(BeNil())
			tree.SetParent(NewBinaryTree("parent"))
			Expect(tree.Parent().Value()).To(Equal("parent"))
		})
	})

	Describe(".SetLeft()", func() {
		It("sets the left child of the tree", func() {
			Expect(tree.Left()).To(BeNil())
			tree.SetLeft(NewBinaryTree("left"))
			Expect(tree.Left().Value()).To(Equal("left"))
		})
	})

	Describe(".SetRight()", func() {
		It("sets the right child of the tree", func() {
			Expect(tree.Right()).To(BeNil())
			tree.SetRight(NewBinaryTree("right"))
			Expect(tree.Right().Value()).To(Equal("right"))
		})
	})
})
