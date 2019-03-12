package Domain

type ConvectorInterface interface {
	FromViewModel(model UserViewModel) UserApplicationModel
	FromApplicationModel(model UserApplicationModel) UserViewModel
}
