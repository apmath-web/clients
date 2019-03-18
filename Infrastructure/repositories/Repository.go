package repositories

import (
	//"clients/Domain"
	"clients/Domain"
	"clients/Infrastructure/applicationModels"
)

type Repository struct {
	clients         map[int]applicationModels.ClientApplicationModel
	numberOfClients int
}

func (r *Repository) GetClient(id int) Domain.ClientDomainModelInterface {
	panic("implement me")
}
func (r *Repository) SetClient(model Domain.ClientDomainModelInterface) int {
	panic("implement me")
}
func (r *Repository) ChangeClient(id int, model Domain.ClientDomainModelInterface) {
	panic("implement me")
}
func GenRepository() Domain.ClientRepositoryInterface {
	repo := &Repository{make(map[int]applicationModels.ClientApplicationModel), 0}
	return repo
}
func (r *Repository) Get(id int) applicationModels.ClientApplicationModel {
	client, ok := r.clients[id]
	if ok {
		return client
	}
	return nil
}
func (r *Repository) Store(client applicationModels.ClientApplicationModel) {
	if client.GetId() != -1 {
		panic("Client Has already got ID!")
	}
	r.numberOfClients++
	client.SetId(r.numberOfClients)
	r.clients[r.numberOfClients] = client
}
func (r *Repository) Remove(client applicationModels.ClientApplicationModel) {
	delete(r.clients, client.GetId())
}
