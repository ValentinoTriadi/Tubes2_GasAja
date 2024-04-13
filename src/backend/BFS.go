package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func bfsScrape2(webEntity web, keyword string, count *int, base *string, saveRes *[][]web) {
	// Initialize local variable 
	var storage = [][]ResultEntity{}
	var currentLevel = []ResultEntity{}
	var nextLevel = []ResultEntity{}
	var found = false
	var BASEURL = *base
	var level = 0



	// Initialize storage
	temp := ResultEntity{0, webEntity}
	storage = append(storage, []ResultEntity{temp})

	// Start BFS scraping until found
	for !found {
		// Get current level
		currentLevel = storage[level]

		// Loop through current level
		for index, res := range currentLevel {

			// temporary code for testing purpose
			fmt.Println("Scraping: ", BASEURL + res.webEntity.Url)
			// fmt.Println("Level: ", level, "Index: ", index)

			// Send a GET request to the URL
			(*count)++ // Increment the total number of web entities scraped
			response, err := http.Get(BASEURL + res.webEntity.Url)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			// Parse the HTML response
			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Slice to keep all the hyperlinks and their titles
			var webs []web

			// Find all anchor tags in the HTML document
			doc.Find("div.mw-body-content").Find("a").Each(func(i int, s *goquery.Selection) {
				// Get the href attribute value
				href, exists := s.Attr("href")
				if exists {
					// Get the title of the hyperlink
					title, texists := s.Attr("title")
					if texists {

						// Add the hyperlink and title to the webs slice if href start with "/wiki/" and does not contain ":"
						if strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":") {

							// Add the hyperlink and title to the webs slice if it is not already in the slice and storage to prevent duplicate and loop
							if !containsWebEntity(web{href, title}, webs) && !isStorageContainsWebEntity(web{href, title}, &storage){
								webs = append(webs, web{href, title})
							}

							// Check if the title suit keyword
							if title == keyword {
								found = true

								// Add to result
								appendToResult(&storage, level, index, web{href, title}, saveRes)
								// fmt.Println("Result: ", *saveRes)
							}
						}
					}
				}
			})

			// Add webs to next level
			for _, w := range webs {
				nextLevel = append(nextLevel, ResultEntity{index, w})
			}
		}

		
		storage = append(storage, nextLevel) 	// Add next level to storage
		level++ 								// Increment level
		nextLevel = []ResultEntity{} 			// empty next level
	}
}
