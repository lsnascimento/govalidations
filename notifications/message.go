package notifications

type Message struct {
	Property    string `json:"property"`
	Description string `json:"description"`
}

type Messages []Message
