package Domain

type ClientRepositoryInterface interface {
	GetClient(id int) (ClientDomainModelInterface, error)
	SetClient(model ClientDomainModelInterface) (int, error)
	ChangeClient(id int, model ClientDomainModelInterface) error
}
