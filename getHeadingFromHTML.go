package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}
	result := doc.Find("h1, h2").First().Text()

	return strings.TrimSpace(result)
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	// Prefer the first <p> inside <main>
	p := doc.Find("main p").First()

	// Fall back to the first <p> anywhere in the document
	if p.Length() == 0 {
		p = doc.Find("p").First()
	}

	return strings.TrimSpace(p.Text())
}
