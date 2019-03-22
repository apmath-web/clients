package services

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/models"
	"github.com/apmath-web/clients/Infrastructure/repositories"
)

type ClientService struct {
	repository Domain.ClientRepositoryInterface
}

func (cs *ClientService) Add(model Domain.ClientDomainModelInterface) (Domain.IdInterface, error) {
	id, err := cs.repository.SetClient(model)
	if err != nil {
		return nil, err
	}
	return models.GenId(id), nil
}

func (cs *ClientService) Get(id Domain.IdInterface) (Domain.ClientDomainModelInterface, error) {
	model, err := cs.repository.GetClient(id.Get())
	return model, err

}

func (cs *ClientService) Update(id Domain.IdInterface, model Domain.ClientDomainModelInterface) error {
	return cs.repository.ChangeClient(id.Get(), model)
}

func (cs *ClientService) GenClientService() Domain.ClientServiceInterface {
	service := new(ClientService)
	service.repository = repositories.GenRepository()
	return service
}
