package main

import "fmt"


func main() {
	data := []string{"one", "", "three"}


	fmt.Printf("%q\n", nonEmpty2(data))
	fmt.Printf("data is: %v\n", data)
	fmt.Printf("%q\n", nonEmpty(data))
	fmt.Printf("data is: %v\n", data)
}

func nonEmpty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonEmpty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}