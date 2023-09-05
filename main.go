package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type HtmlData struct {
	Link string
	Text string
}

func main() {
	// htmlFile := flag.String("html", "ex1.html", "HTML file to parse")
	flag.Parse()
	s := `<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/other-page">A link to another page</a>
	</body>
	</html>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Println(n)
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
