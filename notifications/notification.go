package notifications

type Notification struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

type Notifications []Notification
