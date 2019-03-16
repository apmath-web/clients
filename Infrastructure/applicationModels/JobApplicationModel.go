package applicationModels

import "github.com/apmath-web/clients/Domain"

type JobApplicationModel struct {
	Name string
	Wage int
}

func (j *JobApplicationModel) GetName() string {
	return j.Name
}

func (j *JobApplicationModel) GetWage() int {
	return j.Wage
}

func (j *JobApplicationModel) Hydrate(job Domain.JobModelDomainInterface) {
	j.Wage = job.GetWage()
	j.Name = job.GetName()
}
