package applicationModels

type JobApplicationModel struct {
	name string
	wage int
}

func (j *JobApplicationModel) GetName() string {
	return j.name
}

func (j *JobApplicationModel) GetWage() int {
	return j.wage
}
