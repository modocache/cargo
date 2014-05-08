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
			tree.InsertAll(14, 2, 1, 7, 15, 5, 8)
		})

		It("balances left-heavy trees", func() {
			tree.Insert(4)
			root := Root(tree)
			Expect(IsBalanced(root)).To(BeTrue())
		})

		It("balances right-heavy trees", func() {
			tree.Insert(9)
			root := Root(tree)
			Expect(IsBalanced(root)).To(BeTrue())
		})
	})
})
