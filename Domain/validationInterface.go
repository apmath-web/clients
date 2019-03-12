package Domain

type ValidationInterface interface {
	AddMessage(msg MessageInterface)
	GetMessages() []MessageInterface
	SetMessage(msg string)
	MarshalJSON() (b []byte, e error)
}

type MessageInterface interface {
	MarshalJSON() (b []byte, e error)
}
