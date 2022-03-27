package main

import (
	"html/template"
	"log"
	"os"
)



func main() {
	const templ = `<p>A: {{.A}}</p><p>b: {{.B}}</p>`
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello world!</b>"
	data.B = "<b>Hello world!</b>"

	t := template.Must(template.New("escape").Parse(templ))
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}