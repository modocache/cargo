package trees_test

import (
	"github.com/modocache/cargo/comparators"
	. "github.com/modocache/cargo/trees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearchTree", func() {
	Describe("NewBinarySearchTree()", func() {
		It("returns a BinarySearchTree with the given value", func() {
			proverb := "Dying while young is a boon in old age."
			proverbTree := NewBinarySearchTree(proverb, comparators.StringLess)
			Expect(proverbTree.Value).To(Equal(proverb))

			number := 8675309
			numberTree := NewBinarySearchTree(number, comparators.IntLess)
			Expect(numberTree.Value).To(Equal(number))
		})
	})

	Describe("methods", func() {
		var tree *BinarySearchTree
		BeforeEach(func() {
			tree = NewBinarySearchTree(100, comparators.IntLess)
		})

		Describe(".Insert()", func() {
			var value interface{}

			Context("when the inserted value is less than the root value", func() {
				BeforeEach(func() { value = 50 })
				It("adds a left child tree", func() {
					tree.Insert(value)
					Expect(tree.Left.Value).To(Equal(50))
				})
			})

			Context("when the inserted value is not less than the root value", func() {
				BeforeEach(func() { value = 150 })
				It("adds a right child tree", func() {
					tree.Insert(value)
					Expect(tree.Right.Value).To(Equal(150))
				})
			})

			Describe("many values", func() {
				BeforeEach(func() {
					tree.Insert(200)
					tree.Insert(300)
					tree.Insert(250)
					tree.Insert(150)
					tree.Insert(275)
				})
				It("adds left and right trees unbalanced", func() {
					Expect(tree.Left).To(BeNil())
					Expect(tree.Right.Value).To(Equal(200))
					Expect(tree.Right.Left.Value).To(Equal(150))
					Expect(tree.Right.Right.Value).To(Equal(300))
					Expect(tree.Right.Right.Left.Value).To(Equal(250))
					Expect(tree.Right.Right.Left.Right.Value).To(Equal(275))
				})
			})
		})

		Describe(".Find()", func() {
			var value int
			BeforeEach(func() {
				tree.Insert(75)
				tree.Insert(80)
				tree.Insert(50)
				tree.Insert(25)
				tree.Insert(30)
			})

			Context("when the tree contains the value", func() {
				BeforeEach(func() { value = 30 })
				It("returns the child tree with that value", func() {
					child := tree.Find(value)
					Expect(child.Value).To(Equal(30))
				})
			})

			Context("when the tree does not contain the value", func() {
				BeforeEach(func() { value = 200 })
				It("returns nil", func() {
					Expect(tree.Find(value)).To(BeNil())
				})
			})
		})
	})
})
