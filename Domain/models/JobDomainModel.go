package models

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
