package Infrastructure

import (
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Domain/services"
	"github.com/apmath-web/clients/Infrastructure/repositories"
	"sync"
)

type serviceManagerClass struct {
}

func (sm *serviceManagerClass) GetClientRepository() Domain.ClientRepositoryInterface {
	return repositories.GenClientRepository()
}

func (sm *serviceManagerClass) GetClientService() Domain.ClientServiceInterface {
	service := services.GenClientService(sm.GetClientRepository())
	return service
}

var serviceManager *serviceManagerClass
var once sync.Once

func GetServiceManager() *serviceManagerClass {
	once.Do(func() {
		serviceManager = &serviceManagerClass{}
	})
	return serviceManager
}
