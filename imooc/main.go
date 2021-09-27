package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
)

func getHeaders() (header map[string]string) {
	header = map[string]string{
		//"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":          `login_sid_t=e3498f7cca72a692d34ae777ebc63039; cross_origin_proto=SSL; _s_tentry=-; Apache=8988982910612.766.1609150868347; SINAGLOBAL=8988982910612.766.1609150868347; ULV=1609150868353:1:1:1:8988982910612.766.1609150868347:; SSOLoginState=1630897648; UOR=,,www.google.com; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWOg3rc-lellg-YikzmUDqZ5JpX5KMhUgL.FozN1hqfehMpS0z2dJLoIpjLxK-LB.-L1K5LxKqL1-zL1K.LxKnLB.-L1h.t; ALF=1663029958; SCF=Al6SMtYhRRaWwDB15z5E5DhyjbPLGqwZ52MHaYUy0NiG3N7MO-tTkjKv4untB1fQ01qIeCkLkC8z_cofwgLW5iw.; SUB=_2A25MOu8YDeRhGeRJ41QU8CnNzD6IHXVvTkfQrDV8PUNbmtB-LRCmkW9NUneWxRXeO-zBwm4fp9-SWgRfAqIzAUUb; WBStorage=d335429e|undefined`,
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
		"Referer":         "http://weibo.com/minipublish",
		"Accept":          "*/*",

		"Accept-Encoding": "gzip, deflate, br",
		"Host":            "picupload.weibo.com",
		"Connection":      "keep-alive",
	}
	return
}
func Download(url string,dir string) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Referer", "http://www.imooc.com/")

	resp, err := (&http.Client{}).Do(req)

	s := strings.Split(url, "/")
	name := s[len(s)-1]

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(dir+name, data, 0644)
	//img := uploadSinaImg(name)
	return
}
func Post(url string, jsonStr io.Reader) (body []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, jsonStr)
	if err != nil {
		log.Println("err", err)
	}
	header := getHeaders()

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("http err ", res.Status)
		return
	}
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	return
}
func uploadSinaImg(fileName string) string {

	urlStr := "https://picupload.weibo.com/interface/pic_upload.php?s=json&ori=1&data=1&rotate=0&wm=&app=miniblog&mime=image%2Fjpeg"

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Println("err", err)
	}
	b, _ := Post(urlStr, file)

	var res Res
	err = json.Unmarshal(b, &res)
	if err != nil {
		log.Println(err)
		return ""
	}
	fmt.Println("filename ", fileName)
	if res.Code != "A00006" {
		fmt.Println("err", string(b))
		return ""
	}
	name := res.Data.Pics.Pic1.PID
	return "https://tvax1.sinaimg.cn/large/" + name

}
func replaceImg(s string) {
	var index int
	for k, v := range s {
		if v == 40 {
			index = k
		}
	}
	old := s[index+1 : len(s)-1]

	s = strings.ReplaceAll(s, old, "new")
}
func main() {
	//getList("http://www.imooc.com/wiki/lambda/lambdaintro.html", 50, "./imooc/java/")
	//

	//	output, needHandle, err := handleMdImg(dir+"/"+file.Name())
	//fmt.Println(file.Name())
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//	if needHandle {
	//		err = writeToFile(dir+"/"+file.Name(), output)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//
	//}

}

type Posts struct {
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
	IsPublic int    `json:"is_public"`
	Html     string `json:"html"`
}

func importPost(title, content string) {
	client := &http.Client{}
	s, _ := json.Marshal(Posts{
		Title:    title,
		Markdown: content,
	})

	var data = strings.NewReader(string(s))
	req, err := http.NewRequest("POST", "https://api.mdnice.com/articles", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJyb2xlIjoidXNlciIsInVzZXJJZCI6Ik9URXdOdz09Iiwic3ViIjoiNDc0NDk3MDk3QHFxLmNvbSIsImlzcyI6IjkwYjlhNjNjODFjYzYzNTg4NDg2IiwiaWF0IjoxNjMxNjEyOTc0LCJhdWQiOiJtZG5pY2UtYXBpIiwiZXhwIjoxNjM0MjA0OTc0LCJuYmYiOjE2MzE2MTI5NzR9.dejutJy1hcnArLB3GES_quaWVNp6lOo2z4UGlaIXrEM")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
func getContent(link string, title string) {
	fileName := title
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
		GuessLang: func(s string) (string, error) { return "javascript", nil },
	})

	output, needHandle, err := handleMdImg(fileName)
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

func handleMdImg(fileName string) ([]byte, bool, error) {
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
				str1 := []byte("### 微信公众号老徐说\n")
				str2 := []byte("\n![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)")
				output = append(output, str1...)
				output = append(output, str2...)

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
		} else if strings.Contains(string(line), "img.mukewang.com") {
			s := string(line)
			var index int
			for k, v := range s {
				if v == 40 {
					index = k
				}
			}
			old := s[index+1 : len(s)-1]
			if !strings.Contains(old, "http://") && !strings.Contains(old, "https://") {
				old = "https:"+old
			}
			Download(old,"./img/")
			newByte := []byte(s)
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
				str1 := []byte("### 微信公众号老徐说\n")
				str2 := []byte("\n![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)")
				output = append(output, str1...)
				output = append(output, str2...)

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

func getList(url string, start int, dir string) {

	res, err := http.Get(url)
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
		i = i + start
		title := strings.TrimSpace(s.Text())

		title = "Java从零开始（" + strconv.Itoa(i) + "）" + strings.ReplaceAll(title, "/", "") + ".md"
		link, _ := s.Attr("href")
		link = "http://www.imooc.com" + link

		getContent(link, dir+title)
		time.Sleep(1 * time.Second)
	})
}
