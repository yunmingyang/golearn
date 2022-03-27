package main

import "fmt"


type tree struct {
	value int
	left, right *tree
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	printTree(root)
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int{
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree{
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func printTree(root *tree) {
	if root == nil {
		return
	}
	printTree(root.left)
	fmt.Println(root.value)
	printTree(root.right)
}


func main()  {
	values := []int{1, 30, 5, 66, 78, 69, 100}

	sort(values)
	fmt.Println(values)
}