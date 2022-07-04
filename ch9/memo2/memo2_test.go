package memo2_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"golearn/ch9/memo2"
	"golearn/ch9/memoutils"
)

func TestConcurrent(t *testing.T) {
	var n sync.WaitGroup
	m := memo2.New(memoutils.HttpGetBody)

	for url := range memoutils.IncomingURLS() {
		n.Add(1)

		go func(url string) {
			start := time.Now()
			fmt.Println(start)
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
