package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)


func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invaild := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charout: %s\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invaild ++
			continue
		}

		counts[r] ++
		utflen[n] ++
	}

	 fmt.Printf("rune\tcount\n")
	 for c, n := range counts {
		 fmt.Printf("%q\t%d\n", c, n)
	 }

	 fmt.Printf("\nlen\tcount\n")
	 for i, n := range utflen {
		 if i > 0 {
			 fmt.Printf("%d\t%d\n", i, n)
		 }
	 }

	 if invaild > 0 {
		 fmt.Printf("\n %d invaild UTF-8 characters\n", invaild)
	 }
}