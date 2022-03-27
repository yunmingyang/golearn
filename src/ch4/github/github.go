package github

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	"net/url"
	// "os"
	"strings"
	"time"
)


const IssueURL = "https://api.github.com/search/issues"


type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time
	Body string // markdown format
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}


func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)

	}

	var result IssueSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()
	return &result, nil
}

// func main() {
// 	result, err := SearchIssues(os.Args[1:])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%d issues\n", result.TotalCount)
// 	for _, item := range result.Items {
// 		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
// 	}

// }