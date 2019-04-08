package main

// package main

// import (
// 	"fmt"

// 	"github.com/PuerkitoBio/goquery"
// )

// func GetPage(url string) {
// 	doc, _ := goquery.NewDocument(url)
// 	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
// 		url, _ := s.Attr("src")
// 		fmt.Println(url)
// 	})
// }

// func main() {
// 	url := "http://blog.golang.org/"
// 	GetPage(url)
// }
//----------------------------------------------------------
//golangサイトから画像の抽出処理
//----------------------------------------------------------
// package main

// import (
// 	"flag"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"sync"

// 	"github.com/PuerkitoBio/goquery"
// )

// var stock = []string{}
// var base = "http://blog.golang.org/"
// var i int = 0
// var wg = new(sync.WaitGroup)

// func main() {
// 	flag.Parse()
// 	fmt.Println("It works!")
// 	doc, _ := goquery.NewDocument(base)
// 	results := makeUrl(doc)
// 	for len(results) > 0 {
// 		results = GetUrl(results)
// 	}
// 	wg.Wait()
// }

// func containsInStock(value string) bool {
// 	l := len(stock)
// 	for i := 0; i < l; i++ {
// 		if stock[i] == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetUrl(urls []*url.URL) []*url.URL {
// 	saurceUrl := []*url.URL{}
// L:
// 	for _, item := range urls {
// 		url_string := item.String()
// 		if !strings.Contains(url_string, base) {
// 			continue L
// 		}
// 		if containsInStock(url_string) {
// 			continue L
// 		}
// 		fmt.Println(url_string)
// 		stock = append(stock, url_string)
// 		doc, _ := goquery.NewDocument(url_string)
// 		results := makeUrl(doc)
// 		wg.Add(1)
// 		go GetImage(doc)
// 		saurceUrl = append(saurceUrl, results...)
// 	}
// 	return saurceUrl
// }

// func makeUrl(doc *goquery.Document) []*url.URL {
// 	var result []*url.URL
// 	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
// 		target, _ := s.Attr("href")
// 		base, _ := url.Parse(base)
// 		targets, _ := url.Parse(target)
// 		result = append(result, base.ResolveReference(targets))
// 	})
// 	return result
// }

// func GetImage(doc *goquery.Document) {
// 	var result []*url.URL
// 	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
// 		target, _ := s.Attr("src")
// 		base, _ := url.Parse(base)
// 		targets, _ := url.Parse(target)
// 		result = append(result, base.ResolveReference(targets))
// 	})
// 	for _, imageUrl := range result {
// 		imageUrl_String := imageUrl.String()
// 		if containsInStock(imageUrl_String) {
// 			continue
// 		}
// 		response, err := http.Get(imageUrl_String)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer response.Body.Close()
// 		file, err := os.Create(fmt.Sprintf("hoge%d.jpg", i))
// 		i++
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer file.Close()
// 		io.Copy(file, response.Body)
// 	}
// 	wg.Done()
// }
//----------------------------------------------------------

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// url := "https://www.nikkei.com/"
	url := "http://anime.eiga.com/program/"

	doc, _ := goquery.NewDocument(url)
	fmt.Println("--- <title>タグ ----------")
	fmt.Println(doc.Find("title").Text())
	fmt.Println("")

	fmt.Println("--- イベント新着情報 ----------")
	doc.Find("ul.eventInfoList > li").Each(func(i int, s *goquery.Selection) {
		//rank := s.Find("a > span:nth-child(1)").Text()
		// title := s.Find("a > span.m-sub_access_ranking_title").Text()
		eventDate := s.Find("div.eventInfoDate").Text()
		eventTitle := s.Find("div.eventInfoTtl > a").Text()

		fmt.Println(fixTitle(eventDate) + "-" + fixTitle(eventTitle))

	})
}

func fixTitle(title string) string {
	return strings.TrimSpace(title)
}
