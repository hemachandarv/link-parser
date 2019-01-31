package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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
	return fmt.Sprintf("Href: %s\nText: %s\n\n", link.href, link.text)
}

/*
Parse will read the contents from io.Reader "r"
and return a slice of Link found in it.
*/
func Parse(r io.Reader) (links []Link, err error) {
	node, err := html.Parse(r)
	if err != nil {
		return
	}
	links = dfs(node)
	return
}

func dfs(node *html.Node) (links []Link) {
	if node.Type == html.ElementNode && node.DataAtom == atom.A {
		return buildLink(node)
	}
	for next := node.FirstChild; next != nil; next = next.NextSibling {
		childLinks := dfs(next)
		links = append(links, childLinks...)
	}
	return
}

func buildLink(node *html.Node) (links []Link) {
	var href, text string
	for _, attr := range node.Attr {
		if attr.Key == atom.Href.String() {
			href = attr.Val
			break
		}
	}
	text = buildLinkText(node)
	return []Link{
		Link{href: href, text: strings.TrimSpace(text)},
	}
}

func buildLinkText(node *html.Node) (text string) {
	if node.Type == html.TextNode {
		return node.Data
	}
	for next := node.FirstChild; next != nil; next = next.NextSibling {
		text += buildLinkText(next)
	}
	return
}
