package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
)

func main() {

	getContent("", "")
}
func getContent(link string, title string) {
	res, err := http.Get("http://www.imooc.com/wiki/javalesson/operators.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	content := doc.Find(".content-item")
	file, err := os.OpenFile("file.md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	content.Each(func(i int, s *goquery.Selection) {
		h := s.Children().Children().Not(".code-bottom").Not(".code-top")
		
		fmt.Println(h.Text())
		html,_ := h.Html()
		fmt.Println(html)
		godown.Convert(file, strings.NewReader(html), nil)
		// var buf bytes.Buffer
		// h, _ := s.Html()
		// godown.Convert(&buf, strings.NewReader(h),nil)
		// fmt.Println(buf.String())

	})
	// fmt.Println(content.Text())

	
	

}
func getList() {
	res, err := http.Get("http://www.imooc.com/wiki/javalesson/operators.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".jie").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Text()
		link, _ := s.Attr("href")
		link = "http://www.imooc.com" + link
		getContent(link, title)
	})
}
