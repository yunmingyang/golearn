package main

import "fmt"

func reverse(t *[]int) {
	for i, j := 0, len(*t) - 1; i < j; i, j = i + 1, j - 1 {
		(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
	}
}

func main() {
	numArr := new([]int)

	*numArr = append(*numArr, 1, 2, 3, 4, 5)
	reverse(numArr)

	fmt.Println(numArr)
}