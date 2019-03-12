package Domain

type Validation interface {
	AddMessage(msg Message)
	GetMessages() []Message
}

type Message interface {
	GetText() string
	SetText() string
}
