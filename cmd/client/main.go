package main

import (
	"bufio"
	"chatterNats/internal/chat"
	nats2 "chatterNats/pkg/nats"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"strings"
)

const (
	natsURL    = nats.DefaultURL
	CHATSTREAM = "CHAT_STREAM"
	channel    = "general"
)

func main() {
	fmt.Println("Please enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	nc, err := nats2.NatsConnection(natsURL)
	if err != nil {
		log.Fatalf("could not connect to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nats2.JetStreamContext(nc)
	if err != nil {
		log.Fatalf("could not create JetStream context: %v", err)
	}

	chatStream, err := chat.NewChatStream(js, CHATSTREAM)
	if err != nil {
		log.Fatalf("could not create chat stream: %v", err)
	}

	err = chatStream.SubscribeToChannel(channel, func(message *chat.Message) {
		fmt.Printf("\rReceived Message -> %s: %s\n> ", message.Username, message.Content)
	})
	if err != nil {
		log.Fatalf("could not subscribe to channel: %v", err)
	}

	fmt.Printf("Welcome, %s! Start chatting...\n", username)

	// user input
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading input: %v", err)
			break
		}

		text := strings.TrimSpace(input)
		if text == "/quit" {
			break
		}

		// Send the message
		msg := chat.NewMessage(username, text)
		err = chatStream.PublishMessage(channel, msg)
		if err != nil {
			fmt.Printf("error sending message: %v\n", err)
		}
	}
}
