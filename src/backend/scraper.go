package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type web struct {
	Url		string `json:"url"`
	Title 	string `json:"title"`
}

type Input struct {
	Keyword string `json:"keyword"`
	Start   string `json:"start"`
	Limit   int    `json:"limit"`
}

func main() {
	// Create Router
    router := mux.NewRouter()

	// Handle route
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/api/scrape", scrapeHandler).Methods("POST")

	enchancedRouter := enableCORS(jsonContentTypeMiddleware(router))

	log.Fatal(http.ListenAndServe(":8000", enchancedRouter))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start scraping...")
	
	var i Input
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	// Start scraping
	timeStart := time.Now()
	webs := getWeb(web{"/wiki/" + i.Start, i.Start}, i.Keyword, i.Limit, []web{})
	timeEnd := time.Now()

	result := struct {
		Webs    []web
		Time    string
	}{
		Webs:    webs,
		Time:    timeEnd.Sub(timeStart).String(),
	}
	
	fmt.Println(result)
	fmt.Println("End scraping...")

	json.NewEncoder(w).Encode(result)
}

func containsWebEntity (webEntity web, Res []web) bool {
	for _, w := range Res {
		if w.Url == webEntity.Url {
			return true
		}
	}
	return false
}

func getWeb(webEntity web, keyword string, limit int, Res []web) []web {

	if limit == 0 || containsWebEntity(webEntity, Res) {
		return Res
	}

	// Base URL
	BASEURL := "https://id.wikipedia.org"

	// Found Condition
	found := false

	// fmt.Println("Scraping: ", BASEURL+webEntity.Url)

	// Send a GET request to the URL
	response, err := http.Get(BASEURL + webEntity.Url)
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
			if get[len(get)-1].Title == keyword {
				Res := get
				return Res
			}
		}
	}

	return Res
}
