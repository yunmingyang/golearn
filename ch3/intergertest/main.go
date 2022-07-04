package main

import (
	"fmt"
)



func main() {
	// Bit operations
	var x int8 = 1 << 1 | 1 << 5
	var y int8 = 1 << 1 | 1 << 2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x & y)
	fmt.Printf("%08b\n", x | y)
	fmt.Printf("%08b\n", x ^ y)
	fmt.Printf("%08b\n", x &^ y)

	for i := uint(0); i < 8; i++ {
		if x & (1 << i) != 0 {
			fmt.Printf("%08b\n", i)
		}
	}

	fmt.Printf("%08b\n", x << 1)
	fmt.Printf("%08b\n", y >> 1)

	// len should not return a unsigned integrity, since it will cause ifinite loop
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	// Print about positional notation
	// Octal
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	// Hexadecimal
	x1 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x, %#[1]x, %#[1]X\n", x1)

	// Characters
	ascii := 'a'
	// rune
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}