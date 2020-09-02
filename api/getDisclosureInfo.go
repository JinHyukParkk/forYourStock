package api

import (
	"forYourStock/model"
	"net/http"
	"log"
	"os"
	"fmt"
)

func GetDisclosureInfo(companyInfo model.CompanyStruct, key string) {
	req, err := http.NewRequest("GET", "https://opendart.fss.or.kr/api/list.json", nil)
    if err != nil {
        log.Print(err)
        os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("crtfc_key", key)
	q.Add("corp_code", companyInfo.Code)

	req.URL.RawQuery = q.Encode()

    fmt.Println(req.URL.String())	
}