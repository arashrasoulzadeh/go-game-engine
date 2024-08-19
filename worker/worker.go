package worker

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func Init() *gochannel.GoChannel {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(true, true),
	)

	messages, err := pubSub.Subscribe(context.Background(), "worker")
	if err != nil {
		panic(err)
	}

	go process(messages)
	return pubSub
}

func process(messages <-chan *message.Message) {
	fmt.Println("starting worker")
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
