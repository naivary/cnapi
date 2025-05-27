package main

import (
	"context"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
)

type BrokerConfig struct {
	URL string

	Subject string
}

type Closer interface {
	Close(ctx context.Context) error
}

type Broker struct {
	sender   Closer
	receiver Closer

	senderClient   client.Client
	receiverClient client.Client

	cfg *BrokerConfig
}

func NewBroker(ctx context.Context, cfg BrokerConfig) (*Broker, error) {
	br := &Broker{
		cfg: &cfg,
	}
	if err := br.initSender(ctx); err != nil {
		return nil, err
	}
	if err := br.initReceiver(ctx); err != nil {
		return nil, err
	}
	return br, nil
}

func (b *Broker) initReceiver(ctx context.Context) error {
	receiver, err := cenats.NewConsumer(b.cfg.URL, b.cfg.Subject, cenats.NatsOptions())
	if err != nil {
		return err
	}
	cl, err := cloudevents.NewClient(receiver)
	if err != nil {
		return err
	}
	b.receiverClient = cl
	b.receiver = receiver
	return nil
}

func (b *Broker) initSender(ctx context.Context) error {
	sender, err := cenats.NewSender(b.cfg.URL, b.cfg.Subject, cenats.NatsOptions())
	if err != nil {
		return err
	}
	cl, err := cloudevents.NewClient(sender)
	if err != nil {
		return err
	}
	b.senderClient = cl
	b.sender = sender
	return nil
}

func (b *Broker) Shutdown(ctx context.Context) error {
	if err := b.sender.Close(ctx); err != nil {
		return err
	}
	return b.receiver.Close(ctx)
}
