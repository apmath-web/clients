package applicationModels

import (
	"database/sql"
	"errors"
	"github.com/apmath-web/clients/Domain"
	"github.com/jmoiron/sqlx"
)

type DbClient struct {
	Id            int    `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	BirthDate     string `db:"birth_date"`
	Sex           string `db:"sex"`
	MaritalStatus string `db:"marital_status"`
	Children      int    `db:"children"`
}

type ClientApplicationModel struct {
	DbClient
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

func (c *ClientApplicationModel) updateId(id int) {
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
	c.updateId(clientId)
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

func (c *ClientApplicationModel) GetClient(id int, db *sqlx.DB) error {
	cl := DbClient{}
	if err := db.Get(&cl, "SELECT * FROM clients WHERE id=$1", id); err != nil {
		return err
	}
	c.DbClient = cl
	if err := c.Passport.GetPassport(id, db); err != nil {
		return err
	}
	jobs := []JobApplicationModel{}
	if err := db.Select(&jobs, "SELECT * FROM jobs WHERE client_id=$1", id); err != nil {
		return err
	}
	c.Jobs = jobs
	return nil
}

func (c *ClientApplicationModel) UpdateClient(cl Domain.ClientDomainModelInterface, tx *sql.Tx) error {
	if c.FirstName != cl.GetFirstName() {
		if _, err := tx.Exec("UPDATE clients SET first_name=$2 WHERE id=$1",
			c.Id, cl.GetFirstName()); err != nil {
			return err
		}
	}
	if c.LastName != cl.GetLastName() {
		if _, err := tx.Exec("UPDATE clients SET last_name=$2 WHERE id=$1",
			c.Id, cl.GetLastName()); err != nil {
			return err
		}
	}
	if c.BirthDate != cl.GetBirthDate() {
		if _, err := tx.Exec("UPDATE clients SET birth_date=$2 WHERE id=$1",
			c.Id, cl.GetBirthDate()); err != nil {
			return err
		}
	}
	if c.Sex != cl.GetSex() {
		if _, err := tx.Exec("UPDATE clients SET sex=$2 WHERE id=$1",
			c.Id, cl.GetSex()); err != nil {
			return err
		}
	}
	if c.MaritalStatus != cl.GetMaritalStatus() {
		if _, err := tx.Exec("UPDATE clients SET marital_status=$2 WHERE id=$1",
			c.Id, cl.GetMaritalStatus()); err != nil {
			return err
		}
	}
	if c.Children != cl.GetChildren() {
		if _, err := tx.Exec("UPDATE clients SET children=$2 WHERE id=$1",
			c.Id, cl.GetChildren()); err != nil {
			return err
		}
	}
	if err := c.Passport.UpdatePassport(cl.GetPassport(), tx); err != nil {
		return errors.New("Not unic creditionals")
	}
	for _, domainJob := range cl.GetJobs() {
		exist := false
		for _, job := range c.Jobs {
			if job.Wage == domainJob.GetWage() && job.Name == domainJob.GetName() {
				exist = true
			}
		}
		if !exist {
			newJob := new(JobApplicationModel)
			newJob.Hydrate(domainJob)
			newJob.ClientId = c.Id
			if err := newJob.SaveJob(tx); err != nil {
				return err
			}
		}
	}
	for _, job := range c.Jobs {
		exist := false
		for _, domainJob := range cl.GetJobs() {
			if job.Wage == domainJob.GetWage() && job.Name == domainJob.GetName() {
				exist = true
			}
		}
		if !exist {
			if err := job.DeleteJob(tx); err != nil {
				return err
			}
		}
	}
	return nil
}
