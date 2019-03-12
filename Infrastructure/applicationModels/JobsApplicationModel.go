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

func (j *JobsApplicationModel) SetName(name string) {
	j.name = name
}
func (j *JobsApplicationModel) SetWage(wage int) {
	j.wage = wage
}
