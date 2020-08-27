package api

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
)

type myCloser interface {
	Close() error
}

// closeFile is a helper function which streamlines closing
// with error checking on different file types.
func closeFile(f myCloser) {
	err := f.Close()
	check(err)
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

// check is a helper function which streamlines error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
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