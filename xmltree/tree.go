package xmltree

import "encoding/xml"

type Node interface{} // Chardata or *Element

type CharData string

type Element struct {
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}