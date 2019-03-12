package Validation

import "github.com/apmath-web/clients/Domain"

type FirstNameMessage struct {
	firstName string
}

func (fn FirstNameMessage) GetText() string {
	return fn.firstName
}

type LastNameMessage struct {
	lastName string
}

func (ln LastNameMessage) GetText() string {
	return ln.lastName
}

type BirthDateMessage struct {
	birthDate string
}

func (bd BirthDateMessage) GetText() string {
	return bd.birthDate
}

type SexMessage struct {
	sex string
}

func (s SexMessage) GetText() string {
	return s.sex
}

type MaritalStatusMessage struct {
	maritalStatus string
}

func (ms MaritalStatusMessage) GetText() string {
	return ms.maritalStatus
}

type ChildrenMessage struct {
	children string
}

func (c ChildrenMessage) GetText() string {
	return c.children
}

type PassportSeriesMessage struct {
	passportSeries string
}

func (ps PassportSeriesMessage) GetText() string {
	return ps.passportSeries
}

type PassportNumberMessage struct {
	passportNumber string
}

func (pn PassportNumberMessage) GetText() string {
	return pn.passportNumber
}

type JobNameMessage struct {
	jobName string
}

func (jn JobNameMessage) GetText() string {
	return jn.jobName
}

type JobWageMessage struct {
	jobWage string
}

func (jw JobWageMessage) GetText() string {
	return jw.jobWage
}

type Validation struct {
	messages []Domain.MessageInterface
}

func (v *Validation) AddMessage(message Domain.MessageInterface) {
	v.messages = append(v.messages, message)
}

func (v *Validation) GetMessages() []Domain.MessageInterface {
	return v.messages
}
