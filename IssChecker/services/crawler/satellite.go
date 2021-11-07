package crawler

import (
	"encoding/json"
	"fmt"
)

const satelliteUrl = "https://api.wheretheiss.at/v1/satellites"

type Satellite struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (Satellite) MapToData(jsonText string) (Data, error) {
	var s Satellite
	var ss []Satellite
	err := json.Unmarshal([]byte(jsonText), &ss)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	s = ss[0]
	fmt.Println(s)
	return s, nil
}

func (Satellite) GetUrl() string {
	return satelliteUrl
}

func (s Satellite) PrintInfo() {
	fmt.Println("Satellite", s.Name)
}
