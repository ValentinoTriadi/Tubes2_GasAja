package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

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

	var wg sync.WaitGroup
	
	// Start BFS scraping until found
	for !found {
		// Get current level
		currentLevel = storage[level]

		ch := make(chan ResultEntity)

		// Loop through current level
		for index, res := range currentLevel {
			wg.Add(1)
			go bfsScrape(BASEURL + res.webEntity.Url, count, &storage, keyword, &found, level, index, saveRes, ch, &wg)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		// Add webs to next level
		for w := range ch {
			nextLevel = append(nextLevel, w)
		}

		storage = append(storage, nextLevel) 	// Add next level to storage
		level++ 								// Increment level
		nextLevel = []ResultEntity{} 			// empty next level
	}
}


func bfsScrape(URL string, count *int, storage *[][]ResultEntity, keyword string, found *bool, level int, index int, saveRes *[][]web, ch chan<- ResultEntity, wg *sync.WaitGroup) []web {

	defer wg.Done()

	// temporary code for testing purpose
	fmt.Println("Scraping: ", URL)
	// fmt.Println("Level: ", level, "Index: ", index)

	// Send a GET request to the URL
	(*count)++ // Increment the total number of web entities scraped
	response, err := http.Get(URL)
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
					if !containsWebEntity(web{href, title}, webs) && !isStorageContainsWebEntity(web{href, title}, storage){
						webs = append(webs, web{href, title})
						ch <- ResultEntity{index, web{href, title}}
					}

					// Check if the title suit keyword
					if title == keyword {
						*found = true

						// Add to result
						appendToResult(storage, level, index, web{href, title}, saveRes)
						// fmt.Println("Result: ", *saveRes)
					}
				}
			}
		}
	})

	return webs
}