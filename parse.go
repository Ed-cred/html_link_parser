package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represent a link (<a href="...">) in an html document
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
