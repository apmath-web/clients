package viewModels

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/clients/Domain"
)

type BasicClient struct {
	FirstName     string                            `json:"firstName"`
	LastName      string                            `json:"lastName"`
	BirthDate     string                            `json:"birthDate"`
	Passport      Domain.PassportViewModelInterface `json:"passport"`
	Jobs          []Domain.JobViewModelInterface    `json:"jobs"`
	Sex           string                            `json:"sex"`
	MaritalStatus string                            `json:"maritalStatus"`
	Children      int                               `json:"children"`
}

type ClientViewModel struct {
	BasicClient
	validation Domain.ValidationInterface
}

func (c *ClientViewModel) GetFirstName() string {
	return c.FirstName
}

func (c *ClientViewModel) GetLastName() string {
	return c.LastName
}

func (c *ClientViewModel) GetBirthDate() string {
	return c.BirthDate
}

func (c *ClientViewModel) GetPassport() Domain.PassportViewModelInterface {
	return c.Passport
}

func (c *ClientViewModel) GetJobs() []Domain.JobViewModelInterface {
	return c.Jobs
}

func (c *ClientViewModel) GetSex() string {
	return c.Sex
}

func (c *ClientViewModel) GetMaritalStatus() string {
	return c.MaritalStatus
}

func (c *ClientViewModel) GetChildren() int {
	return c.Children
}

func (c *ClientViewModel) Validate() bool {
	return true
}

func (c *ClientViewModel) GetValidation() Domain.ValidationInterface {
	return c.validation
}

func (c *ClientViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"firstName":     c.FirstName,
		"lastName":      c.LastName,
		"birthDate":     c.BirthDate,
		"passport":      c.Passport,
		"jobs":          c.Jobs,
		"sex":           c.Sex,
		"maritalStatus": c.MaritalStatus,
		"children":      c.Children,
	})
}
