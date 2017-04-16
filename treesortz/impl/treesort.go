package treesortz

type tree struct {
	value       int
	left, right *tree
	count       int
}

<<<<<<< HEAD
//Sort ...
=======
// Sort ...
>>>>>>> 4c4086c1fa00138e6c8d30d047cc22095c776248
func Sort(values []int) {
	var t *tree
	for _, v := range values {
		t = add(t, v)
	}
	values = values[:0]
	appendValues(values, t)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
<<<<<<< HEAD
func appendValues(values []int, t *tree) {
=======
func appendValues(values []int, t *tree) []int {
>>>>>>> 4c4086c1fa00138e6c8d30d047cc22095c776248

	if t.left != nil {
		appendValues(values, t.left)
	}

<<<<<<< HEAD
	values[len(values)] = t.value

=======
	for i := 0; i < t.count; i++ {
		values = append(values, t.value)
	}
>>>>>>> 4c4086c1fa00138e6c8d30d047cc22095c776248
	if t.right != nil {
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = &tree{count: 1}
		t.value = value
		return t
	}

	if value == t.value {
		t.count++
	} else if value < t.value {
		t.left = add(t.left, value)
	} else if value > t.value {
		t.right = add(t.right, value)
	}

	return t
}
