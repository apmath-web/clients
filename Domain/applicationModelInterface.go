package Domain

type UserApplicationModelInterface interface {
	GetId() int
	GetPassport() PassportApplicationModelInterface
	GetJobs() []JobsApplicationModelInterface
	UserModelInterface
}

type PassportApplicationModelInterface interface {
	PassportModelInterface
}

type JobsApplicationModelInterface interface {
	JobsModelInterface
}
