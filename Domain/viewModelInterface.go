package Domain

type UserViewModelInterface interface {
	UserModelInterface
	GetPassport() PassportViewModelInterface
	GetJobs() []JobsViewModelInterface
	Validate() bool
	GetValidation() ValidationInterface
}

type PassportViewModelInterface interface {
	PassportModelInterface
}

type JobsViewModelInterface interface {
	JobsModelInterface
}
