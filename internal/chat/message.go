package chat

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func NewMessage(username, content string) *Message {
	return &Message{
		ID:        uuid.New().String(),
		Username:  username,
		Content:   content,
		Timestamp: time.Now(),
	}
}
