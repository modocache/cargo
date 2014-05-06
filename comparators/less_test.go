package comparators_test

import (
	. "github.com/modocache/cargo/comparators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Less", func() {
	Describe("IntLess", func() {
		var value, treeValue int
		Context("when value is less than treeValue", func() {
			BeforeEach(func() {
				value = 1
				treeValue = 2
			})

			It("returns true", func() {
				Expect(IntLess(value, treeValue)).To(BeTrue())
			})
		})

		Context("when value is not less than treeValue", func() {
			BeforeEach(func() {
				value = 1
				treeValue = 1
			})

			It("returns false", func() {
				Expect(IntLess(value, treeValue)).To(BeFalse())
			})
		})
	})

	Describe("StringLess", func() {
		var value, treeValue string
		Context("when value is less than treeValue", func() {
			BeforeEach(func() {
				value = "aardvark"
				treeValue = "bear"
			})

			It("returns true", func() {
				Expect(StringLess(value, treeValue)).To(BeTrue())
			})
		})

		Context("when value is not less than treeValue", func() {
			BeforeEach(func() {
				value = "cougar"
				treeValue = "bear"
			})

			It("returns false", func() {
				Expect(StringLess(value, treeValue)).To(BeFalse())
			})
		})
	})
})
