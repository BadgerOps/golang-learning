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
				Content   string    `json:"mobiledoc"`
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
		var mobiledoc struct {
			Cards [][]interface{} `json:"cards"`
		}
		if err := json.Unmarshal([]byte(post.Content), &mobiledoc); err != nil {
			fmt.Printf("Error extracting content for post %s: %v\n", post.Title, err)
			continue
		}

		var markdownContent string
		for _, card := range mobiledoc.Cards {
			cardType, ok := card[0].(string)
			if !ok {
				fmt.Printf("Unexpected card type format for post %s\n", post.Title)
				continue
			}

			switch cardType {
			case "markdown":
				content, _ := card[1].(map[string]interface{})["markdown"].(string)
				markdownContent += content
			case "hr":
				markdownContent += "\n---\n"
			case "code":
				codeContent, ok := card[1].(map[string]interface{})["code"].(string)
				language, langOk := card[1].(map[string]interface{})["language"].(string)
				if ok {
					if !langOk {
						language = ""
					}
					markdownContent += fmt.Sprintf("\n```%s\n%s\n```\n", language, codeContent)
				}
			case "image":
				imageURL, ok := card[1].(map[string]interface{})["src"].(string)
				altText, altOk := card[1].(map[string]interface{})["alt"].(string)
				if ok {
					if !altOk {
						altText = ""
					}
					markdownContent += fmt.Sprintf("\n![%s](%s)\n", altText, imageURL)
				}
			default:
				fmt.Printf("Unhandled card type '%s' for post %s\n", cardType, post.Title)
			}
		}

		if markdownContent == "" {
			fmt.Printf("No content extracted for post %s\n", post.Title)
			continue
		}

		// Hugo front matter (TOML format)
		frontMatter := fmt.Sprintf(`+++
title = "%s"
date = "%s"
draft = false
+++

`, post.Title, post.CreatedAt.Format(time.RFC3339))

		content := frontMatter + markdownContent
		filename := filepath.Join(*outDir, post.Slug+".md")
		err = os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Written: %s\n", filename)
		}
	}
}
