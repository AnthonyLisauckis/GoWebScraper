package main

import (
	"fmt"

	colly "github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	fmt.Printf("Collector Created")
	// Your scraping logic will be added here

	// Define a callback for when an <a> tag is found
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// fmt.Printf("OnHTML Function Called")
		// Extract the title and URL from the <a> tag
		title := e.Text
		link := e.Attr("href")
		fmt.Printf("Title: %s, URL: %s\n", title, link)
	})

	// Start the scraping process by visiting the website
	c.Visit("https://google.com")
	fmt.Printf("Visit Func Called")
}
