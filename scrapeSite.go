package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func scrapeWebsite(url string, c chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal("Error:", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("Error:", err)
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		dataSlice := make([]string, 0, 5) 
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			dataSlice = append(dataSlice, s.Text())
		})
		if len(dataSlice) == 5 {
			c <- dataSlice
		}
	})
}

func scrapeAndSaveToFile() {
	pages := []string{
		"https://atts.org/breed-statistics/statistics-page1/",
		"https://atts.org/breed-statistics/statistics-page2/",
		"https://atts.org/breed-statistics/statistics-page3/",
		"https://atts.org/breed-statistics/statistics-page4/",
		"https://atts.org/breed-statistics/statistics-page5/",
		"https://atts.org/breed-statistics/statistics-page6/",
		"https://atts.org/breed-statistics/statistics-page7/",
		"https://atts.org/breed-statistics/statistics-page8/",
	}
	c := make(chan []string, len(pages))
	var wg sync.WaitGroup
	wg.Add(len(pages))
	for _, page := range pages {
		go scrapeWebsite(page, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	bts := BreedTestScoreSlice{}
	for value := range c {
		cs := CreateScoreFromSlice(value)
		bts = append(bts, *cs)
	}

	bts.saveToFile("scores.txt")
}