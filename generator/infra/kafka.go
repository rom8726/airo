package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const kafkaEnvFormat = "# Kafka\nKAFKA_BROKERS=%s:9092\nKAFKA_CLIENT_ID=app"

func init() {
	addInfra("kafka", InfraInfo{
		Code:      "kafka",
		Title:     "Kafka",
		Processor: &KafkaProcessor{},
		order:     2,
	})
}

//go:embed templates/kafka.tmpl
var tmplKafka string

type KafkaProcessor struct {
	cfg *config.ProjectConfig
}

func (k *KafkaProcessor) SetConfig(cfg *config.ProjectConfig) {
	k.cfg = cfg
}

func (k *KafkaProcessor) Import() string {
	return `"github.com/IBM/sarama"`
}

func (k *KafkaProcessor) Config() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: k.cfg.ProjectName,
	}

	return render(tmplKafka, "config", renderData)
}

func (k *KafkaProcessor) ConfigField() string {
	return "Kafka Kafka `envconfig:\"KAFKA\"`"
}

func (k *KafkaProcessor) Constructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: k.cfg.ProjectName,
	}

	return render(tmplKafka, "constructor", renderData)
}

func (k *KafkaProcessor) InitInAppConstructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: k.cfg.ProjectName,
	}

	return render(tmplKafka, "init_in_app_constructor", renderData)
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
	renderData := struct {
		ProjectName string
	}{
		ProjectName: k.cfg.ProjectName,
	}

	return render(tmplKafka, "close", renderData)
}

func (k *KafkaProcessor) DockerCompose() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: k.cfg.ProjectName,
	}

	return render(tmplKafka, "docker_compose", renderData)
}

func (k *KafkaProcessor) ComposeEnv() string {
	host := k.cfg.ProjectName + "-kafka"

	return fmt.Sprintf(kafkaEnvFormat, host)
}

func (k *KafkaProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(kafkaEnvFormat, host)
}
