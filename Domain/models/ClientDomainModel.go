package models

import "github.com/apmath-web/clients/Domain"

type ClientDomainModel struct {
	FirstName     string
	LastName      string
	BirthDate     string
	Sex           string
	MaritalStatus string
	Children      int
	Passport      Domain.PassportDomainModelInterface
	Jobs          []Domain.JobDomainModelInterface
}

func (c *ClientDomainModel) GetFirstName() string {
	return c.FirstName
}

func (c *ClientDomainModel) GetLastName() string {
	return c.LastName
}

func (c *ClientDomainModel) GetBirthDate() string {
	return c.BirthDate
}

func (c *ClientDomainModel) GetPassport() Domain.PassportDomainModelInterface {
	return c.Passport
}

func (c *ClientDomainModel) GetJobs() []Domain.JobDomainModelInterface {
	return c.Jobs
}

func (c *ClientDomainModel) GetSex() string {
	return c.Sex
}

func (c *ClientDomainModel) GetMaritalStatus() string {
	return c.MaritalStatus
}

func (c *ClientDomainModel) GetChildren() int {
	return c.Children
}

func GenClientDomainModel(firstName string, lastName string, birthDate string, sex string,
	maritalStatus string, children int, passport Domain.PassportDomainModelInterface,
	jobs []Domain.JobDomainModelInterface) Domain.ClientDomainModelInterface {
	dm := new(ClientDomainModel)
	dm.FirstName = firstName
	dm.LastName = lastName
	dm.BirthDate = birthDate
	dm.Sex = sex
	dm.MaritalStatus = maritalStatus
	dm.Children = children
	dm.Passport = passport
	dm.Jobs = jobs
	return dm
}
