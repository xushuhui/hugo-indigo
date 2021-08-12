package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"

	"log"
	"os"
)

func main() {
	//file, err := os.OpenFile("1.html", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	f, err := os.Open("./geektime/1.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	h1, err := doc.Find("#app").Find("h1").Html()
	_ = strings.TrimSpace(h1)
	body := doc.Find("._2c4hPkl9").Text()
	fmt.Println(body)
}
