package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"realeased"`
	// some details: https://pkg.go.dev/encoding/json
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movie = []Movie{
	{Title: "Caseblance", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}


func main() {
	// Unindent
	data, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("JSON marshal failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// Indent
	data, err = json.MarshalIndent(movie, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshalIndent failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//Unmarshaling
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshling failed: %s\n", err)
	}
	fmt.Printf("%s\n", titles)
}