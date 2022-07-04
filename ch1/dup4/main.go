package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLinesDup4(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLinesDup4(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLinesDup4(f *os.File, counts map[string]int) {
	flag := false
	input := bufio.NewScanner(f)

	for input.Scan() {
		if counts[input.Text()] > 1 && f != os.Stdin && !flag {
			fmt.Println(f.Name())
			flag = true
		}
		counts[input.Text()]++
	}
}
