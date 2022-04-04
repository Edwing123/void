package html

type TrustedString string

type Attributes map[string]string

type Node struct {
	Element    string
	Attributes Attributes
	Children   []any
}
