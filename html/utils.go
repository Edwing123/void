package html

import (
	"html"
)

func validateChildren(children []any) bool {
	for _, child := range children {
		switch child.(type) {
		case Node:
			break

		case string:
			break

		case TrustedString:
			break

		default:
			return false
		}
	}

	return true
}

func sanatizeText(text string) string {
	return html.EscapeString(text)
}
