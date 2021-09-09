package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func getHeaders() (header map[string]string) {
	header = map[string]string{
		//"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":          `login_sid_t=e3498f7cca72a692d34ae777ebc63039; cross_origin_proto=SSL; _s_tentry=-; Apache=8988982910612.766.1609150868347; SINAGLOBAL=8988982910612.766.1609150868347; ULV=1609150868353:1:1:1:8988982910612.766.1609150868347:; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWOg3rc-lellg-YikzmUDqZ5JpX5KMhUgL.FozN1hqfehMpS0z2dJLoIpjLxK-LB.-L1K5LxKqL1-zL1K.LxKnLB.-L1h.t; ALF=1662433647; SSOLoginState=1630897648; SCF=Al6SMtYhRRaWwDB15z5E5DhyjbPLGqwZ52MHaYUy0NiGY_RM0_MGhaQ3kWQ-UxbUK5jiZh8nfglHPLPZcGFFahg.; SUB=_2A25MMfWhDeRhGeRJ41QU8CnNzD6IHXVvR2BprDV8PUNbmtB-LWfEkW9NUneWxXSugxFDViEsC9aVsc7ukE55QHgD; UOR=,,www.google.com`,
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
		"Referer":         "http://weibo.com/minipublish",
		"Accept":          "*/*",
		"Content-Length":  "79664",
		"Accept-Encoding": "gzip, deflate, br",
		"Host":            "picupload.weibo.com",
	}
	return
}
func createReqBody(filePath string) (string, io.Reader, error) {
	var err error

	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf) // body writer

	f, err := os.Open(filePath)
	if err != nil {
		log.Println("err", err)
	}
	defer f.Close()

	// file part1
	_, fileName := filepath.Split(filePath)
	fw1, err := bw.CreateFormFile("pic1", fileName)
	if err != nil {
		log.Println("err", err)
	}
	_, err = io.Copy(fw1, f)
	if err != nil {
		log.Println("err", err)
	}
	bw.Close() //write the tail boundry
	return base64.StdEncoding.EncodeToString(buf.Bytes()), buf, nil
}

func Post(url string, jsonStr io.Reader) (res *http.Response, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, jsonStr)
	if err != nil {
		log.Println("err", err)
	}
	header := getHeaders()

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err = client.Do(req)
	if err != nil {
		log.Println("err", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("http err ", res.Status)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	return
}
func uploadSinaImg() {

	//_, body, err := createReqBody("./1.jpg")

	//now := utils.Int642String(time.Now().Unix())
	//https://photo.weibo.com/upload/photo
	urlStr := "https://picupload.weibo.com/interface/pic_upload.php?s=xml&ori=1&data=1&rotate=0&wm=&app=miniblog&mime=image%2Fjpeg"

	//&logo=&nick=0&marks=1&app=miniblog`
	//urlStr = url.QueryEscape(urlStr)
	body, err := os.OpenFile("./1.jpg", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Println("err", err)
	}
	Post(urlStr, body)

}
func main() {
	//getList()
	uploadSinaImg()
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

	res, err := http.Get("http://www.imooc.com/wiki/lambda/lambdaintro.html")
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
		i = i + 50
		title := strings.TrimSpace(s.Text())
		title = "Java从零开始（" + strconv.Itoa(i) + "）" + title + ".md"
		link, _ := s.Attr("href")
		link = "http://www.imooc.com" + link

		getContent(link, title)
		time.Sleep(10 * time.Second)
	})
}
