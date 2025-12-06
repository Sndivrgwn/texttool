package finder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SearchFile(root *string, ext *string) []string {
	var file []string

	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == *ext {
			file = append(file, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return file
}