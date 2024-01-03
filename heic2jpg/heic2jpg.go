package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func convertHeicToJpg(inputPath string, outputPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Construct the output filename with .jpg extension
	outputFile := filepath.Join(outputPath, filepath.Base(inputPath))
	outputFile = outputFile[0:len(outputFile)-len(filepath.Ext(outputFile))] + ".jpg"

	// Construct the ImageMagick convert command
	cmd := exec.Command("convert", inputPath, outputFile)

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error converting file %s: %v\n", inputPath, err)
		return
	}

	fmt.Printf("Converted %s to %s\n", inputPath, outputFile)
}

func main() {
	// Define and parse command-line flags
	dirPtr := flag.String("d", "", "Directory containing HEIC files to convert")
	outPtr := flag.String("o", "", "Output directory for JPG files")
	flag.Parse()

	// Validate input directory
	inputDir, err := filepath.Abs(*dirPtr)
	if err != nil {
		fmt.Printf("Invalid input directory: %s\n", *dirPtr)
		os.Exit(1)
	}

	// Validate output directory
	outputDir, err := filepath.Abs(*outPtr)
	if err != nil {
		fmt.Printf("Invalid output directory: %s\n", *outPtr)
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Find all HEIC files in the input directory
	var wg sync.WaitGroup
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".HEIC" {
			wg.Add(1)
			go convertHeicToJpg(path, outputDir, &wg)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through input directory: %v\n", err)
		os.Exit(1)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Conversion complete.")
}
