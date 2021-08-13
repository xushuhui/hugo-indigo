package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func makefile(fileName string,f *os.File)  {
	mdf := strings.TrimSuffix(fileName, ".html")


	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	h1 := doc.Find("#app").Find("h1").First().Text()
	h1 = strings.TrimSpace(h1)
	h1 = `<h1>`+h1+`</h1>`
	//   _29HP61GA_0
	body,_ := doc.Find("._2c4hPkl9").Html()
	//_2QuafIpq_0
	file, err := os.OpenFile("./md/go/"+mdf+".md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	godown.Convert(file, strings.NewReader(h1+body), &godown.Option{

	})
	f.Close()
}
func main() {
	dir :="./go"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {

		f, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}
		makefile(file.Name(),f)

	}


}
func rename()  {
	//newName := strings.ReplaceAll(file.Name(), "html", "")
	//newName := file.Name()+".html"
	//os.Rename(filepath.Join(dir, file.Name()), filepath.Join(dir, newName))
}