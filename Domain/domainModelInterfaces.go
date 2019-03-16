package Domain

type ClientDomainModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
	GetPassport() PassportDomainModelInterface
	GetJobs() []JobModelDomainInterface
}

type PassportDomainModelInterface interface {
	GetSeries() int
	GetNumber() int
}

type JobModelDomainInterface interface {
	GetName() string
	GetWage() int
}
