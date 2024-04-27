package main

import (
	//"encoding/json"
	"fmt"
	//"log"
	//"net/http"
	"strings"
	//"time"
	"regexp"
	//"sync"

	"github.com/gocolly/colly/v2"
	//"github.com/gorilla/mux"
)


type Node struct {
	Value    web
	Depth    int
	Children []*Node
}

type Tree struct {
	Root Node
}

func expandTree(node *Node, baseURL string, limit int, count *int, allWebs *[]web) {
	if node.Depth >= limit {
		return
	}

	c := colly.NewCollector(
		colly.Async(true),
		colly.URLFilters(
			regexp.MustCompile(baseURL + `/wiki/[^:]*$`),
		),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})

	c.CacheDir = "./cache"
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		title := e.Attr("title")
		if strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":") {
			childWeb := web{Url: href, Title: title}
			childNode := Node{Value: childWeb, Depth: node.Depth + 1}
			node.Children = append(node.Children, &childNode)
			if node.Depth+1 < limit {
				expandTree(&childNode, baseURL, limit, count,allWebs)
			}
		}
	})

	c.OnResponse(func(r *colly.Response) {
		if !containsWebEntity(web{r.Request.URL.Path,r.Request.URL.Path},*allWebs){
			*allWebs = append(*allWebs,web{r.Request.URL.Path,r.Request.URL.Path})
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraped", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.Request.URL, err)
	})

	c.Visit(baseURL + node.Value.Url)
	c.Wait()
}

func findSolution(node *Node, keyword string, count *int, baseURL string, webs *[]web, allWebs *[]web) {
	limit := 1
	found := false
	path := []web{}
	for !found {
		expandTree(node, baseURL, limit, count, allWebs)
		found = searchSolution(node, keyword, &path, webs)
		limit++
	}
}

func searchSolution(node *Node, keyword string, currentPath *[]web, webs *[]web) bool {
	*currentPath = append(*currentPath, node.Value)
	if node.Value.Title == keyword {
		reversePath := make([]web, len(*currentPath))
		copy(reversePath, *currentPath)
		*webs = append(*webs, reversePath...)
		return true
	}
	for _, child := range node.Children {
		if searchSolution(child, keyword, currentPath, webs) {
			return true
		}
	}
	*currentPath = (*currentPath)[:len(*currentPath)-1]
	return false
}


