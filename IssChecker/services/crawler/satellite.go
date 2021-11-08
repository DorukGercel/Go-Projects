package crawler

import (
	"fmt"
	"github.com/DorukGercel/Go-Projects/IssChecker/services/json_handler"
)

const satelliteUrl = "https://api.wheretheiss.at/v1/satellites"

type Satellite struct {
	name string
	id   float64
}

func (Satellite) MapToData(jsonText string) ([]Data, error) {
	var jsonHandler json_handler.JsonHandler
	var d []Data

	err := jsonHandler.Unmarshall(jsonText)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, v := range jsonHandler {
		s := Satellite{
			name: v["name"].(string),
			id:   v["id"].(float64),
		}
		d = append(d, s)
	}
	return d, nil
}

func (Satellite) GetUrl() string {
	return satelliteUrl
}

func (s Satellite) Print() {
	fmt.Println("Satellite: ", "Name: ", s.name, ", ID: ", s.id)
}
