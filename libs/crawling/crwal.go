package crawling

import (
	"fmt"
	"net/http"

	"forYourStock/libs/common"
)

//Scrapper function
func Scrapper(code string) {
	var baseURL string = "https://finance.naver.com/item/frgn.nhn?code=" + code

	res, err := http.Get(baseURL)
	common.CheckErr(err)
	common.CheckCode(res)

	defer res.Body.Close()

	fmt.Println(res)
	fmt.Println("Done")
}
