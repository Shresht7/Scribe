package markdown

import "github.com/Shresht7/Scribe/scribe"

//* LINK *//

// NodeLink represents a link in Markdown
type NodeLink struct {
	URL   string
	Title string
	scribe.NodeText
}

// Instantiate a new link with the given contents
func NewLink(title, url string) *NodeLink {
	return &NodeLink{
		URL:   url,
		Title: title,
	}
}

// Implement the Node interface for NodeLink
func (link *NodeLink) String() string {
	return "[" + link.Title + "](" + link.URL + ")"
}

func Link(title, url string) string {
	return NewLink(title, url).String()
}
