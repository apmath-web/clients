package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type JobViewModel struct {
	Name       string `json:"name"`
	Wage       int    `json:"wage"`
	validation Domain.ValidationInterface
}

func (j *JobViewModel) GetName() string {
	return j.Name
}

func (j *JobViewModel) GetWage() int {
	return j.Wage
}

func (j *JobViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"name": j.Name,
		"wage": j.Wage,
	})
}

func (j *JobViewModel) validateName() {

}

func (j *JobViewModel) validateWage() {

}

func (j *JobViewModel) Validate() bool {
	j.validateName()
	j.validateWage()
	return j.validation.Empty()
}

func (j *JobViewModel) GetValidation() Domain.ValidationInterface {
	return j.validation
}
