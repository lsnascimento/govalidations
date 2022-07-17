package notifications

type Notifier struct {
	Messages Messages
}

func New() *Notifier {
	return &Notifier{
		Messages: Messages{},
	}
}

func (notifier *Notifier) GetMessages() Messages {
	return notifier.Messages
}

func (notifier *Notifier) AddMessage(property, message string) {
	notifier.Messages = append(notifier.Messages, Message{property, message})
}

func (notifier *Notifier) Invalid() bool {
	return len(notifier.Messages) > 0
}

func (notifier *Notifier) Valid() bool {
	return len(notifier.Messages) == 0
}
