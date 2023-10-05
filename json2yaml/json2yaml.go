package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// This program is just me messing around with converting json to yaml
// It is just for learning, don't actually use it. Please.

func main() {
	// Define and parse the command line flags
	var filePath, outputFilename string
	flag.StringVar(&filePath, "f", "", "Path to the JSON file")
	flag.StringVar(&outputFilename, "o", "output.yaml", "Path to the output YAML file")
	flag.Parse()

	if filePath == "" {
		log.Fatal("Please specify a JSON file with the -f flag")
	}

	// Read the JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JSON file %s: %v", filePath, err)
	}

	var jsonObj interface{}

	// Unmarshal the JSON data
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Convert JSON object to YAML
	yamlData, err := yaml.Marshal(jsonObj)
	if err != nil {
		log.Fatalf("Error marshaling to YAML: %v", err)
	}

	// Write YAML to the specified output file
	if err := ioutil.WriteFile(outputFilename, yamlData, 0644); err != nil {
		log.Fatalf("Error writing to %s: %v", outputFilename, err)
	}

	fmt.Printf("Data written to %s\n", outputFilename)
}
