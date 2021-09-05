package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
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
		"Content-Type": "application/json",
		"Cookie":       "_s_tentry=cn.bing.com; UOR=cn.bing.com,open.weibo.com,cn.bing.com; Apache=4992712349818.786.1630676593662; SINAGLOBAL=4992712349818.786.1630676593662; ULV=1630676593675:1:1:1:4992712349818.786.1630676593662:; login_sid_t=18995bbc55ec6c9277c8de66b563c071; cross_origin_proto=SSL; ALF=1662215082; SSOLoginState=1630679083; SUB=_2A25MNkB7DeRhGeRJ41QU8CnNzD6IHXVvQjazrDV8PUNbmtAKLUjfkW9NUneWxXuekD3KzoSL3Eg4y0QvYuB0xwMW; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWOg3rc-lellg-YikzmUDqZ5JpX5KzhUgL.FozN1hqfehMpS0z2dJLoIpjLxK-LB.-L1K5LxKqL1-zL1K.LxKnLB.-L1h.t; XSRF-TOKEN=SJO5Tak7wPdjVFasH4zVeSEQ; WBPSESS=GjA8I2NHhaKAyZg5agD9BG8KVHidYGC8nIa48kcJ-usTlfe1m9e4GGEPkdty4K1KjIecQ_qBe8LGtKvgiOKy_2K4Y5p1cG2xAIeFgJYBhb6DITtemACEpTBRFAKulVPt",
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
	}
	return
}
func createReqBody(filePath string) (string, io.Reader, error) {
	var err error

	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf) // body writer

	f, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	// file part1
	_, fileName := filepath.Split(filePath)
	fw1, _ := bw.CreateFormFile("pic1", fileName)
	io.Copy(fw1, f)

	bw.Close() //write the tail boundry
	return base64.StdEncoding.EncodeToString(buf.Bytes()), buf, nil
}
func Post(url string, jsonStr []byte) (res *http.Response, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

	//now := utils.Int642String(time.Now().Unix())
	urlStr := `https://picupload.weibo.com/interface/pic_upload.php?ori=1&mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog`
	//urlStr := `http://picupload.service.weibo.com/interface/pic_upload.php?mime=image%2Fjpeg&data=base64&url=0&markpos=1
	//&logo=&nick=0&marks=1&app=miniblog`
	//urlStr = url.QueryEscape(urlStr)
	base64, _, err := createReqBody("./2.png")
	if err != nil {
		log.Println("err", err)
	}
	jsons, _ := json.Marshal(map[string]string{
		"b64_data": base64,
	})
	Post(urlStr, jsons)

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
