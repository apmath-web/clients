package Domain

type ClientViewModelInterface interface {
	ClientModelInterface
	GetPassport() PassportViewModelInterface
	GetJobs() []JobViewModelInterface
	Validate() bool
	GetValidation() ValidationInterface
	MarshalJSON() (b []byte, e error)
	UnmarshalJSON(b []byte) error
}

type PassportViewModelInterface interface {
	PassportModelInterface
	MarshalJSON() (b []byte, e error)
	UnmarshalJSON(b []byte) error
}

type JobViewModelInterface interface {
	JobModelInterface
	MarshalJSON() (b []byte, e error)
	UnmarshalJSON(b []byte) error
}
