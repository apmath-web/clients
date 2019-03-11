package Domain

type RepositoryInterface interface {
	GetModel(id int) HelloWorldApplicationModel
	PutModel(model HelloWorldApplicationModel) int
}
