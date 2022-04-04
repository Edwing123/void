# Void API reference

In this document I write about the important functions and types that together
create the public API for Void.

## Types

### Node

A `Node` is a struct type for represeting an HTML Node, it has three fields:

- Element: `string`
- Attributes: `Attributes`
- Children: `[]any`

**Element** is the name of the HTML node, for example: `html`, `p`, `img`.

**Attributes** is a map representing attributes.

**Children** is a list of children of the `Node`.

> Note: Children is of type `[]any`, but when rendering a `Node`, the only
> supported types for a child are: `string|Node|TrustedString`.

#### Examples

```go
// An h1 element with no attributes and a single child.
title := Node{
		Element:  "h1",
		Children: []any{"Hello Void"},
}
```

### Attributes

The `Attributes` type represents a map of attributes for a `Node`.

#### Examples

```go
userName := "Edwing123"

cardTitle := Node{
    Element: "h3",
    Attributes: Attributes{
        "class": "card__title",
        "data-is-admin": "true",
    },
    Children: []any{userName},
}
```

### TrustedString

The type `TrustedString` represents a string that should be not sanatized
when rendering.

This can be used to include raw HTML in your output HTML string, but keep in mind
that doing so can be dangerous, you should only used it with trusted text.

#### Examples

```go
title := Node{
		Element:  "h1",
		Children: []any{TrustedString(`<script>alert('Hello Void')</script>`)},
}

renderedTitle := Render(title) // <h1><script>alert('Hello Void')</script></h1>
```

## Functions

### CreateElement

`CreateElement` creates a `Node` struct with the provided arguments.

**Arguments**:

- name
- attributes
- ...children

> Note: Everything after the second argument becomes child of the created `Node`.

`CreateElement` does some validations to the arguments you pass in:

- It validates that the provided name is a valid HTML element (does it exist?).
- It validates that the children are of the expected types.

If any of this validations fail, `CreateElement` panics with an error.

> Note: You should preffer using `CreateElement` when creating elements, instead of creating
> instances of `Node` directly.

#### Examples

```go
links := []any{
    CreateElement("a", Attributes{"class": "nav__link", "href": "/"}, "Home"),
    CreateElement("a", Attributes{"class": "nav__link", "href": "/about"}, "About"),
    CreateElement("a", Attributes{"class": "nav__link", "href": "/contact"}, "Contact"),
}

mainNavigation := CreateElement(
    "nav",
    Attributes{"class": "nav"},
    links...,
)
```

### Render

`Render` takes in a `Node` and builds an HTML string representing the `Node`,
its attributes and its children.

#### Examples

```go
links := []any{
    CreateElement("a", Attributes{"class": "nav__link", "href": "/"}, "Home"),
    CreateElement("a", Attributes{"class": "nav__link", "href": "/about"}, "About"),
    CreateElement("a", Attributes{"class": "nav__link", "href": "/contact"}, "Contact"),
}

mainNavigation := CreateElement(
    "nav",
    Attributes{"class": "nav"},
    links...,
)

renderedNavigation := Render(mainNavigation) // <nav class="nav"><a class="nav__link" href="/">Home</a><a class="nav__link" href="/about">About</a><a class="nav__link" href="/contact">Contact</a></nav>
```
