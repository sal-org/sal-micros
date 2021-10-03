package model

// ServerResponse .
type ServerResponse struct {
	Meta Meta `json:"meta"`
}

// Meta .
type Meta struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}
