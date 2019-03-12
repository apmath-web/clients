package viewModels

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
