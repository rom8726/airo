package kafka

import (
	"context"
)

type TopicProducer struct {
	producer *Producer
	topic    string
}

func NewTopicProducer(producer *Producer, topic string) *TopicProducer {
	return &TopicProducer{producer: producer, topic: topic}
}

func (p *TopicProducer) Produce(ctx context.Context, data []byte) error {
	return p.producer.Produce(ctx, p.topic, data)
}
