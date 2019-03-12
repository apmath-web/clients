package Domain

type UserModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetPassport() PassportApplicationModelInterface
	GetJobs() []JobsApplicationModelInterface
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
}

type PassportModelInterface interface {
	GetSeries() int
	GetNumber() int
}

type JobsModelInterface interface {
	GetName() string
	GetWage() int
}
