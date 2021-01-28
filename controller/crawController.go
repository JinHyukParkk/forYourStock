package controller

import (
	"forYourStock/libs/crawling"
)

// 주식 정보 크롤링
func CrawStock() {
	crawling.Scrapper("005380")
}
