package xmltree

import (
	"encoding/xml"
	"fmt"
	"io"
)

func GetXMLTree(reader io.Reader) (Element, error) {
	dec := xml.NewDecoder(reader)
	tree, err := getNode(dec)
	if err != nil {
		return Element{}, err
	}
	return tree, nil
}

func getNode(dec *xml.Decoder) (Element, error) {
	element := Element{}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return element, nil
		} else if err != nil {
			return element, fmt.Errorf("xmlselect: %v", err)
		}
		switch tok := tok.(type) {
		case xml.EndElement:
			return element, nil
		case xml.CharData:
			element.Children = append(element.Children, CharData(tok))
		case xml.StartElement:
			element.Type = tok.Name
			element.Attr = tok.Attr
			childNode, err := getNode(dec)
			if err != nil {
				return element, err
			}
			element.Children = append(element.Children, childNode)
		}
	}
}
