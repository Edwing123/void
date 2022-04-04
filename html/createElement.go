package html

import (
	"strconv"

	"github.com/Edwing123/void/html/elements"
)

func CreateElement(name string, attributes Attributes, children ...any) Node {
	// Verify that there's an element with that name.
	if _, ok := elements.Elements[name]; !ok {
		panic("element " + strconv.Quote(name) + " is not a valid HTML element")
	}

	// Validate that children only contains valid child types.
	if !validateChildren(children) {
		panic("invalid child type: child type must be " + validChildTypes)
	}

	return Node{
		Element:    name,
		Attributes: attributes,
		Children:   children,
	}
}
