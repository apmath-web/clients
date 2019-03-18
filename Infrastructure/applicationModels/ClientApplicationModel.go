package applicationModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type JsonClient struct {
	Id            int                      `json:"id"`
	FirstName     string                   `json:"firstName"`
	LastName      string                   `json:"lastName"`
	BirthDate     string                   `json:"birthDate"`
	Sex           string                   `json:"sex"`
	MaritalStatus string                   `json:"maritalStatus"`
	Children      int                      `json:"children"`
	Passport      PassportApplicationModel `json:"passport"`
	Jobs          []JobApplicationModel    `json:"jobs"`
}

type ClientApplicationModel struct {
	JsonClient
}

func (c *ClientApplicationModel) GetId() int {
	return c.Id
}

func (c *ClientApplicationModel) GetFirstName() string {
	return c.FirstName
}

func (c *ClientApplicationModel) GetLastName() string {
	return c.LastName
}

func (c *ClientApplicationModel) GetBirthDate() string {
	return c.BirthDate
}

func (c *ClientApplicationModel) GetPassport() PassportApplicationModel {
	return c.Passport
}

func (c *ClientApplicationModel) GetJobs() []JobApplicationModel {
	return c.Jobs
}

func (c *ClientApplicationModel) GetSex() string {
	return c.Sex
}

func (c *ClientApplicationModel) GetMaritalStatus() string {
	return c.MaritalStatus
}

func (c *ClientApplicationModel) GetChildren() int {
	return c.Children
}

func (c *ClientApplicationModel) Hydrate(client Domain.ClientDomainModelInterface) {
	c.FirstName = client.GetFirstName()
	c.LastName = client.GetLastName()
	c.BirthDate = client.GetBirthDate()
	c.Sex = client.GetSex()
	c.MaritalStatus = client.GetMaritalStatus()
	c.Children = client.GetChildren()
	for _, job := range client.GetJobs() {
		tmpJob := JobApplicationModel{}
		tmpJob.Hydrate(job)
		c.Jobs = append(c.Jobs, tmpJob)
	}

}

func (c *ClientApplicationModel) UnmarshalJSON(b []byte) error {
	tmpClient := JsonClient{}
	err := json.Unmarshal(b, &tmpClient)
	if err := json.Unmarshal(b, &tmpClient); err != nil {
		return err
	}
	c.JsonClient = tmpClient
	return err
}

func (u *ClientApplicationModel) SetId(id int) {
	if u.Id == -1 {
		u.Id = id
	} else {
		panic("Your set one more ID!")
	}
}
