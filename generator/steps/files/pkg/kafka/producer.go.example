package kafka

import (
	"context"
	"log/slog"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.AsyncProducer
}

func NewProducer(addrs []string) (*Producer, error) {
	producerConfig := sarama.NewConfig()
	producerConfig.Version = sarama.MaxVersion
	producerConfig.Producer.RequiredAcks = sarama.WaitForLocal
	producerConfig.Producer.Return.Errors = true
	producerConfig.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(addrs, producerConfig)
	if err != nil {
		return nil, err
	}

	p := &Producer{producer: producer}

	go p.dispatch()

	return p, nil
}

func (p *Producer) Produce(ctx context.Context, topic string, data []byte) error {
	select {
	case <-ctx.Done():
		return context.Canceled
	case p.producer.Input() <- newMessage(topic, data):
		return nil
	}
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

func (p *Producer) dispatch() {
	for {
		select {
		case msg, ok := <-p.producer.Successes():
			if !ok {
				return
			}
			freeMessage(msg)
		case err, ok := <-p.producer.Errors():
			if !ok {
				return
			}
			freeMessage(err.Msg)
			slog.Error(err.Error())
		}
	}
}
