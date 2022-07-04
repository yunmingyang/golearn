package main

import (
	"fmt"
	"unsafe"
)



func main() {
	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Println(x.b)

	// the same with pb := &x
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
}