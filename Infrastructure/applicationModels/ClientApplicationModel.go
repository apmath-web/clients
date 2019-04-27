package applicationModels

import (
	"database/sql"
	"github.com/apmath-web/clients/Domain"
)

type JsonClient struct {
	Id            int    `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	BirthDate     string `db:"birth_date"`
	Sex           string `db:"sex"`
	MaritalStatus string `db:"marital_status"`
	Children      int    `db:"children"`
}

type ClientApplicationModel struct {
	JsonClient
	Passport PassportApplicationModel
	Jobs     []JobApplicationModel
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
	c.Passport.Hydrate(client.GetPassport())

}

func (c *ClientApplicationModel) UpdateId(id int) {
	c.Id = id
	c.Passport.ClientId = id
	tmpJobs := []JobApplicationModel{}
	for _, job := range c.Jobs {
		job.ClientId = id
		tmpJobs = append(tmpJobs, job)
	}
	c.Jobs = tmpJobs
}

func (c *ClientApplicationModel) SaveClient(tx *sql.Tx) (int, error) {
	var clientId int
	if err := tx.QueryRow("INSERT INTO clients "+
		"(first_name, last_name, birth_date, sex, marital_status, children) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6) RETURNING id",
		c.GetFirstName(), c.GetLastName(), c.GetBirthDate(),
		c.GetSex(), c.GetMaritalStatus(), c.GetChildren()).Scan(&clientId); err != nil {
		return 0, err
	}
	c.UpdateId(clientId)
	if err := c.Passport.SavePassport(tx); err != nil {
		return 0, err
	}
	tmpJobs := []JobApplicationModel{}
	for _, job := range c.Jobs {
		if err := job.SaveJob(tx); err != nil {
			return 0, err
		}
		tmpJobs = append(tmpJobs, job)
	}
	c.Jobs = tmpJobs
	return clientId, nil
}
