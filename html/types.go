package html

type TrustedString string

type Attributes map[string]string

type Children []any

type Node struct {
	Element    string
	Attributes Attributes
	Children   Children
}
