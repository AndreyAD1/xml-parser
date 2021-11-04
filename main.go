package main

import (
	"fmt"
	"log"
	"os"
	"github.com/AndreyAD1/xml-parser/xmltree"
)

func printXMLTree(tree xmltree.Element) {
	fmt.Println(tree)
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