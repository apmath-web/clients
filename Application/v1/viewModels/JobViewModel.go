package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/validation"
	"github.com/apmath-web/clients/Domain"
)

type JobViewModel struct {
	Name       string `json:"name"`
	Wage       int    `json:"wage"`
	validation validation.Validation
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
	if j.Name == "" {
		j.validation.AddMessage(validation.GenMessage("jobName", "Is empty"))
	}
}

func (j *JobViewModel) validateWage() {
	if j.Wage <= 0 {
		j.validation.AddMessage(validation.GenMessage("jobWage", "Minus or zero value of wage"))
	}
}

func (j *JobViewModel) Validate() bool {
	j.validateName()
	j.validateWage()
	return j.validation.Empty()
}

func (j *JobViewModel) GetValidation() Domain.ValidationInterface {
	return &j.validation
}

func (j *JobViewModel) Hydrate(job Domain.JobModelDomainInterface) {
	j.Wage = job.GetWage()
	j.Name = job.GetName()
}
