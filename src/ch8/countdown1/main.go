package main

import (
	"fmt"
	"time"
)



func lanuch() {
	fmt.Println("Lift off")
}

func main() {
	fmt.Println("Commencing countdown...")

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<- tick
	}

	lanuch()
}