package html

import (
	"fmt"

	"github.com/Edwing123/void/html/elements"
)

func Render(node Node) string {
	_, isEmptyElement := elements.EmptyElements[node.Element]
	openingTag := createOpeningTag(node.Element, isEmptyElement, node.Attributes)

	if isEmptyElement {
		return openingTag
	}

	var body string
	for _, child := range node.Children {
		switch child.(type) {
		case string:
			body += sanatizeText(child.(string))

		case TrustedString:
			body += string(child.(TrustedString))

		case Node:
			body += Render(child.(Node))
		}
	}

	return openingTag + body + createClosingTag(node.Element)
}

// Creates the opening tag for the element name and add its
// atributes.
func createOpeningTag(name string, isEmptyElement bool, attributes Attributes) string {
	code := "<" + name

	for name, value := range attributes {
		code += fmt.Sprintf(" %s=%q", name, value)
	}

	if isEmptyElement {
		code += " />"
	} else {
		code += ">"
	}

	return code
}

// Creates the closing tag for name.
func createClosingTag(name string) string {
	return "</" + name + ">"
}
