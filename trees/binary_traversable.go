/*
TODO: Add package comment here.

http://golang.org/doc/effective_go.html#commentary
*/
package trees

import (
	"github.com/modocache/cargo"
	"github.com/modocache/cargo/comparators"
	"github.com/modocache/cargo/queues"
	"math"
)

type BinaryTraversable interface {
	Parent() BinaryTraversable
	SetParent(traversable BinaryTraversable)
	Left() BinaryTraversable
	SetLeft(traversable BinaryTraversable)
	Right() BinaryTraversable
	SetRight(traversable BinaryTraversable)
	Value() interface{}
}

type TraversalCallback func(traversable BinaryTraversable) bool

func Root(traversable BinaryTraversable) BinaryTraversable {
	return root(traversable, nil)
}

func Depth(traversable BinaryTraversable) int {
	depth := -1
	callback := func(traversable BinaryTraversable) bool {
		depth++
		return false
	}
	root(traversable, callback)
	return depth
}

func root(traversable BinaryTraversable, callback TraversalCallback) BinaryTraversable {
	if traversable == nil {
		panic("attempt to pass trees.Root() a nil object")
	}

	if callback != nil {
		callback(traversable) // Ignore callback value; always find root
	}

	if isOrhpan(traversable) {
		return traversable
	} else {
		return root(traversable.Parent(), callback)
	}
}

