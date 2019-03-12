package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type BasicUser struct {
	firstName     string
	lastName      string
	birthDate     string
	passport      Domain.PassportViewModelInterface
	jobs          []Domain.JobsViewModelInterface
	sex           string
	maritalStatus string
	children      int
}

type UserViewModel struct {
	BasicUser
	validation Domain.ValidationInterface
}

func (u *UserViewModel) GetFirstName() string {
	return u.firstName
}

func (u *UserViewModel) GetLastName() string {
	return u.lastName
}

func (u *UserViewModel) GetBirthDate() string {
	return u.birthDate
}

func (u *UserViewModel) GetPassport() Domain.PassportViewModelInterface {
	return u.passport
}

func (u *UserViewModel) GetJobs() []Domain.JobsViewModelInterface {
	return u.jobs
}

func (u *UserViewModel) GetSex() string {
	return u.sex
}

func (u *UserViewModel) GetMaritalStatus() string {
	return u.maritalStatus
}

func (u *UserViewModel) GetChildren() int {
	return u.children
}

func (u *UserViewModel) Validate() bool {
	return true
}

func (u *UserViewModel) GetValidation() Domain.ValidationInterface {
	return u.validation
}

func (u *UserViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"firstName":     u.firstName,
		"lastName":      u.lastName,
		"birthDate":     u.birthDate,
		"passport":      u.passport,
		"jobs":          u.jobs,
		"sex":           u.sex,
		"maritalStatus": u.maritalStatus,
		"children":      u.children,
	})
}
