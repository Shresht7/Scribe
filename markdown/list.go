package markdown

import (
	"strconv"
	"strings"

	"github.com/Shresht7/Scribe/scribe"
)

// * LIST ITEM *//

// NodeListItem represents a list item in Markdown
type NodeListItem struct {
	scribe.NodeText
}

// Instantiate a new list item with the given contents
func NewListItem(contents ...any) *NodeListItem {
	// Create a new list item
	item := &NodeListItem{}
	item.WithSeparator(" ")

	// Append the contents to the list item
	item.AppendChild(contents...)

	// Return the list item
	return item
}

// Implement the Node interface for NodeListItem
func (item *NodeListItem) String() string {
	return item.NodeText.String()
}

// Create a new list item with the given contents
func ListItem(contents ...any) string {
	return NewListItem(contents...).String()
}

//* LIST *//

// List is a list of items.
type NodeList struct {
	Ordered   bool
	Decorator string
	Items     []*NodeListItem
}

// Instantiate a new list with the given contents
func NewOrderedList(items []any) *NodeList {
	list := &NodeList{
		Ordered: true,
	}
	for _, item := range items {
		list.Items = append(list.Items, NewListItem(item))
	}
	return list
}

// Instantiate a new list with the given contents
func NewUnorderedList(decorator string, items []any) *NodeList {
	// Validate the decorator
	if decorator != "*" && decorator != "-" && decorator != "+" {
		decorator = "*"
	}

	list := &NodeList{
		Ordered:   false,
		Decorator: decorator,
	}

	for _, item := range items {
		list.Items = append(list.Items, NewListItem(item))
	}

	return list
}

// Implement the Node interface for NodeList
func (list *NodeList) String() string {
	// Create a new string builder
	builder := strings.Builder{}

	// If the list is ordered, add the numbers
	if list.Ordered {
		// Iterate over the items
		for index, item := range list.Items {
			// Add the item to the builder
			builder.WriteString(strconv.Itoa(index+1) + ". " + item.String() + "\n")
		}
	} else {
		// Iterate over the items
		for _, item := range list.Items {
			// Add the item to the builder
			builder.WriteString(list.Decorator + " " + item.String() + "\n")
		}
	}

	// Return the string
	return builder.String()
}

// Write an unordered list to the document
func (doc *MarkdownDocument) WriteUnorderedList(items []any) *MarkdownDocument {
	// Create a new list
	list := NewUnorderedList("*", []any{})
	for _, item := range items {
		list.Items = append(list.Items, NewListItem(item))
	}

	// Add the list to the document
	doc.AppendChild(list)

	// Return the document
	return doc
}

func UnorderedList(items []any) string {
	return NewUnorderedList("*", items).String()
}

// Write an ordered list to the document
func (doc *MarkdownDocument) WriteOrderedList(items []any) *MarkdownDocument {
	// Create a new list
	list := NewOrderedList([]any{})
	for _, item := range items {
		list.Items = append(list.Items, NewListItem(item))
	}

	// Add the list to the document
	doc.AppendChild(list)

	// Return the document
	return doc
}
