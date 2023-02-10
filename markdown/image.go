package markdown

//* IMAGE *//

// Image is an image.
type NodeImage struct {
	URL string
	Alt string
}

// Instantiate a new image with the given contents
func Image(alt, url string) *NodeImage {
	return &NodeImage{
		URL: url,
		Alt: alt,
	}
}

// Implement the Node interface for NodeImage
func (image *NodeImage) String() string {
	return "![" + image.Alt + "](" + image.URL + ")"
}
