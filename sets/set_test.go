package sets_test

import (
	. "github.com/modocache/cargo/sets"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	var set *Set
	BeforeEach(func() { set = NewSet() })
	Describe("adding and removing elements", func() {
		It("adds elements", func() {
			set.Add("Rails")
			set.Add("Django")

			Expect(set.Cardinality()).To(Equal(2))
			Expect(set.Contains("Rails")).To(BeTrue())
			Expect(set.Contains("Django")).To(BeTrue())
		})

		It("removes elements", func() {
			set.Add("iOS")
			set.Add("Android")
			set.Remove("Android")

			Expect(set.Cardinality()).To(Equal(1))
			Expect(set.Contains("iOS")).To(BeTrue())
			Expect(set.Contains("Android")).To(BeFalse())
		})
	})

	Describe(".Elements()", func() {
		Context("when the set is empty", func() {
			It("returns an empty slice", func() {
				Expect(set.Elements()).To(Equal([]interface{}{}))
			})
		})

		Context("when the set contains elements", func() {
			BeforeEach(func() {
				set.Add("Haskell")
				set.Add("Go")
			})
			It("returns an array of those elements", func() {
				// FIXME: This test is brittle; it will fail if the order
				// in which the elements are returned changes, although
				// sets themselves are not unordered. I don't feel like
				// writing a slice equality function at the moment, though.
				Expect(set.Elements()).To(Equal([]interface{}{"Haskell", "Go"}))
			})
		})
	})
})
