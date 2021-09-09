package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func makefile(fileName string, f *os.File) {
	mdf := strings.TrimSuffix(fileName, ".html")

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	h1 := doc.Find("#app").Find("h1").First().Text()
	h1 = strings.TrimSpace(h1)
	h1 = `<h1>` + h1 + `</h1>`
	//   _29HP61GA_0
	body, _ := doc.Find("._2c4hPkl9").Html()
	//_2QuafIpq_0 _2c4hPkl9
	file, err := os.OpenFile("./md/fen/"+mdf+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	godown.Convert(file, strings.NewReader(h1+body), &godown.Option{})
	f.Close()
}
func main() {
	dir := "./html"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {

		f, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}

		makefile(file.Name(), f)
		f.Close()

		//rename(file.Name())
	}

}
func rename(fileName string) {
	mdf := strings.TrimSuffix(fileName, ".html")
	file, err := os.OpenFile("./md/kubernetes/"+mdf+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	pos := int64(0)

	for { //读取每一行内容
		line, err := reader.ReadString('\n')
		if err != nil { //读到末尾
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}

		if strings.Contains(line, "image/jpeg") {
			fmt.Println(line)
			//bytes := []byte("address " + ip + "\n")
			//file.WriteAt(bytes, pos)
		} //每一行读取完后记录位置
		pos += int64(len(line))
	}
}
