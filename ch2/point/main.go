package main

import "fmt"



func main(){
	x := 1
	p := &x

	fmt.Println(p)
	fmt.Println(*p)

	*p = 2

	fmt.Println(x)

	pp := &p

	fmt.Println(pp)
	fmt.Println(*pp)
	fmt.Println(**pp)
	fmt.Println(p == *pp)

	pptr := new(*int)
	fmt.Println(pptr)

	ptr := new(int)
	fmt.Println(*ptr)

	pptr = &ptr
	fmt.Println(**pptr)

	**pptr++
	fmt.Println(*ptr)
	fmt.Println(**pptr)
}