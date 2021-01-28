package common

import (
	"log"
	"net/http"
)

func CheckCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("Status code err: %d %s", res.StatusCode, res.Status)
	}
}
