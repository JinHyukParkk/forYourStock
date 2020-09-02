package controller

import (
	"io/ioutil"
	"forYourStock/model"
	"encoding/json"
	"fmt"
	// "forYourStock/api"
	
)

func Start() {
	// 주식 종목 번호 파일 가져오기
	// api.GiveCompanyList()
	// 종목 번호 파일 읽기
	// api.GiveList()
	szCompanyList := getCompanyList()

	var aCompanyList model.CompanyListStruct
	if err := json.Unmarshal(szCompanyList, &aCompanyList); err != nil {
		panic(err)
	}
	
}

func getCompanyList() []byte {
	dat, err := ioutil.ReadFile("data/companyList.json")
	check(err)
	return dat
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}