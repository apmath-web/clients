package repositories

import (
	"clients/Domain"
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
	return nil
}

func (r *ClientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) {
	r.clients[id] = model
}
