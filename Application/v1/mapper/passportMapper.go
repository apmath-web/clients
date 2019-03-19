package mapper

import (
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
)

func PassportViewMapper(vm viewModels.PassportViewModel) Domain.PassportDomainModelInterface {
	return models.GenPassportDomainModel(vm.GetSeries(), vm.GetNumber())
}
