package xmltree

import (
	"encoding/xml"
	"fmt"
)

type Node interface{} // Chardata or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	output := fmt.Sprintf("<%s %s>", e.Type.Local, e.Attr)
	return output
}
