package viewModels

import "encoding/json"

type JobsViewModel struct {
	name string
	wage int
}

func (j *JobsViewModel) GetName() string {
	return j.name
}

func (j *JobsViewModel) GetWage() int {
	return j.wage
}

func (j *JobsViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"name": j.name,
		"wage": j.wage,
	})
}
