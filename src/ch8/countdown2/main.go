package main

import (
	"fmt"
	"os"
	"time"
)



func lanuch() {
	fmt.Println("Lift off")
}

func main() {
	fmt.Println("Commencing countdown. Please return to abort")

	timer := time.After(10 * time.Second)
	abort := make(chan struct{})

	go func () {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for countdown := 10; countdown >= 0; countdown-- {
		select{
		case <- timer:
			lanuch()
		case <- abort:
			fmt.Println("Lanuch abort")
			return
		default:
			fmt.Println(countdown)
			time.Sleep(time.Second)
		}
	}

	ch := make(chan int, 1)
	for i := 0; i < 10; i++{
		// In
		select {
		case x := <- ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}