func Height(traversable BinaryTraversable) int {
	if traversable == nil || isLeaf(traversable) {
		return 0
	} else {
		leftHeight, rightHeight := 0, 0
		if left := traversable.Left(); left != nil {
			leftHeight = Height(left)
		}
		if right := traversable.Right(); right != nil {
			rightHeight = Height(right)
		}

		return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
}

func IsBalanced(traversable BinaryTraversable) bool {
	if traversable == nil || isLeaf(traversable) {
		return true
	} else {
		leftHeight := Height(traversable.Left())
		rightHeight := Height(traversable.Right())
		return int(math.Abs(float64(leftHeight-rightHeight))) < 2
	}
}

func BreadthFirstSearch(traversable BinaryTraversable, callback TraversalCallback) {
	queue := queues.NewQueue()
	pushOnQueueIfNotNil(queue, traversable)
	breadthFirstSearch(queue, callback)
}

func breadthFirstSearch(queue *queues.Queue, callback TraversalCallback) {
	if queue.IsEmpty() {
		return
	}

	traversable := queue.Pop().(BinaryTraversable)
	if callback(traversable) {
		return
	}

	pushOnQueueIfNotNil(queue, traversable.Left())
	pushOnQueueIfNotNil(queue, traversable.Right())
	breadthFirstSearch(queue, callback)
}

func pushOnQueueIfNotNil(queue *queues.Queue, element interface{}) {
	if element != nil {
		queue.Push(element)
	}
}

func DepthFirstSearch(traversable BinaryTraversable,
	order cargo.TraversalOrder, callback TraversalCallback) {
	switch order {
	case cargo.PreOrder:
		depthFirstSearchPreOrder(traversable, callback)
	case cargo.InOrder:
		depthFirstSearchInOrder(traversable, callback)
	case cargo.PostOrder:
		depthFirstSearchPostOrder(traversable, callback)
	}
}

func depthFirstSearchPreOrder(traversable BinaryTraversable, callback TraversalCallback) {
	if traversable != nil {
		if callback(traversable) {
			return
		}
		depthFirstSearchPreOrder(traversable.Left(), callback)
		depthFirstSearchPreOrder(traversable.Right(), callback)
	}
}

func depthFirstSearchInOrder(traversable BinaryTraversable, callback TraversalCallback) {
	if traversable != nil {
		depthFirstSearchInOrder(traversable.Left(), callback)
		if callback(traversable) {
			return
		}
		depthFirstSearchInOrder(traversable.Right(), callback)
	}
}

func depthFirstSearchPostOrder(traversable BinaryTraversable, callback TraversalCallback) {
	if traversable != nil {
		depthFirstSearchPostOrder(traversable.Left(), callback)
		depthFirstSearchPostOrder(traversable.Right(), callback)
		if callback(traversable) {
			return
		}
	}
}

func IsBinarySearchTree(traversable BinaryTraversable, less comparators.Less) bool {
	leftIs, rightIs := true, true
	left, right := traversable.Left(), traversable.Right()

	if left != nil {
		if hasLeftGrandChildren(traversable) {
			leftIs = IsBinarySearchTree(left, less)
		} else {
			leftIs = less(left.Value(), traversable.Value())
		}
	}

	if right != nil {
		if hasRightGrandChildren(traversable) {
			rightIs = IsBinarySearchTree(right, less)
		} else {
			rightIs = !less(right.Value(), traversable.Value())
		}
	}

	return rightIs && leftIs
}

func rotateLeft(x BinaryTraversable) {
	y := x.Right()

	// y's left is now x's right. Update the parent as well.
	x.SetRight(y.Left())
	if x.Right() != nil {
		x.Right().SetParent(x)
	}

	// y's parent is what x's parent used to be.
	// Update the parent's left/right reference as well.
	y.SetParent(x.Parent())
	if x.Parent() != nil && x.Parent().Left() == x {
		x.Parent().SetLeft(y)
	} else if x.Parent() != nil && x.Parent().Right() == x {
		x.Parent().SetRight(y)
	}

	// Finalize the rotation.
	y.SetLeft(x)
	x.SetParent(y)
}

func rotateRight(y BinaryTraversable) {
	x := y.Left()

	// x's right is now y's left. Update the parent as well.
	y.SetLeft(x.Right())
	if y.Left() != nil {
		y.Left().SetParent(y)
	}

	// x's parent is what y's parent used to be.
	// Update the parent's left/right reference as well.
	x.SetParent(y.Parent())
	if y.Parent() != nil && y.Parent().Left() == y {
		x.Parent().SetLeft(x)
	} else if y.Parent() != nil && y.Parent().Right() == y {
		x.Parent().SetRight(x)
	}

	// Finalize the rotation.
	x.SetRight(y)
	y.SetParent(x)
}

func isOrhpan(searchable BinaryTraversable) bool {
	return searchable.Parent() == nil
}

func isLeaf(searchable BinaryTraversable) bool {
	return searchable.Left() == nil && searchable.Right() == nil
}

func hasGrandChildren(traversable BinaryTraversable) bool {
	return hasLeftGrandChildren(traversable) || hasRightGrandChildren(traversable)
}

func hasLeftGrandChildren(traversable BinaryTraversable) bool {
	if left := traversable.Left(); left == nil {
		return false
	} else {
		return !isLeaf(left)
	}
}

func hasRightGrandChildren(traversable BinaryTraversable) bool {
	if right := traversable.Right(); right == nil {
		return false
	} else {
		return !isLeaf(right)
	}
}

func isRightChild(searchable BinaryTraversable) bool {
	return !isOrhpan(searchable) && searchable == searchable.Parent().Right()
}

func isLeftChild(searchable BinaryTraversable) bool {
	return !isOrhpan(searchable) && searchable == searchable.Parent().Left()
}

func isGrandchild(searchable BinaryTraversable) bool {
	return !isOrhpan(searchable) && !isOrhpan(searchable.Parent())
}

func isParentLeftChild(searchable BinaryTraversable) bool {
	return isGrandchild(searchable) && isLeftChild(searchable.Parent())
}

func rightUncle(searchable BinaryTraversable) BinaryTraversable {
	if isGrandchild(searchable) {
		return searchable.Parent().Parent().Right()
	} else {
		return nil
	}
}

func leftUncle(searchable BinaryTraversable) BinaryTraversable {
	if isGrandchild(searchable) {
		return searchable.Parent().Parent().Left()
	} else {
		return nil
	}
}
