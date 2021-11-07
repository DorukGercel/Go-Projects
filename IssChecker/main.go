package main

import (
	"github.com/DorukGercel/Go-Projects/IssChecker/services/crawler"
)

func main() {
	var s crawler.Satellite
	d, err := crawler.SendGetRequest(s)
	if err != nil {
		return
	}
	d.PrintInfo()
}
