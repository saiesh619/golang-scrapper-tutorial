package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var URL = "https://www.metalsucks.net/"

func main() {
	req, _ := http.NewRequest(
		"GET",
		URL,
		nil,
	)

	resDo, _ := http.DefaultClient.Do(req)
	parseResponse(resDo)

}

func parseResponse(res *http.Response) {
	fmt.Println(res.StatusCode)
	document, _ := goquery.NewDocumentFromReader(res.Body)

	nodes := document.Find("div")
	fmt.Println("Text of all div elements is here ========", nodes.Text())

}
