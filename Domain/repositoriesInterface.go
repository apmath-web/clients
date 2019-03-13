package Domain

type ClientRepositoryInterface interface {
	GetClient(id int) ClientApplicationModelInterface
	SetClient(model ClientApplicationModelInterface) int
	ChangeClient(id int, model ClientApplicationModelInterface)
}
