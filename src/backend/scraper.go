package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type web struct {
	url, title string
}

func main() {
	// // input
	// var keyword, start string
	// var limit int

	// // Get the keyword
	// fmt.Print("Enter the keyword: ")
	// fmt.Scanf("%v", &keyword)

	// // Get the start page
	// fmt.Println("Enter the start keyword: ")
	// fmt.Scanf("%v", &start)

	// // Get the limit
	// fmt.Println("Enter the limit: ")
	// _, err := fmt.Scanf("%d", &limit)
	// if err != nil {
	// 	log.Fatal(err)
	// }


	// temp declaration
	keyword := "Bahasa Jawa"
	start := "Kucing"
	limit := 2

	// Start scraping
	timeStart := time.Now()
	web := getWeb(web{"/wiki/" + start, start}, keyword, limit, []web{})
	timeEnd := time.Now()

	// print the result
	fmt.Println("Result:")
	for _, w := range web {
		fmt.Print(w.title)
		if w.title != keyword {
			fmt.Print(" -> ")
		}
	}
	println()
	println("Executed in", timeEnd.Sub(timeStart).Seconds(), "seconds")
}

func getWeb(webEntity web, keyword string, limit int, Res []web) []web {

	if limit == 0 {
		return Res
	}

	// Base URL
	BASEURL := "https://id.wikipedia.org"

	// Found Condition
	found := false

	// Send a GET request to the URL
	response, err := http.Get(BASEURL + webEntity.url)
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

				// Add the hyperlink and title to the webs slice if title start with "/wiki/"
				if strings.HasPrefix(href, "/wiki/") {
					webs = append(webs, web{href, title})

					// Check if the title suit keyword
					if title == keyword {
						found = true
						Res = append(Res, webEntity)
						Res = append(Res, web{href, title})
					}
				}

			}
		}
	})

	if !found {
		// call getweb with all hyperlink in webs
		Res = append(Res, webEntity)
		for _, w := range webs {
			get := getWeb(w, keyword, limit-1, Res)
			if get[len(get)-1].title == keyword {
				Res := get
				return Res
			}
		}
	}

	return Res
}
