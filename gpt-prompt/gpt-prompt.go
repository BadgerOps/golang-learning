package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	openAIAPIURL = "https://api.openai.com/v1/engines/davinci/completions"
	configFile   = "~/.config/gpt.conf"
)

type Config struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type APIRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type APIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func readConfig() (Config, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	configPath := strings.Replace(configFile, "~", userHome, 1)

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func getResponseFromGPT3(prompt string, token string) (string, error) {
	requestBody, err := json.Marshal(APIRequest{
		Prompt:    prompt,
		MaxTokens: 150,
	})

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response APIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Text, nil
	}
	return "", fmt.Errorf("No response from GPT-3.5")
}

func main() {
	var prompt string
	flag.StringVar(&prompt, "p", "", "Prompt for GPT-3.5")
	flag.Parse()

	if prompt == "" {
		fmt.Println("Error: Please provide a prompt using the -p flag.")
		return
	}

	config, err := readConfig()
	if err != nil {
		fmt.Printf("Error reading config: %s\n", err)
		return
	}

	response, err := getResponseFromGPT3(prompt, config.Token)
	if err != nil {
		fmt.Printf("Error getting response from GPT-3.5: %s\n", err)
		return
	}

	fmt.Println("GPT-3.5 Response:", response)
}
