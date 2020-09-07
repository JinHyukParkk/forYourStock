package controller

import (
	"forYourStock/api"
	"io/ioutil"
)

func Start() {
	// API Key 세팅
	// key := os.Getenv("APIKEY")
	// 주식 종목 번호 파일 가져오기
	// api.GetCompanyList(key)
	// 종목 번호 파일 읽기
	// api.GetList()

	// 회사 리스트 불러오기
	// szCompanyList := getCompanyList()

	// var aCompanyList model.CompanyListStruct
	// if err := json.Unmarshal(szCompanyList, &aCompanyList); err != nil {
	// panic(err)
	// }

	// 회사 정보 가져오기
	// api.GetDisclosureInfo(aCompanyList.CompanyList[0], key)
	api.SendMail()

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
