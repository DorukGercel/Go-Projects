package json_handler

import (
	"encoding/json"
	"fmt"
)

type JsonHandler []map[string]interface{}

func (j *JsonHandler) Unmarshall(jsonText string) error {
	err := json.Unmarshal([]byte(jsonText), &j)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
