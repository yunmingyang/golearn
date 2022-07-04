package main

import "fmt"



func main() {
	double(1)

	a := triple(3)
	fmt.Println(a)
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d \n", x, result)
	}()

	return x + x
}

func triple(x int) (result int) {
	defer func ()  {
		result += x
	}()

	return double(x)
}