package services

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
)

type clientService struct {
	repository Domain.ClientRepositoryInterface
}

func (cs *clientService) Add(model Domain.ClientDomainModelInterface) (Domain.IdInterface, error) {
	id, err := cs.repository.SetClient(model)
	if err != nil {
		return nil, err
	}
	return models.GenId(id), nil
}

func (cs *clientService) Get(id Domain.IdInterface) (Domain.ClientDomainModelInterface, error) {
	model, err := cs.repository.GetClient(id.Get())
	return model, err

}

func (cs *clientService) Update(id Domain.IdInterface, model Domain.ClientDomainModelInterface) error {
	return cs.repository.ChangeClient(id.Get(), model)
}

func GenClientService(repository Domain.ClientRepositoryInterface) Domain.ClientServiceInterface {
	service := new(clientService)
	service.repository = repository
	return service
}
