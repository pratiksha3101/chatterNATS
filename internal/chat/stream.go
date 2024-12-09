package chat

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
)

type ChatStream struct {
	js     nats.JetStreamContext
	stream string
}

func NewChatStream(js nats.JetStreamContext, streamName string) (*ChatStream, error) {
	_, err := js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{streamName + ".>"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create stream: %v", err)
	}

	return &ChatStream{
		js:     js,
		stream: streamName,
	}, nil
}

func (cs *ChatStream) PublishMessage(channel string, msg *Message) error {
	subject := fmt.Sprintf("%s.%s", cs.stream, channel)

	msgData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}

	_, err = cs.js.Publish(subject, msgData)
	return err
}

func (cs *ChatStream) SubscribeToChannel(channel string, handler func(*Message)) error {
	subject := fmt.Sprintf("%s.%s", cs.stream, channel)

	_, err := cs.js.Subscribe(subject, func(m *nats.Msg) {
		var msg *Message
		if err := json.Unmarshal(m.Data, &msg); err == nil {
			handler(msg)
		}
	})

	return err
}
