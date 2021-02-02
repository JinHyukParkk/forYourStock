package crawling

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"forYourStock/libs/common"

	"github.com/PuerkitoBio/goquery"
)

//Scrapper function
func Scrapper(code string) {
	var baseURL string = "https://finance.naver.com/item/main.nhn?code=" + code

	fmt.Println(baseURL)
	res, err := http.Get(baseURL)
	common.CheckErr(err)
	common.CheckCode(res)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	strBody := string(body)

	fmt.Println(strBody)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	common.CheckErr(err)

	doc.Find(".invest_trend").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i)
		fmt.Println(s.Find("em").Text())
	})

	fmt.Println("Done")
}
