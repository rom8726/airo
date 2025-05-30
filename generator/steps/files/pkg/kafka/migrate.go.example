package kafka

import (
	"errors"
	"fmt"

	"github.com/IBM/sarama"
)

type CreateTopicRequest struct {
	Topic             string
	Partitions        int32
	ReplicationFactor int16
}

// Migrate creates topics in the kafka.
func Migrate(addrs []string, reqs ...CreateTopicRequest) error {
	adm, err := sarama.NewClusterAdmin(addrs, nil)
	if err != nil {
		return fmt.Errorf("create cluster admin: %w", err)
	}

	defer func() { _ = adm.Close() }()

	for _, req := range reqs {
		details := sarama.TopicDetail{
			NumPartitions:     req.Partitions,
			ReplicationFactor: req.ReplicationFactor,
		}
		err = adm.CreateTopic(req.Topic, &details, false)
		if err != nil {
			if errors.Is(err, sarama.ErrTopicAlreadyExists) {
				continue
			}
			return fmt.Errorf("create topic %q: %w", req.Topic, err)
		}
	}

	return nil
}
