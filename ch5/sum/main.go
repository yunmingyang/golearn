package main

import (
	"fmt"
	"os"
)


func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))

	values := []int{2, 5, 6}
	fmt.Println(sum(values...))

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	lineNum, name := 12, "count"
	errorf(lineNum, "undefined: %s", name)
}

func sum(vals ...int) int{
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

func f(...int){}

func g([]int){}

func errorf(lineNum int, format string, args ...interface{}){
	fmt.Fprintf(os.Stderr, "Line %d: ", lineNum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}