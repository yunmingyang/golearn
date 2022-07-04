package main

import "fmt"



func main() {
	fmt.Println(gcd(2, 4))
}


func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}