package infra

import (
	_ "embed"
	"fmt"
)

const kafkaEnvFormat = `
# Kafka
KAFKA_BROKERS=%s:9092
KAFKA_CLIENT_ID=app`

func WithKafka() Opt {
	return func(registry *Registry) {
		registry.addInfra("kafka", &InfraInfo{
			Code:      "kafka",
			Title:     "Kafka",
			Processor: &KafkaProcessor{},
			order:     2,
		})
	}
}

//go:embed templates/kafka.tmpl
var tmplKafka string

type KafkaProcessor struct {
	BaseProcessor
}

func (k *KafkaProcessor) Import() string {
	return `"github.com/IBM/sarama"`
}

func (k *KafkaProcessor) Config() string {
	return k.config(tmplKafka)
}

func (k *KafkaProcessor) ConfigField() string {
	return "Kafka Kafka `envconfig:\"KAFKA\"`"
}

func (k *KafkaProcessor) Constructor() string {
	return k.constructor(tmplKafka)
}

func (k *KafkaProcessor) InitInAppConstructor() string {
	return k.initInAppConstructor(tmplKafka)
}

func (k *KafkaProcessor) StructField() string {
	return `KafkaProducer sarama.SyncProducer
	KafkaClient sarama.Client`
}

func (k *KafkaProcessor) FillStructField() string {
	return `KafkaClient: kafkaClient,
		KafkaProducer: kafkaProducer,`
}

func (k *KafkaProcessor) Close() string {
	return k.close(tmplKafka)
}

func (k *KafkaProcessor) DockerCompose() string {
	return k.dockerCompose(tmplKafka)
}

func (k *KafkaProcessor) ComposeEnv() string {
	host := k.cfg.ProjectName + "-kafka"

	return fmt.Sprintf(kafkaEnvFormat, host)
}

func (k *KafkaProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(kafkaEnvFormat, host)
}
