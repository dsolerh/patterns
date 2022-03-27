package structural

import "fmt"

type InterfaceX interface {
	MethodA()
	AddChild(InterfaceX)
}

type Composite struct {
	children []InterfaceX
}

func (c *Composite) MethodA() {
	if len(c.children) == 0 {
		fmt.Println("I'm a leaf ")
		return
	}
	fmt.Println("I'm a composite ")
	for _, child := range c.children {
		child.MethodA()
	}
}

func (c *Composite) AddChild(child InterfaceX) {
	c.children = append(c.children, child)
}
