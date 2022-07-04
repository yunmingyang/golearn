package main

import "fmt"


func squares() func() int {
	var x int

	return func() int {
		x++

		return x * x
	}
}

func main() {
	f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	f1 := squares()

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
}