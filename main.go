package main

import (
	"flag"
	"fmt"
	"texttool/internal/finder"
)

func main() {
	root := flag.String("path", ".", "to find file")
	ext := flag.String("ext", ".go", "to filter file")
	flag.Parse()
	
	file := finder.SearchFile(root, ext)
	if file == nil {
		fmt.Println("tidak ada file yang dicari")
	}
	for _, f := range file {
		fmt.Println(f)
	}
}