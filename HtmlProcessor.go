package iqdbAPI

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func HtmlProcessor(res *http.Response) string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resp = Response{
		Success: false,
		Results: nil,
	}
	// Find the review items
	doc.Find("#pages.pages div table").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			// For each item found, get the title
			resultUrl, _ := s.Find(".image a").Attr("href")
			resultHead := s.Find("tr th").Text()
			resultTitles, _ := s.Find(".image img").Attr("title")
			split := strings.Fields(s.Find("tr:nth-child(4) td").Text())
			split0 := strings.Split(split[0], "×")
			if len(split) <= 1 || len(split0) <= 1 {
				return
			}
			resultHeight := split0[0]
			resultWidth := split0[1]
			resultCategory := split[1]
			resultSimilarity := s.Find("tr:nth-child(5) td").Text()

			resp.Results = append(resp.Results, Result{
				Head:       resultHead,
				Url:        resultUrl,
				Titles:     resultTitles,
				Height:     resultHeight,
				Width:      resultWidth,
				Category:   resultCategory,
				Similarity: resultSimilarity,
			})
		}
	})
	resp.Success = true
	respP := &resp
	j, err := json.Marshal(respP)
	if err != nil {
		fmt.Println(err)
	}
	return string(j)
}
