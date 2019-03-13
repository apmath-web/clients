package viewModels

import "encoding/json"

type JobViewModel struct {
	name string
	wage int
}

func (j *JobViewModel) GetName() string {
	return j.name
}

func (j *JobViewModel) GetWage() int {
	return j.wage
}

func (j *JobViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"name": j.name,
		"wage": j.wage,
	})
}
