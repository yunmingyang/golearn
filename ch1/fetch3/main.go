package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func main(){
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://"){
			url = "http://" + url
		}
		fmt.Printf("fetch from url: %s\n", url)

		res, err := http.Get(url)
		// exercise 1.9
		fmt.Printf("Staue code: %d\n", res.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %s", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, res.Body)

		res.Body.Close()
		fmt.Printf("\nWrite %d bytes", b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error\nurl: %s\n%s", url, err)
			os.Exit(1)
		}
	}
}