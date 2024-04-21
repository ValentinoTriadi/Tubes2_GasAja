package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	// "github.com/PuerkitoBio/goquery"
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
	Lang	string `json:"lang"`
}

type ResultEntity struct {
	index 		int		// index of parent
	webEntity 	web
}

var GlobalLimit int 

func main() {
	// Create Router
    router := mux.NewRouter()

	// Handle route
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/api/scrape/bfs", bfsScrapeHandler).Methods("POST")

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


/* ============================================================================ */
/* =================================BFS Scrape================================= */
/* ============================================================================ */

func bfsScrapeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start BFS scraping...")
	
	// Var lokal
	var count int // Total number of web entities scraped
	var res [][]web // Slice to keep the result of the scraping

	// Decode the request body into an Input struct
	var i Input
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Base URL
	BASEURL := "https://" + i.Lang + ".wikipedia.org"

	// Start scraping
	timeStart := time.Now()
	bfsScrape2(web{"/wiki/" + strings.ReplaceAll(i.Start, " ", "_"), i.Start}, i.Keyword, &count, &BASEURL, &res)
	timeEnd := time.Now()

	// Encode the result into a struct and send it as a response
	result := struct {
		Webs    [][]web
		Time    string
		Total   int
	}{
		Webs:    res,
		Time:    timeEnd.Sub(timeStart).String(),
		Total:   count,
	}
	
	fmt.Println("End scraping...")

	json.NewEncoder(w).Encode(result)
}

/* ============================================================================ */
/* =================================BFS Scrape================================= */
/* ============================================================================ */