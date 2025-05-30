package steps

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rom8726/airo/config"
)

//go:embed files/pkg/httpserver/httpserver.go
var httpServerGo []byte

//go:embed files/pkg/httpserver/httpserver_tls.go
var httpServerTLSGo []byte

//go:embed files/pkg/db/tx_ctx.go.example
var txCtxGo []byte

//go:embed files/pkg/db/tx_manager_pg.go.example
var txManagerGoPG []byte

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
	default:
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
