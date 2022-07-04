package memo1_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"golearn/ch9/memo1"
	"golearn/ch9/memoutils"
)

func TestSequential(t *testing.T) {
	m := memo1.New(memoutils.HttpGetBody)

	for url := range memoutils.IncomingURLS() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func TestConcurrent(t *testing.T) {
	m := memo1.New(memoutils.HttpGetBody)
	var n sync.WaitGroup

	for url := range memoutils.IncomingURLS() {
		n.Add(1)

		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s %s %d bytes\n", url, time.Since(start), len(value.([]byte)))

			n.Done()
		}(url)
	}

	n.Wait()
}
