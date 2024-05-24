package main

import (
	"flag"
	"funnyJSON/builder"
	"log"
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

	drawer := builder.Drawer{}
	err := drawer.ParseJSON(jsonFilePath)
	if err != nil {
		log.Fatalf("ReadFile error: %v", err)
	}

}
