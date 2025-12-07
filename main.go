package main

import (
	"flag"
	"fmt"
	"sort"
	"texttool/internal/finder"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type FileInfo struct {
	Name, Path string
	Size int64
	Mod time.Time
}

func main() {
	var results []FileInfo

	root := flag.String("path", ".", "to find file")
	ext := flag.String("ext", ".go", "to filter file")
	filename := flag.String("filename", "", "to filter file")
	sortKey := flag.String("sort", "name", "sort by key")
	order := flag.String("order", "asc", "asc|desc")

	flag.Parse()

	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	
	t.AppendHeader(table.Row{"File Name", "File Size", "Last Modified", "File Path"})
	file := finder.SearchFile(root, ext, filename)
	
	for _, f := range file {
		size, mod, name := finder.GetFileStat(f)
		modTime, _ := time.Parse("2006-01-02 15:04:05", mod)
		results = append(results, FileInfo{
			Name: name,
			Size: size,
			Mod: modTime,
			Path: f,
		})
	}
	
	sort.Slice(results, func(i, j int) bool {
    var less bool

    switch *sortKey {
    case "size":
        less = results[i].Size < results[j].Size
    case "mod":
        less = results[i].Mod.Before(results[j].Mod)
    default:
        less = results[i].Name < results[j].Name
    }

    if *order == "desc" {
        return !less
    }
    return less
	})

	
	for _, r := range results {
		sizeFix := finder.HumanFileSize(float64(r.Size))
		t.AppendRow(table.Row{
		color.GreenString(r.Name),
		color.MagentaString(sizeFix),
		color.YellowString(r.Mod.Format("2006-01-02 15:04:05")),
		color.CyanString(r.Path),
	})
	}

	if file == nil {
		t.AppendRow(table.Row{color.RedString("File Not found")})
	}

	fmt.Println(t.Render())
}