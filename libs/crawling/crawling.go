package crawling

import (
	"fmt"
	"net/http"

	"forYourStock/libs/common"

	"github.com/PuerkitoBio/goquery"
)

//Scrapper function
func Scrapper(code string) {
	// var baseURL string = "https://finance.naver.com/item/frgn.nhn?code=" + code
	var baseURL string = "https://taepseon.tistory.com/116"

	fmt.Println(baseURL)
	res, err := http.Get(baseURL)
	common.CheckErr(err)
	common.CheckCode(res)

	defer res.Body.Close()

	fmt.Println(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	common.CheckErr(err)

	doc.Find(".tah p11 nv01").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	fmt.Println("Done")
}
