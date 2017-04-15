package treesortz

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var t *tree
	fmt.Println(t)

	for _, v := range values {
		t = add(t, v)
	}

	values = values[:0]
	appendValues(values, t)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {

	if t.left != nil {
		values = appendValues(values, t.left)
	}

	values = append(values, t.value)

	if t.right != nil {
		appendValues(values, t.right)
	}

	return values

}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	}
	if value > t.value {
		t.right = add(t.right, value)
	}

	return t
}