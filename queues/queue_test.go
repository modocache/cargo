package queues_test

import (
	. "github.com/modocache/cargo/queues"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	queue := NewQueue()
	Describe(".Push() and .Pop()", func() {
		Context("when the queue is empty", func() {
			It("panics when popped", func() {
				Expect(func() { queue.Pop() }).To(Panic())
			})
		})

		Context("when the queue is not empty", func() {
			BeforeEach(func() {
				queue.Push("one")
				queue.Push(2)
				queue.Push([]int{3})
			})

			It(".Pop() removes elements from the queue (FIFO)", func() {
				Expect(queue.Pop()).To(Equal("one"))
				Expect(queue.Pop()).To(Equal(2))
				Expect(queue.Pop()).To(Equal([]int{3}))
				Expect(func() { queue.Pop() }).To(Panic())
			})
		})
	})
})
