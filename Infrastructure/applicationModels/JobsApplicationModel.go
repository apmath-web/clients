package applicationModels

type JobsApplicationModel struct {
	name string
	wage int
}

func (j *JobsApplicationModel) GetName() string {
	return j.name
}

func (j *JobsApplicationModel) GetWage() int {
	return j.wage
}
