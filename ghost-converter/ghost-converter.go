package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

func main() {
	filePath := flag.String("f", "", "path to Ghost JSON export")
	outDir := flag.String("o", "", "output directory for the static HTML")

	flag.Parse()

	if *filePath == "" || *outDir == "" {
		fmt.Println("Both -f and -o arguments are required.")
		return
	}

	// Read the JSON file
	data, err := ioutil.ReadFile(*filePath)
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
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Written: %s\n", filename)
		}
	}
}
