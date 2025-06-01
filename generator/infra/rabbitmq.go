package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	rabbitmqEnvFormat = `
# RabbitMQ
RABBITMQ_HOST=%s
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_VHOST=`
)

// WithRabbitMQ returns a registry option that adds RabbitMQ support
func WithRabbitMQ() RegistryOption {
	return WithInfra(
		"rabbitmq",
		"RabbitMQ",
		NewRabbitMQProcessor(),
		6,
	)
}

//go:embed templates/rabbitmq.tmpl
var tmplRabbitMQ string

// NewRabbitMQProcessor creates a new processor for RabbitMQ
func NewRabbitMQProcessor() Processor {
	return NewDefaultProcessor(tmplRabbitMQ,
		WithImport(func(*config.ProjectConfig) string {
			return `amqp "github.com/rabbitmq/amqp091-go"`
		}),
		WithConfigField("RabbitMQ RabbitMQ `envconfig:\"RABBITMQ\"`"),
		WithConfigFieldName("RabbitMQ"),
		WithStructField("RabbitMQConn *amqp.Connection"),
		WithFillStructField("RabbitMQConn: rabbitMQConn,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(rabbitmqEnvFormat, cfg.ProjectName+"-rabbitmq")
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(rabbitmqEnvFormat, "localhost") }()),
	)
}
