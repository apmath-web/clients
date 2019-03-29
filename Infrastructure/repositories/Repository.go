package repositories

import (
	"errors"
	"github.com/apmath-web/clients/Domain"
	"sync"
)

type clientRepository struct {
	clients         map[int]Domain.ClientDomainModelInterface
	NumberOfClients int
}

func (r *clientRepository) GetClient(id int) (Domain.ClientDomainModelInterface, error) {
	client, ok := r.clients[id]
	if ok {
		return client, nil
	}
	return nil, errors.New("No client with such id")
}
func (r *clientRepository) SetClient(model Domain.ClientDomainModelInterface) (int, error) {
	r.NumberOfClients++
	id := r.NumberOfClients
	r.clients[id] = model
	return id, nil
}

func (r *clientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) error {
	_, ok := r.clients[id]
	if ok {
		r.clients[id] = model
		return nil
	}
	return errors.New("No client with such id")
}

var repo *clientRepository
var once sync.Once

func GenClientRepository() Domain.ClientRepositoryInterface {
	once.Do(func() {
		repo = &clientRepository{make(map[int]Domain.ClientDomainModelInterface), 0}
	})
	return repo
}
