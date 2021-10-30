package main

import (
	"fmt"
	"os"
	"github.com/AndreyAD1/xml-parser/xmltree"
)

func printXMLTree(tree xmltree.Element) {
	fmt.Println("draft")
}


func main() {
	xmlTree := xmltree.GetXMLTree(os.Stdin)
	printXMLTree(xmlTree)
}