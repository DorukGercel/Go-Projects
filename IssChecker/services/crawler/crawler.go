package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data interface {
	MapToData(string) (Data, error)
	GetUrl() string
	PrintInfo()
}

func SendGetRequest(d Data) (Data, error) {
	resp, errReq := http.Get(d.GetUrl())
	if errReq != nil && resp.StatusCode == http.StatusOK {
		fmt.Println(errReq)
		return nil, errReq
	}
	bodyBytes, errIo := ioutil.ReadAll(resp.Body)
	if errIo != nil {
		fmt.Println(errIo)
		return nil, errIo
	}

	return d.MapToData(string(bodyBytes))
}
