package Domain

type ConverterInterface interface {
	FromViewModel(model UserViewModel) UserApplicationModel
	FromApplicationModel(model UserApplicationModel) UserViewModel
}
