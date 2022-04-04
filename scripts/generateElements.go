package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
	@author Edwin Evenor Garcia Valle.

	Generate file with elements map.
*/

// This is the head of the auto generated file.
const head = `
package elements

/*
	@author Edwin Evenor Garcia Valle.

	This is an auto generated file.
*/

var EmptyElements = map[string]struct{}{
	"area": {},
	"base": {},
	"br": {},
	"col": {},
	"embed": {},
	"hr": {},
	"img": {},
	"input": {},
	"link": {},
	"meta": {},
	"param": {},
	"source": {},
	"track": {},
	"wbr": {},
}
`

// Path to the directory of the HTML elements.
const elementsDir = "./mdn/element"
const outputFile = "./html/elements/elements.go"

// Represents information about an HTML element.
type Element struct {
	Name           string
	IsDeprecated   bool
	OmitClosingTag bool
}

func checkIfDeprecated(content string) bool {
	lines := strings.SplitN(content, "\n", 20)
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		if strings.Contains(line, "Deprecated") {
			return true
		}
	}

	return false
}

func getElements() ([]Element, error) {
	getErrorValues := func(err error) ([]Element, error) {
		return nil, err
	}

	entries, err := os.ReadDir(elementsDir)
	if err != nil {
		return getErrorValues(err)
	}

	var elements []Element

	for _, entry := range entries {
		path := elementsDir + "/" + entry.Name() + "/index.md"

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return getErrorValues(err)
		}

		// Heading elements, e.g, h1-h6 are all grouped inside heading_elements directory,
		// so I have to create those entries explicitly.
		if entry.Name() == "heading_elements" {
			for i := 1; i <= 6; i++ {
				elements = append(elements, Element{
					Name:           fmt.Sprintf("%s%d", "h", i),
					IsDeprecated:   false,
					OmitClosingTag: false,
				})
			}

			continue
		}

		deprecated := checkIfDeprecated(string(content))

		if !deprecated {
			elements = append(elements, Element{
				Name:         entry.Name(),
				IsDeprecated: deprecated,
			})
		}
	}

	return elements, nil
}

func createElementsMap(elements []Element) string {
	code := "var Elements = map[string]struct{}{\n"

	for _, element := range elements {
		code += fmt.Sprintf("\t%q: {},\n", element.Name)
	}

	code += "}"

	return code
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	elements, err := getElements()
	checkError(err)

	elementsMap := createElementsMap(elements)
	code := fmt.Sprintf("%s\n%s\n", head, elementsMap)

	err = ioutil.WriteFile(outputFile, []byte(code), 0o744)
	checkError(err)
}
