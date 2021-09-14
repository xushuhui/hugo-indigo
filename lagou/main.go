package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"spider/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
)

func handleFile(fileName string, f *os.File) {
	mdf := strings.TrimSuffix(fileName, ".md")

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	h1 := `<h2>` + fileName + `</h2>`
	//   _29HP61GA_0
	html, _ := doc.Html()
	//_2QuafIpq_0 _2c4hPkl9
	file, err := os.OpenFile("./lagou/docker/md/"+mdf+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	godown.Convert(file, strings.NewReader(h1+html), &godown.Option{})
	f.Close()
	fmt.Println(mdf)
}
func do() {
	dir := "./lagou/docker"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}

		handleFile(file.Name(), f)

		//rename(file.Name())
	}
}

const LESSION_URL = `https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons?courseId=%d`
const IndexUrl = `https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail?lessonId=%d`

//
//https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail?lessonId=6863
const PAY_URL = `https://gate.lagou.com/v1/neirong/edu/member/drawCourse?courseId=%s`
const PURCHASE_URL = `https://gate.lagou.com/v1/neirong/kaiwu/getAllCoursePurchasedRecordForPC`

type lesson struct {
	id    int
	title string
}

func getLessons(courseId int) (lessons []lesson) {

	body := Get(fmt.Sprintf(LESSION_URL, courseId))
	var resp LessonsResp
	json.Unmarshal(body, &resp)

	for _, v := range resp.Content.CourseSectionList {
		for _, val := range v.CourseLessons {
			l := lesson{
				title: val.Theme,
				id:    val.ID,
			}
			lessons = append(lessons, l)
		}

	}
	return
}
func getLesson(lessonId int) (content string) {
	body := Get(fmt.Sprintf(IndexUrl, lessonId))
	var resp LessonResp
	json.Unmarshal(body, &resp)
	return resp.Content.TextContent
}

type Course struct {
	id    int
	title string
}

func getCourses() (courses []Course) {
	body := Get(fmt.Sprintf(PURCHASE_URL))
	var resp CourseResp
	json.Unmarshal(body, &resp)

	record := resp.Content.AllCoursePurchasedRecord
	for _, val := range record {
		for _, v := range val.CourseRecordList {
			c := Course{id: v.ID, title: v.Name}
			courses = append(courses, c)
		}
	}
	return
}
func main() {
	courses := getCourses()
	for _, c := range courses {
		lessons := getLessons(c.id)

		for _, v := range lessons {
			lesson := getLesson(v.id)
			dir := `./lagou/` + c.title + `/`
			if !utils.IsDir(dir) {
				os.Mkdir(dir, os.ModePerm)
			}
			title := strings.ReplaceAll(v.title, "|", "-")
			title = strings.ReplaceAll(title, "?", "")
			title = strings.ReplaceAll(title, "/", "-")
			title = strings.ReplaceAll(title, "\\", "-")
			title = strings.ReplaceAll(title, `"`, " ")
			title = strings.ReplaceAll(title, `“`, " ")
			title = strings.ReplaceAll(title, `”`, " ")
			makeMd(title, lesson, dir)

			fmt.Println(title)
		}

		fmt.Println(c.title)
	}

}
func makeMd(title string, content string, dir string) {
	h1 := `<h2>` + title + `</h2>`

	file, err := os.OpenFile(dir+title+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	godown.Convert(file, strings.NewReader(h1+content), &godown.Option{})
}
func getHeaders() (header map[string]string) {
	header = map[string]string{

		`Host`:             `gate.lagou.com`,
		`Connection`:       `keep-alive`,
		`Pragma`:           `no-cache`,
		`Cache-Control`:    `no-cache`,
		`sec-ch-ua`:        `"Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"`,
		`Accept`:           `application/json, text/plain, */*`,
		`Authorization`:    `X1RHVF9UR1QsImlwIjoiMTE2jVyBBDze5J3qeNqdwSHiit5XePRyXTa4bPV9lWi4cOXJi9d9NTGTYT5LOj0bZ2Q`,
		`X-L-REQ-HEADER`:   `iQtX0NBU19UR1RfIiwic3ViIjoi7oVvszBFSwEJCS9MYijVyBBDze5J3qeNqdwSHiit5XePRyXTa4bPV9lWi4cOXJi9d9NTGTYT5LOj0bZ2Q"}`,
		`sec-ch-ua-mobile`: `?0`,
		`User-Agent`:       `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36`,
		`Origin`:           `https://edu.lagou.com`,
		`Sec-Fetch-Site`:   `same-site`,
		`Sec-Fetch-Mode`:   `cors`,
		`Sec-Fetch-Dest`:   `empty`,
		`Referer`:          `https://kaiwu.lagou.com/`,
		`Edu-Referer`:      `https://kaiwu.lagou.com/course/courseInfo.htm?courseId=710&sid=20-h5Url-0&buyFrom=2&pageId=1pz4#/content`,
		`Accept-Encoding`:  `gzip, deflate, br`,
		`Accept-Language`:  `zh-CN,zh;q=0.9,en;q=0.8`,
		`Cookie`:           ``,
	}
	return
}
func Get(url string) (body []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
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
