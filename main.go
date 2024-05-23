package main

import (
	"flag"
)

func main() {

	// 命令行参数
	jsonFilePath := flag.String("f", "", "path to the JSON file")
	styleArg := flag.String("s", "tree", "style for the JSON visualization(tree in Default)")
	//iconFamily := flag.String("i","")
	flag.Parse()
	//fmt.Printf("f: %v; s: %v", *jsonFilePath, *styleArg)
}
