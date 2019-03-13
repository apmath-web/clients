package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type BasicClient struct {
	firstName     string
	lastName      string
	birthDate     string
	passport      Domain.PassportViewModelInterface
	jobs          []Domain.JobViewModelInterface
	sex           string
	maritalStatus string
	children      int
}

type ClientViewModel struct {
	BasicClient
	validation Domain.ValidationInterface
}

func (u *ClientViewModel) GetFirstName() string {
	return u.firstName
}

func (u *ClientViewModel) GetLastName() string {
	return u.lastName
}

func (u *ClientViewModel) GetBirthDate() string {
	return u.birthDate
}

func (u *ClientViewModel) GetPassport() Domain.PassportViewModelInterface {
	return u.passport
}

func (u *ClientViewModel) GetJobs() []Domain.JobViewModelInterface {
	return u.jobs
}

func (u *ClientViewModel) GetSex() string {
	return u.sex
}

func (u *ClientViewModel) GetMaritalStatus() string {
	return u.maritalStatus
}

func (u *ClientViewModel) GetChildren() int {
	return u.children
}

func (u *ClientViewModel) Validate() bool {
	return true
}

func (u *ClientViewModel) GetValidation() Domain.ValidationInterface {
	return u.validation
}

func (u *ClientViewModel) MarshalJSON() (b []byte, e error) {
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
