package api

import (
	"os"
	"bytes"
	"github.com/GiterLab/urllib"
	"bufio"
	"fmt"
)

func GetCompanyList(key string) {
	var b bytes.Buffer
	b.WriteString("https://opendart.fss.or.kr/api/corpCode.xml?crtfc_key=")
	b.WriteString(key)

	url := b.String()

	req := urllib.Get(url)
	req.Debug(true)
	resp, err := req.String()
	check(err)

	listZip, err := os.Create("list.zip")
	check(err)
	w := bufio.NewWriter(listZip)
	n, err := w.WriteString(resp)
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}