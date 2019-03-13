package Domain

type ClientApplicationModelInterface interface {
	GetId() int
	GetPassport() PassportApplicationModelInterface
	GetJobs() []JobApplicationModelInterface
	ClientModelInterface
}

type PassportApplicationModelInterface interface {
	PassportModelInterface
}

type JobApplicationModelInterface interface {
	JobModelInterface
}
