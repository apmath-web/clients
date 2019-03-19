package repositories

import (
	"errors"
	"github.com/apmath-web/clients/Domain"
)

type ClientRepository struct {
	clients         map[int]Domain.ClientDomainModelInterface
	numberOfClients int
}

func GenRepository() Domain.ClientRepositoryInterface {
	repo := &ClientRepository{make(map[int]Domain.ClientDomainModelInterface), 0}
	return repo
}

func (r *ClientRepository) GetClient(id int) Domain.ClientDomainModelInterface {
	client, ok := r.clients[id]
	if ok {
		return client
	}
	return nil
}
func (r *ClientRepository) SetClient(model Domain.ClientDomainModelInterface) int {
	r.numberOfClients++
	id := r.numberOfClients
	r.clients[id] = model
	return id
}

func (r *ClientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) error {
	_, ok := r.clients[id]
	if ok {
		r.clients[id] = model
		return nil
	} else {
		return errors.New("Invalid id")
	}
}
