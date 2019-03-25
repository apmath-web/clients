package Domain

type ClientDomainModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
	GetPassport() PassportDomainModelInterface
	GetJobs() []JobDomainModelInterface
}

type PassportDomainModelInterface interface {
	GetSeries() int
	GetNumber() int
}

type JobDomainModelInterface interface {
	GetName() string
	GetWage() int
}

type IdInterface interface {
	Get() int
	Set(id int)
}
