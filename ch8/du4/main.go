package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)



var verbose = flag.Bool("v", false, "show verbose progress message")
var sema = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select{
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)

			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select{
	case sema <- struct{}{}:
	case <-done:
		return nil
	}

	defer func ()  {
		<- sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}

	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	start := time.Now()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	fileSizes := make(chan int64)

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	go func () {
		temp := make([]byte, 1)
		os.Stdin.Read(temp)
		fmt.Println("temp is ", temp)
		close(done)
	}()

	go func () {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	var tick  <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Microsecond)
	}

loop:
	for {
		select{
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			// It is a channel without cache, so if not draining the channel,
			// all  left goroutine will be blocked
			for range fileSizes {
				// do nothing
			}
		case size, ok := <- fileSizes:
			if !ok {
				// fileSize was closed
				// If only break here, the select loop will be broke, but the for loop will not, so use label break here
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
	fmt.Println("compute period: ", time.Since(start))
	// panic("test")
}
