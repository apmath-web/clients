package mapper

import (
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
)

func JobViewMapper(vm viewModels.JobViewModel) Domain.JobDomainModelInterface {
	return models.GenJobDomainModel(vm.GetName(), vm.GetWage())
}
