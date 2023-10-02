package main

import (
	"github.com/gocolly/colly"
	"log"
	"scrapper/bootstrap"
	"scrapper/config"
	mkWomenHandbagCrossbody "scrapper/handler/michael_kors/women/handbag/crossbody"
	"strings"
)

func init() {
	bootstrap.InitEnv("./.env")
	bootstrap.InitConfig()
	bootstrap.InitStorage()
}

// https://brightdata.com/blog/how-tos/web-scraping-go
func main() {
	c := colly.NewCollector()
	c.UserAgent = config.Get().UserAgent

	c.OnHTML(mkWomenHandbagCrossbody.Selector, mkWomenHandbagCrossbody.Handle)

	err := c.Visit(config.Get().MichaelKorsWomenHandbagsCrossbodyUrl)
	if err != nil {
		log.Println("error happened:", err.Error())
		return
	}
}

func get() {
	//debugger := colly.Debugger(&debug.LogDebugger{})
	c := colly.NewCollector()

	c.OnHTML(".product-detail__name h1.product-name", func(e *colly.HTMLElement) {
		log.Println("Name: ", e.Text)
	})

	c.OnHTML(".product-detail__prices span.value", func(e *colly.HTMLElement) {
		log.Println("price: ", strings.Trim(e.Text, " "))
	})

	c.OnHTML("#aboutDetails .details-dropdown--accordion-body.row span", func(e *colly.HTMLElement) {
		log.Println("about: ", e.Text)
	})

	err := c.Visit("https://usa.tommy.com/en/men/clothing/tops/slim-fit-tipped-polo/78J9447-VLP.html?journey=Tier_18482838")
	if err != nil {
		log.Println("error happened:", err)
		return
	}
}
