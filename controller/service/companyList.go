package service

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"forYourStock/libs/common"

	"github.com/GiterLab/urllib"
)

func GetCompanyList(key string) {
	var b bytes.Buffer
	b.WriteString("https://opendart.fss.or.kr/api/corpCode.xml?crtfc_key=")
	b.WriteString(key)

	url := b.String()

	req := urllib.Get(url)
	req.Debug(true)
	resp, err := req.String()
	common.Check(err)

	listZip, err := os.Create("list.zip")
	common.Check(err)

	w := bufio.NewWriter(listZip)
	n, err := w.WriteString(resp)
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}
