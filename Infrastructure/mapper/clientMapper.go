package mapper

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
	"github.com/apmath-web/clients/Infrastructure/applicationModels"
)

func ClientApplicationMapper(am applicationModels.ClientApplicationModel) Domain.ClientDomainModelInterface {
	passport := PassportApplicationMapper(am.GetPassport())
	var jobs []Domain.JobDomainModelInterface
	for _, job := range am.GetJobs() {
		jobs = append(jobs, JobApplicationMapper(job))
	}
	return models.GenClientDomainModel(am.GetFirstName(), am.GetLastName(),
		am.GetBirthDate(), am.GetSex(), am.GetMaritalStatus(), am.GetChildren(), passport, jobs)
}
