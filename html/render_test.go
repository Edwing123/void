package html

import "testing"

func TestCreateOpeningTag(t *testing.T) {
	tests := []struct {
		name           string
		tag            string
		attributes     Attributes
		isEmptyElement bool
		want           string
	}{
		{
			name:           "div with class container",
			tag:            "div",
			attributes:     Attributes{"class": "container"},
			isEmptyElement: false,
			want:           `<div class="container">`,
		},
		{
			name:           "input with type and name",
			tag:            "input",
			attributes:     Attributes{"type": "text", "name": "..."},
			isEmptyElement: true,
			want:           `<input type="text" name="..." />`,
		},
		{
			name:           "img with src",
			tag:            "img",
			attributes:     Attributes{"src": "..."},
			isEmptyElement: true,
			want:           `<img src="..." />`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := createOpeningTag(tc.tag, tc.isEmptyElement, tc.attributes)

			if got != tc.want {
				t.Errorf("got %q, want %q\n", got, tc.want)
			}
		})
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name string
		node Node
		want string
	}{
		{
			name: "Render single element div",
			node: CreateElement("div", nil),
			want: "<div></div>",
		},
		{
			name: "Render empty element div with id",
			node: CreateElement("div", Attributes{"id": "..."}),
			want: `<div id="..."></div>`,
		},
		{
			name: "Render empty element div with id and text child",
			node: CreateElement("div", Attributes{"id": "..."}, "Hello Void"),
			want: `<div id="...">Hello Void</div>`,
		},
		{
			name: "Render img with attributes",
			node: CreateElement("img", Attributes{"src": "...", "id": "..."}),
			want: `<img src="..." id="..." />`,
		},
		{
			name: "Render h1 with sanatized text",
			node: CreateElement("h1", nil, "<script>alert('Hello World')</script>"),
			want: `<h1>&lt;script&gt;alert(&#39;Hello World&#39;)&lt;/script&gt;</h1>`,
		},
		{
			name: "Render h1 with TrustedString text",
			node: CreateElement("h1", nil, TrustedString("<script>alert('Hello World')</script>")),
			want: `<h1><script>alert('Hello World')</script></h1>`,
		},
		{
			name: "Render nested tree of elements",
			node: CreateElement(
				"html",
				nil,
				CreateElement(
					"head",
					nil,
					CreateElement("title", nil, "Hello Void"),
				),
				CreateElement(
					"body",
					nil,
					CreateElement(
						"div",
						nil,
						CreateElement("h1", Attributes{"class": "title"}, "Hello Void"),
					),
				),
			),
			want: `<html><head><title>Hello Void</title></head><body><div><h1 class="title">Hello Void</h1></div></body></html>`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Render(tc.node)

			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
