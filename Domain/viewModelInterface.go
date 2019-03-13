package Domain

type UserViewModelInterface interface {
	UserModelInterface
	GetPassport() PassportViewModelInterface
	GetJobs() []JobsViewModelInterface
	Validate() bool
	GetValidation() ValidationInterface
	MarshalJSON() (b []byte, e error)
}

type PassportViewModelInterface interface {
	PassportModelInterface
	MarshalJSON() (b []byte, e error)
}

type JobsViewModelInterface interface {
	JobsModelInterface
	MarshalJSON() (b []byte, e error)
}
