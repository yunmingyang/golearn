package main

import (
	"html/template"
	"log"
	"os"

	"golearn/ch4/github"
)

func main() {
	var issuesList = template.Must(template.New("issueList").Parse(`
	<h1>{{.TotalCount}}</h1>
	<table>
		<tr style='text-align: left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
	</table>
	`))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issuesList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
