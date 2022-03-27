package main

import (
	"fmt"
	"log"
	"net/http"
)



func main(){
	http.HandleFunc("/", handlerS1)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handlerS1(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}