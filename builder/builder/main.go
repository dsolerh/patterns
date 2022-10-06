package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

func NewHTMLBuilder(name string) *HTMLBuilder {
	return &HTMLBuilder{
		rootName: name,
		root: HTMLElement{
			name:     name,
			text:     "",
			elements: []HTMLElement{},
		},
	}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(childName, childText string) {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *HTMLBuilder) AddChildFluent(childName, childText string) *HTMLBuilder {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Printf("sb.String(): %v\n", sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	// <ul><li></li>...</ul>
	sb.WriteString("<ul>")
	for _, word := range words {
		sb.WriteString(fmt.Sprintf("<li>%s</li>", word))
	}
	sb.WriteString("</ul>")
	fmt.Printf("sb.String(): %v\n", sb.String())

	b := NewHTMLBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world").
		AddChildFluent("li", "!!")
	fmt.Printf("b: %s\n", b)
}
