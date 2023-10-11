package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type GhostExport struct {
	Db []struct {
		Data struct {
			Posts []struct {
				Title     string    `json:"title"`
				Content   string    `json:"html"`
				Slug      string    `json:"slug"`
				CreatedAt time.Time `json:"created_at"`
			} `json:"posts"`
		} `json:"data"`
	} `json:"db"`
}

func displayHelp() {
	fmt.Println(`Ghost Export to Hugo Content Converter

Usage:
  -f string
        Path to the Ghost JSON export.
  -o string
        Output directory for the Hugo content.

Example:
  go run main.go -f export.json -o contentDir

This tool takes a Ghost CMS JSON export and converts each post into Hugo-compatible content. The resulting files will be saved in the provided output directory.`)
}

func main() {
	filePath := flag.String("f", "", "path to Ghost JSON export")
	outDir := flag.String("o", "", "output directory for the Hugo content")
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
		// Hugo front matter (TOML format in this case, but can be changed to YAML if preferred)
		frontMatter := fmt.Sprintf(`+++
title = "%s"
date = "%s"
draft = false
+++

`, post.Title, post.CreatedAt.Format(time.RFC3339))

		content := frontMatter + post.Content
		filename := filepath.Join(*outDir, post.Slug+".md")
		err = os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Written: %s\n", filename)
		}
	}
}
