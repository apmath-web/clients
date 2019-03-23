package Infrastructure

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/services"
	"github.com/apmath-web/clients/Infrastructure/repositories"
	"sync"
)

type ServiceManager struct {
	repo           Domain.ClientRepositoryInterface
	onceRepository sync.Once
}

func (sm *ServiceManager) GetClientRepository() Domain.ClientRepositoryInterface {
	sm.onceRepository.Do(func() {
		sm.repo = &repositories.ClientRepository{make(map[int]Domain.ClientDomainModelInterface), 0}
	})
	return sm.repo
}

func (sm *ServiceManager) GetClientService() Domain.ClientServiceInterface {
	service := &services.ClientService{sm.GetClientRepository()}
	return service
}

var serviceManager *ServiceManager
var once sync.Once

func GetServiceManager() *ServiceManager {
	once.Do(func() {
		serviceManager = &ServiceManager{}
	})
	return serviceManager
}
