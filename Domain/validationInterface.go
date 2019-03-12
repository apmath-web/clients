package Domain

type ValidationInterface interface {
	AddMessage(msg MessageInterface)
	GetMessages() []MessageInterface
}

type MessageInterface interface {
	GetText() string
}
