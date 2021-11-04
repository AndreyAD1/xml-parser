package xmltree

import (
	"io"
	"log"
	"os"
	"encoding/xml"
)

func GetXMLTree(reader io.Reader) Element {
	dec := xml.NewDecoder(reader)
	var initialElement Element
	tree := getNode(dec, initialElement)
	return tree
}

func getNode(dec *xml.Decoder, element Element) Element {
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return element
		} else if err != nil {
			log.Fatal(os.Stderr, "xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.EndElement:
			return element
		case xml.CharData:
			element.Children = append(element.Children, CharData(tok))
			return element
		case xml.StartElement:
			childElement := Element{
				Type: tok.Name, 
				Attr: tok.Attr,
			}
			childNode := getNode(dec, childElement)
			element.Children = append(element.Children, childNode)
		}
	}
}