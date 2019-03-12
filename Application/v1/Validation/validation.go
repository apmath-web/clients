package Validation

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type FirstNameMessage struct {
	text string
}

func (m *FirstNameMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"firstName": m.text,
	})
}

type LastNameMessage struct {
	text string
}

func (m *LastNameMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"lastName": m.text,
	})
}

type BirthDateMessage struct {
	text string
}

func (m *BirthDateMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"birthDate": m.text,
	})
}

type SexMessage struct {
	text string
}

func (m *SexMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"sex": m.text,
	})
}

type MaritalStatusMessage struct {
	text string
}

func (m MaritalStatusMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"maritalStatus": m.text,
	})
}

type ChildrenMessage struct {
	text string
}

func (m *ChildrenMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"children": m.text,
	})
}

type PassportSeriesMessage struct {
	text string
}

func (m *PassportSeriesMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"passportSeries": m.text,
	})
}

type PassportNumberMessage struct {
	text string
}

func (m *PassportNumberMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"passportNumber": m.text,
	})
}

type JobNameMessage struct {
	text string
}

func (m *JobNameMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"jobName": m.text,
	})
}

type JobWageMessage struct {
	text string
}

func (m *JobWageMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"jobWage": m.text,
	})
}

type Validation struct {
	message  string
	messages []Domain.MessageInterface
}

func (v *Validation) AddMessage(message Domain.MessageInterface) {
	v.messages = append(v.messages, message)
}

func (v *Validation) GetMessages() []Domain.MessageInterface {
	return v.messages
}

func (v *Validation) SetMessage(msg string) {
	v.message = msg
}

func (v *Validation) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"message":     v.message,
		"description": v.messages,
	})
}

func GenValidation() Domain.ValidationInterface {
	v := new(Validation)
	return v
}
