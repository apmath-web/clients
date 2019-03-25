package mapper

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
	"github.com/apmath-web/clients/Infrastructure/applicationModels"
)

func PassportApplicationMapper(am applicationModels.PassportApplicationModel) Domain.PassportDomainModelInterface {
	return models.GenPassportDomainModel(am.GetSeries(), am.GetNumber())
}
