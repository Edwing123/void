package html

import "testing"

func TestCreateElement(t *testing.T) {
	title := CreateElement("title", nil, "Hello Void")

	if len(title.Children) != 1 {
		t.Errorf("expected title to have one child")
	}

	if title.Element != "title" {
		t.Errorf("title.Element = got %q, want %q\n", title.Element, "title")
	}

	if title.Attributes != nil {
		t.Errorf("title.Attributes = got %v, want %v\n", title.Attributes, nil)
	}
}
