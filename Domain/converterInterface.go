package Domain

type ConverterInterface interface {
	FromViewModel(model UserViewModelInterface) UserApplicationModelInterface
	FromApplicationModel(model UserApplicationModelInterface) UserViewModelInterface
}
