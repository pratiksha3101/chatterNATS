package main

import (
	"chatterNats/internal/chat"
	nats2 "chatterNats/pkg/nats"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

const (
	natsURL    = nats.DefaultURL
	CHATSTREAM = "CHAT_STREAM"
)

func main() {
	nc, err := nats2.NatsConnection(natsURL)
	if err != nil {
		log.Fatalf("could not connect to NATS : %v", err)
	}
	defer nc.Close()

	js, err := nats2.JetStreamContext(nc)
	if err != nil {
		log.Fatalf("could not connect JetStream context: %v", err)
	}

	chatStream, err := chat.NewChatStream(js, CHATSTREAM)
	if err != nil {
		log.Fatalf("could not create chat stream: %v", err)
	}

	// subscribing to multiple channels
	channels := []string{"general", "random", "help"}
	for _, ch := range channels {
		err = chatStream.SubscribeToChannel(ch, func(msg *chat.Message) {
			log.Printf("[%s] %s: %s", ch, msg.Username, msg.Content)
		})
		if err != nil {
			log.Printf("could not subscribe to channel %s: %v", ch, err)
		}
	}

	fmt.Println("chat server is running...")
	select {}
}
