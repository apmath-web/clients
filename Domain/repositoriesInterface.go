package Domain

type RepositoryInterface interface {
	GetUser(id int) UserApplicationModel
	SetUser(model UserApplicationModel) int
	ChangeUser(id int, model UserApplicationModel)
}
