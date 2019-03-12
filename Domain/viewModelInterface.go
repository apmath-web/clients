package Domain

type UserViewModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetPassport() PassportViewModelInterface
	GetJobs() []JobsViewModelInterface
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
	Validate() bool
	GetValidation() ValidationInterface
	MarshalJSON() (b []byte, e error)
}

type PassportViewModelInterface interface {
	GetSeries() int
	GetNumber() int
	MarshalJSON() (b []byte, e error)
}

type JobsViewModelInterface interface {
	GetName() string
	GetWage() int
	MarshalJSON() (b []byte, e error)
}
