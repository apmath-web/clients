package Domain

type RepositoryInterface interface {
	GetUser(id int) UserApplicationModelInterface
	SetUser(model UserApplicationModelInterface) int
	ChangeUser(id int, model UserApplicationModelInterface)
}
