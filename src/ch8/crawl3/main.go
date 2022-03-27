package main

import (
	"fmt"
	"golearn/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)

	link, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}

	return link
}

func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		workList <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
