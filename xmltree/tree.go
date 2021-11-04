package xmltree

import (
	"encoding/xml"
	"fmt"
)

type Node interface{} // Chardata or *Element

type CharData string

type Element struct {
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}

func (e Element) String() string {
	output := fmt.Sprintf(
		"Type: %s, Attrs: %s, Children: \n\t%s", 
		e.Type, 
		e.Attr,
		e.Children,
	)
	return output
}