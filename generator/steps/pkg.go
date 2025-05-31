package steps

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rom8726/airo/config"
)

// http server
//
//go:embed files/pkg/httpserver/httpserver.go
var httpServerGo []byte

//go:embed files/pkg/httpserver/httpserver_tls.go
var httpServerTLSGo []byte

// db
//
//go:embed files/pkg/db/tx_ctx.go.example
var txCtxGo []byte

//go:embed files/pkg/db/tx_manager_pg.go.example
var txManagerGoPG []byte

//go:embed files/pkg/db/tx_manager_mysql.go.example
var txManagerGoMySQL []byte

// kafka
//
//go:embed files/pkg/kafka/consumer.go.example
var consumerGo []byte

//go:embed files/pkg/kafka/producer.go.example
var producerGo []byte

//go:embed files/pkg/kafka/consumer_group_handler.go.example
var consumerGroupHandlerGo []byte

//go:embed files/pkg/kafka/messages_pool.go.example
var messagesPoolGo []byte

//go:embed files/pkg/kafka/migrate.go.example
var migrateGo []byte

//go:embed files/pkg/kafka/topic_producer.go.example
var topicProducerGo []byte

type PkgStep struct{}

func (PkgStep) Description() string {
	return "Create pkg directory"
}

func (PkgStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := pkgHttpServerDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	httpserverPath := filepath.Join(dir, "httpserver.go")
	if err := os.WriteFile(httpserverPath, httpServerGo, 0644); err != nil {
		return fmt.Errorf("failed to write httpserver.go: %w", err)
	}

	httpserverTLSPath := filepath.Join(dir, "httpserver_tls.go")
	if err := os.WriteFile(httpserverTLSPath, httpServerTLSGo, 0644); err != nil {
		return fmt.Errorf("failed to write httpserver_tls.go: %w", err)
	}

	switch cfg.DB {
	case config.DBTypePostgres:
		if err := copyPkgDB(cfg, txManagerGoPG); err != nil {
			return fmt.Errorf("failed to copy pkg/db: %w", err)
		}
	case config.DBTypeMySQL:
		if err := copyPkgDB(cfg, txManagerGoMySQL); err != nil {
			return fmt.Errorf("failed to copy pkg/db: %w", err)
		}
	default:
	}

	if err := copyPkgKafka(cfg); err != nil {
		return fmt.Errorf("failed to copy pkg/kafka: %w", err)
	}

	return nil
}

func copyPkgDB(cfg *config.ProjectConfig, txManagerGoData []byte) error {
	dir := pkgDBDir(cfg)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	txCtxFilePath := filepath.Join(dir, "tx_ctx.go")
	if err := os.WriteFile(txCtxFilePath, txCtxGo, 0644); err != nil {
		return fmt.Errorf("failed to write tx_ctx.go: %w", err)
	}

	txManagerFilePath := filepath.Join(dir, "tx_manager.go")
	if err := os.WriteFile(txManagerFilePath, txManagerGoData, 0644); err != nil {
		return fmt.Errorf("failed to write tx_manager.go: %w", err)
	}

	return nil
}

func copyPkgKafka(cfg *config.ProjectConfig) error {
	dir := pkgKafkaDir(cfg)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	consumerFilePath := filepath.Join(dir, "consumer.go")
	if err := os.WriteFile(consumerFilePath, consumerGo, 0644); err != nil {
		return fmt.Errorf("failed to write consumer.go: %w", err)
	}

	producerFilePath := filepath.Join(dir, "producer.go")
	if err := os.WriteFile(producerFilePath, producerGo, 0644); err != nil {
		return fmt.Errorf("failed to write producer.go: %w", err)
	}

	consumerGroupHandlerFilePath := filepath.Join(dir, "consumer_group_handler.go")
	if err := os.WriteFile(consumerGroupHandlerFilePath, consumerGroupHandlerGo, 0644); err != nil {
		return fmt.Errorf("failed to write consumer_group_handler.go: %w", err)
	}

	messagesPoolFilePath := filepath.Join(dir, "messages_pool.go")
	if err := os.WriteFile(messagesPoolFilePath, messagesPoolGo, 0644); err != nil {
		return fmt.Errorf("failed to write messages_pool.go: %w", err)
	}

	migrateFilePath := filepath.Join(dir, "migrate.go")
	if err := os.WriteFile(migrateFilePath, migrateGo, 0644); err != nil {
		return fmt.Errorf("failed to write migrate.go: %w", err)
	}

	topicProducerFilePath := filepath.Join(dir, "topic_producer.go")
	if err := os.WriteFile(topicProducerFilePath, topicProducerGo, 0644); err != nil {
		return fmt.Errorf("failed to write topic_producer.go: %w", err)
	}

	return nil
}
