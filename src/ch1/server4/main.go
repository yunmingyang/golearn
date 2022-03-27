package main

import (
	"golearn/ch1/lissajous"
	"log"
	"net/http"
)

func main() {
	// handler := func(w http.ResponseWriter, r *http.Request){
	// 	lissajous.Lissajous(w)
	// }
	// http.HandleFunc("/", handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:9099", nil))
}
