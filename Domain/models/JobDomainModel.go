package models

import "github.com/apmath-web/clients/Domain"

type JobDomainModel struct {
	Name string `json:"name"`
	Wage int    `json:"wage"`
}

func (j *JobDomainModel) GetName() string {
	return j.Name
}

func (j *JobDomainModel) GetWage() int {
	return j.Wage
}

func GenJobDomainModel(name string, wage int) Domain.JobDomainModelInterface {
	dm := new(JobDomainModel)
	dm.Wage = wage
	dm.Name = name
	return dm
}
