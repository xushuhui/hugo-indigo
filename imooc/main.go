package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
)

func main() {
	getList()

}
func getContent(link string, title string) {
	fileName := "imooc/" + title
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	res, err := http.Get(link)
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
	content := doc.Find(".content").Find(".code-bottom").Remove().End().Find(".code-top").Remove().End().
		Find(".select-options").Remove().End().Find(".note-edit").Remove().End().Find(".note-detail").Remove().End()

	html, _ := content.Html()
	godown.Convert(file, strings.NewReader(html), &godown.Option{
		GuessLang: func(s string) (string, error) { return "java", nil },
	})

	output, needHandle, err := handleMd(fileName)
	if err != nil {
		panic(err)
	}
	if needHandle {
		err = writeToFile(fileName, output)
		if err != nil {
			panic(err)
		}
	}
}
func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
func handleMd(fileName string) ([]byte, bool, error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				str1 := []byte("### 微信公众号\n")
				str2 := []byte("\n![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)")
				output = append(output, str1...)
				output = append(output, str2...)
				fmt.Println("end")
				return output, needHandle, nil
			}
			return nil, needHandle, err
		}
		if string(line) == "<!---->" {
			newByte := []byte("")
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
	return output, needHandle, nil

}
func replace(r io.Reader, w io.Writer) error {
	// use scanner to read line by line
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if line == "<!---->" {
			line = ""
		}
		if _, err := io.WriteString(w, line+"\n"); err != nil {
			return err
		}
	}
	return sc.Err()
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
		title := strings.TrimSpace(s.Text())
		title = "Java从零开始（" + strconv.Itoa(i) + "）" + title + ".md"
		link, _ := s.Attr("href")
		link = "http://www.imooc.com" + link

		getContent(link, title)
		time.Sleep(10 * time.Second)
	})
}
