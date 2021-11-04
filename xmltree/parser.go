package xmltree

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func GetXMLTree(reader io.Reader) (Element, error) {
	dec := xml.NewDecoder(reader)
	tok, err := dec.Token()
	if err == io.EOF {
		return Element{}, err
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
	initialToken, ok := tok.(xml.StartElement)
	if !ok {
		return Element{}, fmt.Errorf("Invalid XML")
	}
	initialElement := Element{
		Type: initialToken.Name, 
		Attr: initialToken.Attr,
	}
	tree, err := getNode(dec, initialElement)
	if err != nil {
		return Element{}, err
	}
	return tree, nil
}

func getNode(dec *xml.Decoder, element Element) (Element, error) {
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return element, nil
		} else if err != nil {
			return element, fmt.Errorf("xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.EndElement:
			return element, nil
		case xml.CharData:
			element.Children = append(element.Children, CharData(tok))
			return element, nil
		case xml.StartElement:
			childElement := Element{
				Type: tok.Name, 
				Attr: tok.Attr,
			}
			childNode, err := getNode(dec, childElement)
			if err != nil {
				return element, err
			}
			element.Children = append(element.Children, childNode)
		}
	}
}