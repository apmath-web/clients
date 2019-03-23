package repositories

import (
	"errors"
	"github.com/apmath-web/clients/Domain"
)

type ClientRepository struct {
	clients         map[int]Domain.ClientDomainModelInterface
	numberOfClients int
}

func (r *ClientRepository) GetClient(id int) (Domain.ClientDomainModelInterface, error) {
	client, ok := r.clients[id]
	if ok {
		return client, nil
	}
	return nil, nil
}
func (r *ClientRepository) SetClient(model Domain.ClientDomainModelInterface) (int, error) {
	r.numberOfClients++
	id := r.numberOfClients
	r.clients[id] = model
	return id, nil
}

func (r *ClientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) error {
	_, ok := r.clients[id]
	if ok {
		r.clients[id] = model
		return nil
	}
	return errors.New("No client with such id")
}
