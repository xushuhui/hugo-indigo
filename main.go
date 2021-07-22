package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/godown"
)

func main() {
div := `<div class="content">
	<div class="content-item text-item">
	  <div class="cl-preview-section">
		<a id="anchor_0" class="virtual_anchor" target="_blank">
		</a>
		<h1 id="java-运算符">Java 运算符</h1>
	  </div>
	</div>
	<div class="content-item code-item">
	  <div class="code-box js-code-box">
		<div class="code-top">
		  <div class="left">
			<i class="imv2-code">
			</i>
			<span>实例演示</span>
		  </div>
		  <div class="right">
			<span class="preview">预览</span>
			<span class="copy">复制
			  <textarea type="text" class="code"></textarea>
			</span>
			<div class="copysuccess hide">复制成功！</div>
		  </div>
		</div>
		<div class="code-con">
		  <div class="cl-preview-section">
			<pre class=" language-java">
			  <code class="prism  language-java">
				<span class="token keyword">public</span>
				<span class="token keyword">class</span>
				<span class="token class-name">ArithmeticOperators1</span>
			   
			  </code>
			  <ul class="pre-numbering">
			  
				<li>9</li>
			  </ul>
			</pre>
		  </div>
		</div>
		<div class="code-bottom">
		  <a href="/wiki/run/434.html" target="_blank">运行案例</a>
		  <span>点击 "运行案例" 可查看在线运行效果</span>
		</div>
		<!---->
	  </div>
	</div>
	</div>`
	doc, err :=  goquery.NewDocumentFromReader(strings.NewReader(div))
	if err != nil {
		log.Fatal(err)
	}
	println(doc)
	getContent("", "")
}
func getContent(link string, title string) {
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
	content := doc.Find(".content-item")
	file, err := os.OpenFile("file.md", os.O_RDWR|os.O_CREATE, 0766) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	content.Each(func(i int, s *goquery.Selection) {
		h := s.Children().Children().Not(".code-bottom").Not(".code-top")
		
		fmt.Println(h.Text())
		html,_ := h.Html()
		fmt.Println(html)
		godown.Convert(file, strings.NewReader(html), nil)
		// var buf bytes.Buffer
		// h, _ := s.Html()
		// godown.Convert(&buf, strings.NewReader(h),nil)
		// fmt.Println(buf.String())

	})
	// fmt.Println(content.Text())

	
	

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
		title := s.Text()
		link, _ := s.Attr("href")
		link = "http://www.imooc.com" + link
		getContent(link, title)
	})
}
