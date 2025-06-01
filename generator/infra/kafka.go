package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const kafkaEnvFormat = `
# Kafka
KAFKA_BROKERS=%s:9092
KAFKA_CLIENT_ID=app`

// WithKafka returns a registry option that adds Kafka support
func WithKafka() RegistryOption {
	return WithInfra(
		"kafka",
		"Kafka",
		NewKafkaProcessor(),
		2,
	)
}

//go:embed templates/kafka.tmpl
var tmplKafka string

// NewKafkaProcessor creates a new processor for Kafka
func NewKafkaProcessor() Processor {
	return NewDefaultProcessor(tmplKafka,
		WithImport(func(*config.ProjectConfig) string {
			return `"github.com/IBM/sarama"`
		}),
		WithConfigField("Kafka Kafka `envconfig:\"KAFKA\"`"),
		WithConfigFieldName("Kafka"),
		WithStructField(`KafkaProducer sarama.SyncProducer
	KafkaClient sarama.Client`),
		WithFillStructField(`KafkaClient: kafkaClient,
	KafkaProducer: kafkaProducer,`),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(kafkaEnvFormat, cfg.ProjectName+"-kafka")
		}),
	)
}
