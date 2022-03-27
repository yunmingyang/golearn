package main

import (
	"fmt"
	"math"
)


func main() {
	// margin of error in float32
	var f float32 = 1 << 24 // 16777216
	fmt.Println(f == f + 1)

	// FP print
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	// Division by zero, infinity and NaN
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	// No value is equal to Nan
	// nan := math.NaN()
	// fmt.Println(nan == nan, nan < nan, nan > nan)
}