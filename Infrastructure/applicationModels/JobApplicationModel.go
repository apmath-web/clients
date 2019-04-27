package applicationModels

import (
	"database/sql"
	"github.com/apmath-web/clients/Domain"
)

type JobApplicationModel struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Wage     int    `db:"wage"`
	ClientId int    `db:"client_id"`
}

func (j *JobApplicationModel) GetName() string {
	return j.Name
}

func (j *JobApplicationModel) GetWage() int {
	return j.Wage
}

func (j *JobApplicationModel) Hydrate(job Domain.JobDomainModelInterface) {
	j.Wage = job.GetWage()
	j.Name = job.GetName()
}

func (j *JobApplicationModel) SaveJob(tx *sql.Tx) error {
	var jobId int
	if err := tx.QueryRow("INSERT INTO jobs "+
		"(name, wage, client_id) VALUES "+
		"($1, $2, $3) RETURNING id",
		j.GetName(), j.Wage, j.ClientId).Scan(&jobId); err != nil {
		return err
	}
	j.Id = jobId
	return nil
}
