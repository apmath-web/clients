package Domain

type ClientModelInterface interface {
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

type JobModelInterface interface {
	GetName() string
	GetWage() int
}
