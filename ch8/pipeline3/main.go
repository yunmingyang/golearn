package main

import (
	"fmt"
)



func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares,naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for v := 0; v < 100; v++ {
		out <- v
	}

	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}

	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}