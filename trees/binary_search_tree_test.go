package trees_test

import (
	"github.com/modocache/cargo/comparators"
	. "github.com/modocache/cargo/trees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearchTree", func() {
	var tree *BinarySearchTree
	var value interface{}

	Describe(".Insert()", func() {
		BeforeEach(func() { tree = NewBinarySearchTree(100, comparators.IntLess) })
		Context("when the inserted value is less than the root value", func() {
			BeforeEach(func() { value = 50 })
			It("adds a left child tree", func() {
				tree.Insert(value)
				Expect(tree.Left().Value()).To(Equal(50))
			})
		})

		Context("when the inserted value is not less than the root value", func() {
			BeforeEach(func() { value = 150 })
			It("adds a right child tree", func() {
				tree.Insert(value)
				Expect(tree.Right().Value()).To(Equal(150))
			})
		})

		Describe("when many values are inserted", func() {
			BeforeEach(func() {
				tree.Insert(200)
				tree.Insert(300)
				tree.Insert(250)
				tree.Insert(150)
				tree.Insert(275)
			})
			It("adds left and right leaves", func() {
				Expect(tree.Left()).To(BeNil())
				Expect(tree.Right().Value()).To(Equal(200))
				Expect(tree.Right().Left().Value()).To(Equal(150))
				Expect(tree.Right().Right().Value()).To(Equal(300))
				Expect(tree.Right().Right().Left().Value()).To(Equal(250))
				Expect(tree.Right().Right().Left().Right().Value()).To(Equal(275))
			})
		})
	})

	Describe(".Find()", func() {
		BeforeEach(func() {
			tree = NewBinarySearchTree("Terra", comparators.StringLess)
			tree.Insert("Locke")
			tree.Insert("Edgar")
			tree.Insert("Sabin")
			tree.Insert("Cyan")
			tree.Insert("Strago")
		})

		Context("when the tree contains the value", func() {
			BeforeEach(func() { value = "Strago" })
			It("returns the child tree with that value", func() {
				child := tree.Find(value)
				Expect(child.Value()).To(Equal("Strago"))
			})
		})

		Context("when the tree does not contain the value", func() {
			BeforeEach(func() { value = "Kefka" })
			It("returns nil", func() {
				Expect(tree.Find(value)).To(BeNil())
			})
		})
	})
})
