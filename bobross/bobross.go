package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getBob() string {
	bobURL := string "https://www.bobrossquotes.com/text.php"
	reqURL := fmt.Sprintf("%v", bobURL)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error woth http Client: %s\n", err)
		os.Exit(1)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading body of request: %s\n", err)
		os.Exit(1)
	}
	return string(resBody)
}

func main() {
	bobQ := getBob()
	fmt.Printf("%v", bobQ)
}
