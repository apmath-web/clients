package services

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
)

type ClientService struct {
	Repository Domain.ClientRepositoryInterface
}

func (cs *ClientService) Add(model Domain.ClientDomainModelInterface) (Domain.IdInterface, error) {
	id, err := cs.Repository.SetClient(model)
	if err != nil {
		return nil, err
	}
	return models.GenId(id), nil
}

func (cs *ClientService) Get(id Domain.IdInterface) (Domain.ClientDomainModelInterface, error) {
	model, err := cs.Repository.GetClient(id.Get())
	return model, err

}

func (cs *ClientService) Update(id Domain.IdInterface, model Domain.ClientDomainModelInterface) error {
	return cs.Repository.ChangeClient(id.Get(), model)
}

func GenClientService(repository Domain.ClientRepositoryInterface) Domain.ClientServiceInterface {
	service := new(ClientService)
	service.Repository = repository
	return service
}
