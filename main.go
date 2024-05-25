package main

import (
	"flag"
	"funnyJSON/JSONExplorer"
	"log"
)

func main() {

	// 命令行参数
	var jsonFilePath, styleArg, iconFamily string
	flag.StringVar(&jsonFilePath, "f", "", "path to the JSON file")
	flag.StringVar(&styleArg, "s", "", "style for the JSON visualization(tree in default)")
	flag.StringVar(&iconFamily, "i", "", "icon family for the JSON visualization(null in default)")
	flag.Parse()
	if jsonFilePath == "" {
		log.Fatalf("Error: '-f' flag is required but not provided.")
	}
	var drawer JSONExplorer.Explorer

	drawer = &JSONExplorer.Drawer{}
	err := drawer.ParseJSON(jsonFilePath)
	if err != nil {
		log.Fatalf("ReadFile error: %v", err)
	}

	if iconFamily != "" {
		drawer.InitIcon(iconFamily)
	}
	if styleArg != "" {
		drawer.InitStyle(styleArg)
	}
	drawer.Show()
}
