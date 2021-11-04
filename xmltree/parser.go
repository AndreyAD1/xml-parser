package xmltree

import (
	"encoding/xml"
	"fmt"
	"io"
)

func GetXMLTree(reader io.Reader) (Element, error) {
	dec := xml.NewDecoder(reader)
	var element Element
	tok, err := dec.Token()
	if err == io.EOF {
		return element, nil
	} else if err != nil {
		return element, fmt.Errorf("xmlselect: %v", err)
	}
	el := tok.(xml.StartElement)
	element.Type = el.Name
	element.Attr = el.Attr
	tree, err := getNode(dec, element)
	if err != nil {
		return Element{}, err
	}
	return tree, nil
}

var emptyElement Element

func getNode(dec *xml.Decoder, element Element) (Element, error) {
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
			nextElement := Element{
				Type: tok.Name,
				Attr: tok.Attr,
			}
			childNode, err := getNode(dec, nextElement)
			if err != nil {
				return element, err
			}
			element.Children = append(element.Children, childNode)
		}
	}
}
