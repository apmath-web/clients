package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type JsonClient struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	BirthDate     string `json:"birthDate"`
	Sex           string `json:"sex"`
	MaritalStatus string `json:"maritalStatus"`
	Children      int    `json:"children"`
}

type ClientViewModel struct {
	JsonClient
	Passport   Domain.PassportViewModelInterface
	Jobs       []Domain.JobViewModelInterface
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

func (c *ClientViewModel) validateFirstName() {

}

func (c *ClientViewModel) validateLastName() {

}

func (c *ClientViewModel) validateBirthDate() {

}

func (c *ClientViewModel) validateSex() {

}

func (c *ClientViewModel) validateMaritalStatus() {

}

func (c *ClientViewModel) validateChildren() {

}

func (c *ClientViewModel) validatePassport() {

}

func (c *ClientViewModel) validateJobs() {

}

func (c *ClientViewModel) Validate() bool {
	c.validateFirstName()
	c.validateLastName()
	c.validateBirthDate()
	c.validateSex()
	c.validateMaritalStatus()
	c.validateChildren()
	c.validateJobs()
	c.validatePassport()
	return c.validation.Empty()
}

func (c *ClientViewModel) GetValidation() Domain.ValidationInterface {
	return c.validation
}

func (c *ClientViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"firstName":     c.FirstName,
		"lastName":      c.LastName,
		"birthDate":     c.BirthDate,
		"Passport":      c.Passport,
		"Jobs":          c.Jobs,
		"sex":           c.Sex,
		"maritalStatus": c.MaritalStatus,
		"children":      c.Children,
	})
}

func (c *ClientViewModel) UnmarshalJSON(b []byte) error {
	tmpClient := JsonClient{}
	err := json.Unmarshal(b, &tmpClient)
	if err := json.Unmarshal(b, &tmpClient); err != nil {
		return err
	}
	tmpPassport := struct{ Passport PassportViewModel }{}
	if err := json.Unmarshal(b, &tmpPassport); err != nil {
		return err
	}
	tmpJobs := struct{ Jobs []JobViewModel }{}
	if err := json.Unmarshal(b, &tmpJobs); err != nil {
		return err
	}
	c.JsonClient = tmpClient
	c.Passport = &tmpPassport.Passport
	for _, value := range tmpJobs.Jobs {
		c.Jobs = append(c.Jobs, &value)
	}
	return err
}
