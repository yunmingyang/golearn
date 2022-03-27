package memo4_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"golearn/ch9/memo4"
	"golearn/ch9/memoutils"
)

func TestConcurrent(t *testing.T) {
	var n sync.WaitGroup
	m := memo4.New(memoutils.HttpGetBody)

	for url := range memoutils.IncomingURLS() {
		n.Add(1)

		go func(url string) {
			start := time.Now()
			res, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s %s %d bytes\n", url, time.Since(start), len(res.([]byte)))

			n.Done()
		}(url)
	}

	n.Wait()
}
