package finder

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func SearchFile(root *string, ext *string, filename *string) []string {
	var file []string
	targetFileName := strings.ToLower(*filename)


	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if targetFileName == "" {
			if !info.IsDir() && filepath.Ext(path) == *ext {
				file = append(file, path)
			}
		} else {
			if !info.IsDir() {
				currentFileName := strings.ToLower(info.Name())
				if strings.Contains(currentFileName, targetFileName) {
					file = append(file, path)
				}
			}
		}



		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func GetFileStat(root string) (int64, string, string) {
	file, err := os.Open(root)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	fileLastModified := fileInfo.ModTime().Format("2006-01-02 15:04:05")
	fileName := fileInfo.Name()

	return fileSize, fileLastModified, fileName
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {
	if size <= 0 {
		return "0 B"
	}

	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	base := math.Log(size) / math.Log(1024)
	idx := int(math.Floor(base))

	if idx < 0 {
		idx = 0
	}
	if idx >= len(suffixes) {
		idx = len(suffixes) - 1
	}

	scaled := size / math.Pow(1024, float64(idx))
	scaled = Round(scaled, 0.5, 2)

	return strconv.FormatFloat(scaled, 'f', -1, 64) + " " + suffixes[idx]
}

