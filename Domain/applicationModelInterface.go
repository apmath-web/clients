package Domain

type UserApplicationModelInterface interface {
	GetId() int
	UserModelInterface
}

type PassportApplicationModelInterface interface {
	PassportModelInterface
}

type JobsApplicationModelInterface interface {
	JobsModelInterface
}
