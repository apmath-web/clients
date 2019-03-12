package Domain

type UserApplicationModelInterface interface {
	GetId() int
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetPassport() PassportApplicationModelInterface
	GetJobs() []JobsApplicationModelInterface
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
	SetFirstName(firstName string)
	SetLastName(lastName string)
	SetBirthDate(birthDate string)
	SetPassport(passport PassportApplicationModelInterface)
	SetJobs(jobs []JobsApplicationModelInterface)
	SetSex(sex string)
	SetMaritalStatus(status string)
	SetChildren(number int)
}

type PassportApplicationModelInterface interface {
	GetSeries() int
	GetNumber() int
	SetNumber(number int)
	SetSeries(series int)
}

type JobsApplicationModelInterface interface {
	GetName() string
	GetWage() int
	SetName(name string)
	SetWage(wage int)
}
