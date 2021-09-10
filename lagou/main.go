package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func handleFile(fileName string,f *os.File)  {
	mdf := strings.TrimSuffix(fileName, ".md")


	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	h1 := `<h2>`+fileName+`</h2>`
	//   _29HP61GA_0
	html,_ := doc.Html()
	//_2QuafIpq_0 _2c4hPkl9
	file, err := os.OpenFile("./lagou/docker/md/"+mdf+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	godown.Convert(file, strings.NewReader(h1+html), &godown.Option{

	})
	f.Close()
	fmt.Println(mdf)
}
func main() {
	dir :="./lagou/docker"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir(){
			continue
		}
		f, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}

		handleFile(file.Name(),f)


		//rename(file.Name())
	}
}