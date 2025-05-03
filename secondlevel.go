package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	// Fetch HTML
	resp, err := http.Get("https://books.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the HTML content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Regex to extract title attributes
	re := regexp.MustCompile(`title="([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(body), -1)

	// Print book titles
	for _, match := range matches {
		if len(match) > 1 {
			fmt.Println(match[1])
		}
	}
}
