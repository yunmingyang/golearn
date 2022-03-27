package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)



var verbose = flag.Bool("v", false, "show verbose progress message")

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
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
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	go func () {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}

		close(fileSizes)
	}()

	var tick  <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Microsecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select{
		case size, ok := <- fileSizes:
			if !ok {
				break loop // fileSize was closed, if only break here, the select loop will be broke, but the for loop will not, so use label break here
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}
