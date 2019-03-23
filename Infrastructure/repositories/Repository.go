package repositories

import (
	"errors"
	"github.com/apmath-web/clients/Domain"
)

type ClientRepository struct {
	Clients         map[int]Domain.ClientDomainModelInterface
	NumberOfClients int
}

func (r *ClientRepository) GetClient(id int) (Domain.ClientDomainModelInterface, error) {
	client, ok := r.Clients[id]
	if ok {
		return client, nil
	}
	return nil, nil
}
func (r *ClientRepository) SetClient(model Domain.ClientDomainModelInterface) (int, error) {
	r.NumberOfClients++
	id := r.NumberOfClients
	r.Clients[id] = model
	return id, nil
}

func (r *ClientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) error {
	_, ok := r.Clients[id]
	if ok {
		r.Clients[id] = model
		return nil
	}
	return errors.New("No client with such id")
}
