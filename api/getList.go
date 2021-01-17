package api

import (
	"archive/zip"
	"io/ioutil"

	// "encoding/xml"
	"io"
	"os"
	"path/filepath"
	// "forYourStock/model"
	"fmt"
)

func GetList() {
	fmt.Println("Exec Get")
	// XML 데이터 생성
	createXmlData()

	// xml unmarshal
	// 회사 리스트 불러오기
	// szCompanyList := getCompanyList()

	// var aCompanyList model.CompanyXmlResult
	// if err := xml.Unmarshal(szCompanyList, &aCompanyList); err != nil {
	// 	panic(err)
	// }

	// // 회사 정보 가져오기
	// fmt.Println(aCompanyList)

	// 파일 크기가 너무 큼 Memory 부족

}

func getCompanyList() []byte {
	dat, err := ioutil.ReadFile("data/CORPCODE.xml")
	check(err)
	return dat
}

// readAll is a wrapper function for ioutil.ReadAll. It accepts a zip.File as
// its parameter, opens it, reads its content and returns it as a byte slice.
func readAll(file *zip.File) []byte {
	fc, err := file.Open()
	check(err)
	defer closeFile(fc)

	content, err := ioutil.ReadAll(fc)
	check(err)

	return content
}

func createXmlData() {
	zipFile := "list.zip"

	// xml 파일로 떨구기  -- 파일이 너무 커서 안됨.
	err := unzip(zipFile, "data")
	check(err)
}

/*
뽑아보니깐 .. 상장하지 않는 회사들도 많네 ..ㅠ
*/

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
