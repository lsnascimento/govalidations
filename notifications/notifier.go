package notifications

type Notifier struct {
	Notifications Notifications
}

func New() *Notifier {
	return &Notifier{
		Notifications: Notifications{},
	}
}

func (notifier *Notifier) GetNotifications() Notifications {
	return notifier.Notifications
}

func (notifier *Notifier) AddNotification(property, message string) {
	notifier.Notifications = append(notifier.Notifications, Notification{property, message})
}

func (notifier *Notifier) Invalid() bool {
	return len(notifier.Notifications) > 0
}

func (notifier *Notifier) Valid() bool {
	return len(notifier.Notifications) == 0
}
