package trees

import (
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

type traversableCallback func(traversable BinaryTraversable)

func Root(traversable BinaryTraversable) BinaryTraversable {
	return root(traversable, nil)
}

func Depth(traversable BinaryTraversable) int {
	depth := -1
	callback := func(traversable BinaryTraversable) { depth++ }
	root(traversable, callback)
	return depth
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

func root(traversable BinaryTraversable, callback traversableCallback) BinaryTraversable {
	if traversable == nil {
		panic("attempt to pass trees.Root() a nil object")
	}

	if callback != nil {
		callback(traversable)
	}

	if isOrhpan(traversable) {
		return traversable
	} else {
		return root(traversable.Parent(), callback)
	}
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
