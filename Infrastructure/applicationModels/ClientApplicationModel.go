package applicationModels

type ClientApplicationModel struct {
	id            int
	firstName     string
	lastName      string
	birthDate     string
	passport      PassportApplicationModel
	jobs          []JobApplicationModel
	sex           string
	maritalStatus string
	children      int
}

func (u *ClientApplicationModel) GetId() int {
	return u.id
}

func (u *ClientApplicationModel) GetFirstName() string {
	return u.firstName
}

func (u *ClientApplicationModel) GetLastName() string {
	return u.lastName
}

func (u *ClientApplicationModel) GetBirthDate() string {
	return u.birthDate
}

func (u *ClientApplicationModel) GetPassport() PassportApplicationModel {
	return u.passport
}

func (u *ClientApplicationModel) GetJobs() []JobApplicationModel {
	return u.jobs
}

func (u *ClientApplicationModel) GetSex() string {
	return u.sex
}

func (u *ClientApplicationModel) GetMaritalStatus() string {
	return u.maritalStatus
}

func (u *ClientApplicationModel) GetChildren() int {
	return u.children
}
