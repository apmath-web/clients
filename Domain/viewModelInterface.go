package Domain

type BasicViewModel interface {
	Validate() bool
	GetValidation() ValidationInterface
	MarshalJSON() (b []byte, e error)
}

type ClientViewModelInterface interface {
	BasicViewModel
	ClientModelInterface
	GetPassport() PassportViewModelInterface
	GetJobs() []JobViewModelInterface
	UnmarshalJSON(b []byte) error
}

type PassportViewModelInterface interface {
	BasicViewModel
	PassportModelInterface
}

type JobViewModelInterface interface {
	JobModelInterface
	BasicViewModel
}
