package Domain

type UserApplicationModel interface {
	GetId() int
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetPassport() PassportApplicationModel
	GetJobs() []JobsApplicationModel
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
	SetFirstName(firstName string)
	SetLastName(lastName string)
	SetBirthDate(birthDate string)
	SetPassport(passport PassportApplicationModel)
	SetJobs(jobs []JobsApplicationModel)
	SetSex(sex string)
	SetMaritalStatus(status string)
	SetChildren(number int)
}

type PassportApplicationModel interface {
	GetSeries() int
	GetNumber() int
	SetNumber(number int)
	SetSeries(series int)
}

type JobsApplicationModel interface {
	GetName() string
	GetWage() int
	SetName(name string)
	SetWage(wage int)
}
