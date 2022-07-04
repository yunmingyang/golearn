package main

import (
	"bytes"
	"fmt"
	"io"
)



const debug = false



func main() {
	var buf io.Writer // If *bytes.Buffer will cause a panic when debug is false, since its a pointer which point to a bytes.Buffer, and its value is nil, so out != nil will be true.
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	// if debug {

	// }
}

func f(out io.Writer) {
	fmt.Printf("%v\n", out != nil)
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}