package Domain

type ClientServiceInterface interface {
	Add(model ClientDomainModelInterface) (IdInterface, error)
	Get(id IdInterface) (ClientDomainModelInterface, error)
	Update(id IdInterface, model ClientDomainModelInterface) error
}
