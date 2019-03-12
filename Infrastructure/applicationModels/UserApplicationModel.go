package applicationModels

import (
	"github.com/apmath-web/clients/Domain"
)

type UserApplicationModel struct {
	id            int
	firstName     string
	lastName      string
	birthDate     string
	passport      Domain.PassportApplicationModelInterface
	jobs          []Domain.JobsApplicationModelInterface
	sex           string
	maritalStatus string
	children      int
}

func (u *UserApplicationModel) GetId() int {
	return u.id
}

func (u *UserApplicationModel) GetFirstName() string {
	return u.firstName
}

func (u *UserApplicationModel) GetLastName() string {
	return u.lastName
}

func (u *UserApplicationModel) GetBirthDate() string {
	return u.birthDate
}

func (u *UserApplicationModel) GetPassport() Domain.PassportApplicationModelInterface {
	return u.passport
}

func (u *UserApplicationModel) GetJobs() []Domain.JobsApplicationModelInterface {
	return u.jobs
}

func (u *UserApplicationModel) GetSex() string {
	return u.sex
}

func (u *UserApplicationModel) GetMaritalStatus() string {
	return u.maritalStatus
}

func (u *UserApplicationModel) GetChildren() int {
	return u.children
}
