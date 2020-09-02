package api

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
)

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

func GiveList() {
	zipFile := "list.zip"

	zf, err := zip.OpenReader(zipFile)
	check(err)
	defer closeFile(zf)

	for _, file := range zf.File {
		// fmt.Printf("=%s\n", file.Name)
		fmt.Printf("%s\n\n", readAll(file)) // file content
	}
}
/*
뽑아보니깐 .. 상장하지 않는 회사들도 많네 ..ㅠ 
*/