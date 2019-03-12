package Domain

type UserViewModelInterface interface {
	UserModelInterface
	Validate() bool
	GetValidation() ValidationInterface
}

type PassportViewModelInterface interface {
	PassportModelInterface
}

type JobsViewModelInterface interface {
	JobsModelInterface
}
