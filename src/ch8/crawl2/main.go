package main

import (
	"fmt"
	"log"
	"os"

	"golearn/ch5/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}

	return list
}

func main() {
	workList := make(chan []string)
	var n int

	n++
	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
