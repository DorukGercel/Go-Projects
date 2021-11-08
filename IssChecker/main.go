package main

import (
	"fmt"
	"github.com/DorukGercel/Go-Projects/IssChecker/services/crawler"
	"time"
)

func main() {
	var s crawler.Satellite
	channel := make(chan []crawler.Data)
	for {
		go func(c chan []crawler.Data) {
			fmt.Println("Satellite request send...")
			crawler.SendGetRequest(s, c)
		}(channel)
		crawler.PrintInfo(<-channel)
		// Send new request after 5 seconds
		time.Sleep(5 * time.Second)
	}
}
