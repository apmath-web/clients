package viewModels

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreditViewModel(t *testing.T) {
	js := map[string]interface{}{
		"firstName": "Lev",
		"lastName":  "Kovalenko",
		"birthDate": "2019-01-02",
		"Passport": map[string]int{
			"series": 2342,
			"number": 434345,
		},
		"Jobs": []map[string]interface{}{
			{
				"name": "test1",
				"wage": 536446,
			},
			{
				"name": "test4",
				"wage": 53644346,
			},
		},
		"sex":           "male",
		"maritalStatus": "married",
		"children":      4,
	}
	strTest, _ := json.Marshal(js)
	fmt.Println(string(strTest))
	m := new(ClientViewModel)
	_ = json.Unmarshal(strTest, m)
	fmt.Printf("%+v\n", m)
	if !m.Validate() {
		msg, _ := json.Marshal(m.GetValidation())
		fmt.Println(string(msg))
	}
	strResult, _ := json.Marshal(m)
	fmt.Println(string(strResult))

}
