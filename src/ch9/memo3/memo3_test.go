package memo3_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"golearn/ch9/memo3"
	"golearn/ch9/memoutils"
)

func TestSequential(t *testing.T) {
	var n sync.WaitGroup
	m := memo3.New(memoutils.HttpGetBody)

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
