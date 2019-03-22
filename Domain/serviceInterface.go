package Domain

type ClientServiceInterface interface {
	Add(model ClientDomainModelInterface) IdInterface
	Get(id IdInterface) ClientDomainModelInterface
	Update(id IdInterface, model ClientDomainModelInterface) error
}
