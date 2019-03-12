package Domain

type UserModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
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
