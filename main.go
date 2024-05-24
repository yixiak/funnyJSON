package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	// 命令行参数
	var jsonFilePath, styleArg, iconFamily string
	flag.StringVar(&jsonFilePath, "f", "", "path to the JSON file")
	flag.StringVar(&styleArg, "s", "tree", "style for the JSON visualization(tree in default)")
	flag.StringVar(&iconFamily, "i", "null", "icon family for the JSON visualization(null in default)")
	flag.Parse()
	if jsonFilePath == "" {
		log.Fatalf("Error: '-f' flag is required but not provided.")
	}
	file, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer file.Close()
	//fmt.Printf("f: %v; s: %v", *jsonFilePath, *styleArg)
}
