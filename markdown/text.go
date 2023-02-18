package markdown

import "github.com/Shresht7/Scribe/scribe"

// Instantiates a new NodeText with the given contents
func Text(contents ...string) *scribe.NodeText {
	return scribe.Text(contents...)
}
