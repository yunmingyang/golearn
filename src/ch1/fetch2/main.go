package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)



func main(){
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, res.Body)

		res.Body.Close()
		fmt.Printf("\nWrite %d bytes", b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error\nurl: %s\n%v", url, err)
			os.Exit(1)
		}
	}
}