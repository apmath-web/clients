package Domain

type ClientRepositoryInterface interface {
	GetClient(id int) ClientDomainModelInterface
	SetClient(model ClientDomainModelInterface) int
	ChangeClient(id int, model ClientDomainModelInterface) error
}
