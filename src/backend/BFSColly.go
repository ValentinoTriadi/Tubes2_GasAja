package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"sync"

	"github.com/gocolly/colly/v2"
)


func gocollyScrapeBase(w web, keyword string, BASEURL string, saveRes *[][]web, allWebs *[]web) {
	// Initialize local variable
	timeStart := time.Now()
	var storage = [][]ResultEntity{}
	var nextLevel = []ResultEntity{}
	var found = false
	var level = 0

	// Initialize storage
	temp := ResultEntity{0, w}
	storage = append(storage, []ResultEntity{temp})
	if w.Title == keyword{
		*saveRes = [][]web{{w}}
		return
	}

	var wg sync.WaitGroup

	go tb.AddTokens(MaxToken, 1*time.Second) // Tambahkan Max token setiap detik

	// Start BFS scraping until found
	for !found {

		ch := make(chan ResultEntity)

		// Loop through current level
		for index, res := range storage[level] {
			wg.Add(1)
			go gocollyScrape(res.webEntity, keyword, BASEURL, index, &found, level, saveRes, &storage, ch, &wg, allWebs)

			// Limit time to 5 minutes
			if time.Since(timeStart) > 5*time.Minute {
				fmt.Println("Time limit reached")
				break
			}
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		// Add webs to next level
		for w := range ch {
			nextLevel = append(nextLevel, w)
		}

		storage = append(storage, nextLevel) // Add next level to storage
		level++                              // Increment level
		nextLevel = []ResultEntity{}         // empty next level

		// Limit time to 5 minutes
		if time.Since(timeStart) > 5*time.Minute {
			fmt.Println("Time limit reached")
			return
		}
	}
}

func gocollyScrape(w web, keyword string, BASEURL string, index int, found *bool, level int, saveRes *[][]web, storage *[][]ResultEntity, ch chan<- ResultEntity, wg *sync.WaitGroup, allWebs *[]web) {

	defer wg.Done()

	// Tunggu sampai token tersedia
	for !tb.Consume() {
		time.Sleep(100 * time.Millisecond) // Tidur selama durasi singkat jika token tidak tersedia
	}

	var webs []web
	
	// Instantiate default collector
	c := colly.NewCollector(
		// colly.AllowedDomains(BASEURL),
		// colly.MaxDepth(3),
		colly.Async(true),
		colly.URLFilters(
			regexp.MustCompile(BASEURL + `/wiki/[^:]*$`),
		),
	)

	c.CacheDir = "./cache"

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})

	c.SetRequestTimeout(100 * time.Second)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.Attr("title")
		if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") && !strings.Contains(link, "%3A"){

			if !containsWebEntity(web{link, title}, webs) {
				webs = append(webs, web{link, title})
				ch <- ResultEntity{index, web{link, title}}
			}
			if (title == keyword) {
				*found = true
				appendToResult(storage, level, index, web{link, title}, saveRes)
			}
		}
	})

	c.OnResponse(func(r *colly.Response) {
		if !containsWebEntity(web{r.Request.URL.Path, r.Request.URL.Path}, *allWebs) {
			*allWebs = append(*allWebs, web{r.Request.URL.Path, r.Request.URL.Path})
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraped", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.Request.URL, err)
	})

	// Start scraping
	c.Visit(BASEURL + w.Url)

	c.Wait()
}