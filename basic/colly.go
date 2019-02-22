package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
)

func main() {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains("www.shopify.com"))
	//c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("http://10.164.177.170:8080", "https://10.164.177.170:8080")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Requesting " + request.URL.String() + " ...")
	})

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
		//log.Printf("%s\n", r.Body)
		fmt.Println("r.StatusCode", r.StatusCode)
	})

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	// Start the collector
	c.Visit("https://www.shopify.com/sitemap.xml")

	fmt.Println("len(knownUrls)", len(knownUrls))
	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println(url)
		fmt.Println()
	}
	fmt.Println("Collected", len(knownUrls), "URLs")
}
