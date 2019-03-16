package applicationModels

import "github.com/apmath-web/clients/Domain"

type ClientApplicationModel struct {
	Id            int
	FirstName     string
	LastName      string
	BirthDate     string
	Passport      PassportApplicationModel
	Jobs          []JobApplicationModel
	Sex           string
	MaritalStatus string
	Children      int
}

func (u *ClientApplicationModel) GetId() int {
	return u.Id
}

func (u *ClientApplicationModel) GetFirstName() string {
	return u.FirstName
}

func (u *ClientApplicationModel) GetLastName() string {
	return u.LastName
}

func (u *ClientApplicationModel) GetBirthDate() string {
	return u.BirthDate
}

func (u *ClientApplicationModel) GetPassport() PassportApplicationModel {
	return u.Passport
}

func (u *ClientApplicationModel) GetJobs() []JobApplicationModel {
	return u.Jobs
}

func (u *ClientApplicationModel) GetSex() string {
	return u.Sex
}

func (u *ClientApplicationModel) GetMaritalStatus() string {
	return u.MaritalStatus
}

func (u *ClientApplicationModel) GetChildren() int {
	return u.Children
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
