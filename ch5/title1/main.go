package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)




func main() {
	fmt.Println(title(os.Args[1]))
}

func title(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	checkContent := res.Header.Get("Content-Type")
	if checkContent != "text/html" && !strings.HasPrefix(checkContent, "text/html;") {
		res.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, checkContent)
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return fmt.Errorf("pasrsing %s as HTML: %v", url, err)
	}

	visitNode := func (n *html.Node)  {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}