package main

import (
	"fmt"
	"github.com/AndreyAD1/xml-parser/xmltree"
	"log"
	"os"
)

func getXmlString(tree xmltree.Node, tabNumber int) string {
	var tabs string
	for i := 0; i < tabNumber; i++ {
		tabs += "\t"
	}
	if node, ok := tree.(xmltree.CharData); ok {
		return fmt.Sprintf("%s %s\n", tabs, node)
	}
	node := tree.(xmltree.Element)
	result := fmt.Sprintf("%s<%s %s>\n", tabs, node.Type.Space, node.Type.Local)
	for _, child := range node.Children {
		result += getXmlString(child, tabNumber + 1)
	}
	result += fmt.Sprintf("%s<%s %s>\n", tabs, node.Type.Space, node.Type.Local)
	return result
}

func printXMLTree(tree xmltree.Element) {
	fmt.Print(getXmlString(tree, 0))
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("The script expects an argument: file path of XML file")
	}
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	xmlTree, err := xmltree.GetXMLTree(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	printXMLTree(xmlTree)
}
