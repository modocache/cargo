package trees_test

import (
	"github.com/modocache/cargo/comparators"
	. "github.com/modocache/cargo/trees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RedBlackTree", func() {
	var tree *RedBlackTree
	BeforeEach(func() {
		tree = NewRedBlackTree(11, comparators.IntLess)
	})

	Describe(".Insert()", func() {
		BeforeEach(func() {
			tree.Insert(14)
			tree.Insert(2)
			tree.Insert(1)
			tree.Insert(7)
			tree.Insert(15)
			tree.Insert(5)
			tree.Insert(8)
		})

		It("balances left-heavy trees", func() {
			tree.Insert(4)

			root := Root(tree)
			Expect(root.Value()).To(Equal(7))
			Expect(root.Left().Value()).To(Equal(2))
			Expect(root.Left().Left().Value()).To(Equal(1))
			Expect(root.Left().Right().Value()).To(Equal(5))
			Expect(root.Left().Right().Left().Value()).To(Equal(4))
			Expect(root.Right().Value()).To(Equal(11))
			Expect(root.Right().Left().Value()).To(Equal(8))
			Expect(root.Right().Right().Value()).To(Equal(14))
			Expect(root.Right().Right().Right().Value()).To(Equal(15))
		})

		It("balances right-heavy trees", func() {
			tree.Insert(9)

			root := Root(tree)
			Expect(root.Value()).To(Equal(7))
			Expect(root.Left().Value()).To(Equal(2))
			Expect(root.Left().Left().Value()).To(Equal(1))
			Expect(root.Left().Right().Value()).To(Equal(5))
			Expect(root.Right().Value()).To(Equal(11))
			Expect(root.Right().Left().Value()).To(Equal(8))
			Expect(root.Right().Left().Right().Value()).To(Equal(9))
			Expect(root.Right().Right().Value()).To(Equal(14))
			Expect(root.Right().Right().Right().Value()).To(Equal(15))
		})
	})
})
