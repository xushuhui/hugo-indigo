package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
	//md "github.com/JohannesKaufmann/html-to-markdown"
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
	body,_ := doc.Find("._2c4hPkl9").Html()
//_2QuafIpq_0
	file, err := os.OpenFile("2.md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	godown.Convert(file, strings.NewReader(body), &godown.Option{
		GuessLang: func(s string) (string, error) { return "java", nil },
	})
	//converter := md.NewConverter("", true, nil)
	//markdown := converter.Convert(body)
	//file.WriteString(markdown)
}
