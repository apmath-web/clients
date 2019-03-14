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
			"series": 23423,
			"number": 4342345,
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
		"maritalStatus": "maried",
		"children":      4,
	}
	strTest, _ := json.Marshal(js)
	fmt.Println(string(strTest))
	m := new(ClientViewModel)
	err := json.Unmarshal(strTest, m)
	fmt.Println(err)
	fmt.Printf("%+v\n", m)
	strResult, _ := json.Marshal(m)
	fmt.Println(string(strResult))

}
