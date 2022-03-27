package memoutils

import (
	"io/ioutil"
	"log"
	"net/http"
)



func HttpGetBody(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func IncomingURLS() <-chan string {
	ch := make(chan string)

	go func ()  {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}