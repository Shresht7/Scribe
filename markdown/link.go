package markdown

import "github.com/Shresht7/Scribe/scribe"

//* LINK *//

// Link is a link.
type NodeLink struct {
	URL   string
	Title string
	scribe.NodeText
}

// Instantiate a new link with the given contents
func Link(title, url string) *NodeLink {
	return &NodeLink{
		URL:   url,
		Title: title,
	}
}

// Implement the Node interface for NodeLink
func (link *NodeLink) String() string {
	return "[" + link.Title + "](" + link.URL + ")"
}
