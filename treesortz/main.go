package main

import (
	"fmt"

	treesortz "github.com/teo-mateo/gostuff/treesortz/impl"
)

func main() {
	values := []int{2, 44, 3, 6, 55, 17, 19, 6, 3}
	treesortz.Sort(values)
	fmt.Println(values)
}
