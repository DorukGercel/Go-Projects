package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data interface {
	MapToData(string) ([]Data, error)
	GetUrl() string
	Print()
}

func SendGetRequest(d Data, c chan []Data) {
	resp, errReq := http.Get(d.GetUrl())
	if errReq != nil && resp.StatusCode == http.StatusOK {
		fmt.Println(errReq)
		c <- make([]Data, 0)
		return
	}
	bodyBytes, errIo := ioutil.ReadAll(resp.Body)
	if errIo != nil {
		fmt.Println(errIo)
		c <- make([]Data, 0)
		return
	}

	val, errMap := d.MapToData(string(bodyBytes))
	if errMap != nil {
		fmt.Println(errMap)
		c <- make([]Data, 0)
		return
	}
	c <- val
}

func PrintInfo(d []Data) {
	for _, v := range d {
		v.Print()
	}
}
