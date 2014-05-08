/*
TODO: Add package comment here.

http://golang.org/doc/effective_go.html#commentary
*/
package trees

import (
	"github.com/modocache/cargo/comparators"
	"math"
	"reflect"
)

type BinarySearchable interface {
	Parent() BinarySearchable
	SetParent(searchable BinarySearchable)
	Left() BinarySearchable
	SetLeft(searchable BinarySearchable)
	Right() BinarySearchable
	SetRight(searchable BinarySearchable)
	Value() interface{}
	Insert(value interface{}) BinarySearchable
	Find(value interface{}) BinarySearchable
	Less() comparators.Less
}

type childConstructor func(parent BinarySearchable, value interface{}) BinarySearchable
type searchableCallback func(searchable BinarySearchable)

func Root(searchable BinarySearchable) BinarySearchable {
	return root(searchable, nil)
}

func Depth(searchable BinarySearchable) int {
	depth := -1
	callback := func(searchable BinarySearchable) { depth++ }
	root(searchable, callback)
	return depth
}

func Height(searchable BinarySearchable) int {
	if isLeaf(searchable) {
		return 0
	} else {
		leftHeight, rightHeight := 0, 0
		if left := searchable.Left(); left != nil {
			leftHeight = Height(left)
		}
		if right := searchable.Right(); right != nil {
			rightHeight = Height(right)
		}

		return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
}

func IsBalanced(searchable BinarySearchable) bool {
	return true // TODO
}

func root(searchable BinarySearchable, callback searchableCallback) BinarySearchable {
	if searchable == nil {
		panic("attempt to pass trees.Root() a nil object")
	}

	if callback != nil {
		callback(searchable)
	}

	if isOrhpan(searchable) {
		return searchable
	} else {
		return root(searchable.Parent(), callback)
	}
}

func insert(searchable BinarySearchable, value interface{}, constructor childConstructor) BinarySearchable {
	if searchable.Less()(value, searchable.Value()) {
		if searchable.Left() == nil {
			searchable.SetLeft(constructor(searchable, value))
			return searchable.Left()
		} else {
			return searchable.Left().Insert(value)
		}
	} else {
		if searchable.Right() == nil {
			searchable.SetRight(constructor(searchable, value))
			return searchable.Right()
		} else {
			return searchable.Right().Insert(value)
		}
	}
}

func find(tree *BinarySearchTree, value interface{}) BinarySearchable {
	if reflect.DeepEqual(value, tree.Value()) {
		return tree
	} else if tree.Less()(value, tree.Value()) {
		if tree.Left() == nil {
			return nil
		} else {
			return tree.Left().Find(value)
		}
	} else {
		if tree.Right() == nil {
			return nil
		} else {
			return tree.Right().Find(value)
		}
	}
}

func rotateLeft(x BinarySearchable) {
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

func rotateRight(y BinarySearchable) {
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

func isOrhpan(searchable BinarySearchable) bool {
	return searchable.Parent() == nil
}

func isLeaf(searchable BinarySearchable) bool {
	return searchable.Left() == nil && searchable.Right() == nil
}

func isRightChild(searchable BinarySearchable) bool {
	return !isOrhpan(searchable) && searchable == searchable.Parent().Right()
}

func isLeftChild(searchable BinarySearchable) bool {
	return !isOrhpan(searchable) && searchable == searchable.Parent().Left()
}

func isGrandchild(searchable BinarySearchable) bool {
	return !isOrhpan(searchable) && !isOrhpan(searchable.Parent())
}

func isParentLeftChild(searchable BinarySearchable) bool {
	return isGrandchild(searchable) && isLeftChild(searchable.Parent())
}

func rightUncle(searchable BinarySearchable) BinarySearchable {
	if isGrandchild(searchable) {
		return searchable.Parent().Parent().Right()
	} else {
		return nil
	}
}

func leftUncle(searchable BinarySearchable) BinarySearchable {
	if isGrandchild(searchable) {
		return searchable.Parent().Parent().Left()
	} else {
		return nil
	}
}
