package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/validation"
	"github.com/apmath-web/clients/Domain"
	"time"
)

type JsonClient struct {
	FirstName     string            `json:"firstName"`
	LastName      string            `json:"lastName"`
	BirthDate     string            `json:"birthDate"`
	Sex           string            `json:"sex"`
	MaritalStatus string            `json:"maritalStatus"`
	Children      int               `json:"children"`
	Passport      PassportViewModel `json:"passport"`
	Jobs          []JobViewModel    `json:"jobs"`
}

type ClientViewModel struct {
	JsonClient
	validation validation.Validation
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

func (c *ClientViewModel) GetPassport() PassportViewModel {
	return c.Passport
}

func (c *ClientViewModel) GetJobs() []JobViewModel {
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
	if c.FirstName == "" {
		c.validation.AddMessage(validation.GenMessage("firstName", "Is empty"))
	}
}

func (c *ClientViewModel) validateLastName() {
	if c.LastName == "" {
		c.validation.AddMessage(validation.GenMessage("lastName", "Is empty"))
	}
}

func (c *ClientViewModel) validateBirthDate() {
	if c.BirthDate == "" {
		c.validation.AddMessage(validation.GenMessage("BirthDate", "Is empty"))
	}
	if _, err := time.Parse("2006-01-02", c.BirthDate); err != nil {
		c.validation.AddMessage(validation.GenMessage("birthDate", "Incorrect date format"))
	}
}

func (c *ClientViewModel) validateSex() {
	if c.Sex != "male" && c.Sex != "female" {
		c.validation.AddMessage(validation.GenMessage("sex", "Unknown value"))
	}
}

func (c *ClientViewModel) validateMaritalStatus() {
	if c.MaritalStatus != "single" && c.MaritalStatus != "married" {
		c.validation.AddMessage(validation.GenMessage("maritalStatus", "Unknown value"))
	}
}

func (c *ClientViewModel) validateChildren() {
	if c.Children < 0 {
		c.validation.AddMessage(validation.GenMessage("children", "Minus value"))
	}
	if c.Children > 0 && c.MaritalStatus != "married" {
		c.validation.AddMessage(validation.GenMessage("children",
			"You can't have any children if you aren't married"))
	}
}

func (c *ClientViewModel) validatePassport() {
	if !c.Passport.Validate() {
		for _, msg := range c.Passport.GetValidation().GetMessages() {
			c.validation.AddMessage(msg)
		}
	}
}

func (c *ClientViewModel) validateJobs() {
	for _, job := range c.Jobs {
		if !job.Validate() {
			for _, msg := range job.GetValidation().GetMessages() {
				c.validation.AddMessage(msg)
			}
		}
	}
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
	return &c.validation
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

func (c *ClientViewModel) UnmarshalJSON(b []byte) error {
	tmpClient := JsonClient{}
	err := json.Unmarshal(b, &tmpClient)
	if err := json.Unmarshal(b, &tmpClient); err != nil {
		return err
	}
	c.JsonClient = tmpClient
	return err
}

func (c *ClientViewModel) Hydrate(client Domain.ClientDomainModelInterface) {
	c.FirstName = client.GetFirstName()
	c.LastName = client.GetLastName()
	c.BirthDate = client.GetBirthDate()
	c.Sex = client.GetSex()
	c.MaritalStatus = client.GetMaritalStatus()
	c.Children = client.GetChildren()
	for _, job := range client.GetJobs() {
		tmpJob := JobViewModel{}
		tmpJob.Hydrate(job)
		c.Jobs = append(c.Jobs, tmpJob)
	}

}
