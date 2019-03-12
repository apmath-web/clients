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
}

type PassportViewModelInterface interface {
	GetSeries() int
	GetNumber() int
}

type JobsViewModelInterface interface {
	GetName() string
	GetWage() int
}
