package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)




func main() {
	fmt.Println(fetch(os.Args[1]))
}

func fetch(url string) (filename string, n int64, err error) {
	res, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	local := path.Base(res.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, res.Body)
	if closeErr := f.Close(); err != nil {
		err = fmt.Errorf("copy error: %v\nclose error: %v", err, closeErr)
	}

	return local, n, err
}