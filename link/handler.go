package link

import (
	"fmt"
	"io"
)

/*
Link contains the href and text attributes of
a parsed Anchor tag.
*/
type Link struct {
	href string
	text string
}

func (link Link) String() string {
	return fmt.Sprintf("Href: %s\nText:%s\n\n", link.href, link.text)
}

/*
Parse will read the contents from io.Reader "r"
and return a slice of Link found in it.
*/
func Parse(r io.Reader) (links []Link) {
	return
}
