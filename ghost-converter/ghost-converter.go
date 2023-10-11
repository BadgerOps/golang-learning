package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type GhostExport struct {
	Db []struct {
		Data struct {
			Posts []struct {
				Title   string `json:"title"`
				Content string `json:"html"`
				Slug    string `json:"slug"`
			} `json:"posts"`
		} `json:"data"`
	} `json:"db"`
}

func displayHelp() {
	fmt.Println(`Ghost Export to Static HTML Converter

Usage:
  -f string
        Path to the Ghost JSON export.
  -o string
        Output directory for the static HTML files.

Example:
  go run main.go -f export.json -o outputDir

This tool takes a Ghost CMS JSON export and converts each post into a static HTML file. The resulting HTML files will be saved in the provided output directory.`)
}

func main() {
	filePath := flag.String("f", "", "path to Ghost JSON export")
	outDir := flag.String("o", "", "output directory for the static HTML")
	help := flag.Bool("h", false, "display help")
	flag.BoolVar(help, "help", false, "display help")

	flag.Parse()

	if *help || (*filePath == "" && *outDir == "") {
		displayHelp()
		return
	}

	// Read the JSON file
	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var export GhostExport
	err = json.Unmarshal(data, &export)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	if len(export.Db) == 0 || len(export.Db[0].Data.Posts) == 0 {
		fmt.Println("No posts found in the Ghost export.")
		return
	}

	// Ensure the output directory exists
	err = os.MkdirAll(*outDir, 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	for _, post := range export.Db[0].Data.Posts {
		filename := filepath.Join(*outDir, post.Slug+".html")
		content := fmt.Sprintf("<html><head><title>%s</title></head><body>%s</body></html>", post.Title, post.Content)
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Written: %s\n", filename)
		}
	}
}
