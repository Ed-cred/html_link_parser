package link

import "io"

// Link represent a link (<a href="...">) in an html document
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
