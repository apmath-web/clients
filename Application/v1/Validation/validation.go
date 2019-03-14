package Validation

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type Message struct {
	text string
}

type FirstNameMessage struct {
	Message
}

func (m *FirstNameMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"firstName": m.text,
	})
}

type LastNameMessage struct {
	Message
}

func (m *LastNameMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"lastName": m.text,
	})
}

type BirthDateMessage struct {
	Message
}

func (m *BirthDateMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"birthDate": m.text,
	})
}

type SexMessage struct {
	Message
}

func (m *SexMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"sex": m.text,
	})
}

type MaritalStatusMessage struct {
	Message
}

func (m MaritalStatusMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"maritalStatus": m.text,
	})
}

type ChildrenMessage struct {
	Message
}

func (m *ChildrenMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"children": m.text,
	})
}

type PassportSeriesMessage struct {
	Message
}

func (m *PassportSeriesMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"passportSeries": m.text,
	})
}

type PassportNumberMessage struct {
	Message
}

func (m *PassportNumberMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"passportNumber": m.text,
	})
}

type JobNameMessage struct {
	Message
}

func (m *JobNameMessage) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"jobName": m.text,
	})
}

type JobWageMessage struct {
	Message
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

func (v *Validation) Empty() bool {
	return (len(v.messages) == 0) && v.message == ""
}

func GenValidation() Domain.ValidationInterface {
	v := new(Validation)
	return v
}
