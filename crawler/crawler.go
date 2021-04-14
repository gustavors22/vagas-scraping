package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func Crawler(url string){
	domain := "programathor.com.br"
	id := 1

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnHTML(`.cell-list`, func(e *colly.HTMLElement){
		link := e.ChildAttr("a","href")
		h3 := e.ChildText("h3")

		fmt.Printf("\nVaga %d -> %s", id, h3)
		fmt.Printf(" -> %s%s\n", domain,link)
		id++
	})

	c.OnHTML(`a.page-link`, func(e *colly.HTMLElement){
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\n\nVisiting %s\n", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit(url)
}