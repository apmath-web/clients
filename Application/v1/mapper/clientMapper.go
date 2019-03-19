package mapper

import (
	"github.com/apmath-web/clients/Application/v1/viewModels"
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
)

func ClientViewMapper(vm viewModels.ClientViewModel) Domain.ClientDomainModelInterface {
	passport := PassportViewMapper(vm.GetPassport())
	var jobs []Domain.JobDomainModelInterface
	for _, job := range vm.GetJobs() {
		jobs = append(jobs, JobViewMapper(job))
	}
	return models.GenClientDomainModel(vm.GetFirstName(), vm.GetLastName(),
		vm.GetBirthDate(), vm.GetSex(), vm.GetMaritalStatus(), vm.GetChildren(), passport, jobs)
}
