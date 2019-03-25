package mapper

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
	"github.com/apmath-web/clients/Infrastructure/applicationModels"
)

func JobApplicationMapper(am applicationModels.JobApplicationModel) Domain.JobDomainModelInterface {
	return models.GenJobDomainModel(am.GetName(), am.GetWage())
}
