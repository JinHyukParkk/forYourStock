package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"forYourStock/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetDisclosureInfo(data model.CompanyListStruct, key string) bytes.Buffer {
	// 결과값
	var szReturn bytes.Buffer

	// 회사 리스트
	aCompanyList := data.CompanyList

	// 시간 세팅
	now := time.Now()
	current := now.Format("20060102")
	age := now.AddDate(-1, 0, 0).Format("20060102")

	for _, oCompanyInfo := range aCompanyList {
		req, err := http.NewRequest("GET", "https://opendart.fss.or.kr/api/list.json", nil)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}

		q := req.URL.Query()
		q.Add("crtfc_key", key)
		q.Add("corp_code", oCompanyInfo.Code)
		q.Add("bgn_de", age)
		q.Add("end_de", current)

		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		// Client객체에서 Request 실행
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// 결과 출력
		response, _ := ioutil.ReadAll(resp.Body)

		var aCompanyDisclosure model.CompanyDisclosure
		if err := json.Unmarshal(response, &aCompanyDisclosure); err != nil {
			panic(err)
		}
		if aCompanyDisclosure.Status != "013" {
			var url bytes.Buffer
			url.WriteString("https://dart.fss.or.kr/dsaf001/main.do?rcpNo=")
			url.WriteString(aCompanyDisclosure.List[0].RceptNo)

			szReportUrl := url.String()
			szReturn.WriteString("############################# ")
			szReturn.WriteString(aCompanyDisclosure.List[0].CorpName)
			szReturn.WriteString(" #############################")
			szReturn.WriteString("\n")
			szReturn.WriteString("날짜 : ")
			szReturn.WriteString(aCompanyDisclosure.List[0].RceptDt)
			szReturn.WriteString("\n")
			szReturn.WriteString("공시명 : ")
			szReturn.WriteString(aCompanyDisclosure.List[0].ReportNm)
			szReturn.WriteString("\n")
			szReturn.WriteString("공시URL : ")
			szReturn.WriteString(szReportUrl)
			szReturn.WriteString("\n")
			szReturn.WriteString("##########################################################")
			szReturn.WriteString("\n")
			szReturn.WriteString("\n")
		}
	}

	return szReturn
}
