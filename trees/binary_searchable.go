/*
TODO: Add package comment here.

http://golang.org/doc/effective_go.html#commentary
*/
package trees

import (
	"github.com/modocache/cargo/comparators"
	"reflect"
)

type BinarySearchable interface {
	BinaryTraversable
	Insert(value interface{}) BinarySearchable
	Find(value interface{}) BinarySearchable
	Less() comparators.Less
}

type childConstructor func(parent BinarySearchable, value interface{}) BinarySearchable

func insert(searchable BinarySearchable, value interface{}, constructor childConstructor) BinarySearchable {
	if searchable.Less()(value, searchable.Value()) {
		if searchable.Left() == nil {
			searchable.SetLeft(constructor(searchable, value))
			return searchable.Left().(BinarySearchable)
		} else {
			return searchable.Left().(BinarySearchable).Insert(value)
		}
	} else {
		if searchable.Right() == nil {
			searchable.SetRight(constructor(searchable, value))
			return searchable.Right().(BinarySearchable)
		} else {
			return searchable.Right().(BinarySearchable).Insert(value)
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
			return tree.Left().(BinarySearchable).Find(value)
		}
	} else {
		if tree.Right() == nil {
			return nil
		} else {
			return tree.Right().(BinarySearchable).Find(value)
		}
	}
}
