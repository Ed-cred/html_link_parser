package main

import (
	"fmt"
	"strings"

	link "github.com/Ed-cred/html_link_parser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/some-page">
	A different link
	<span> Some span </span
  </a>	
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
