package Domain

type UserViewModel interface {
	GetFirstName() string
	GetLastName() string
	GetBirthDate() string
	GetPassport() PassportViewModel
	GetJobs() []JobsViewModel
	GetSex() string
	GetMaritalStatus() string
	GetChildren() int
	Validate() bool
	GetValidation() Validation
}

type PassportViewModel interface {
	GetSeries() int
	GetNumber() int
}

type JobsViewModel interface {
	GetName() string
	GetWage() int
}
