package steps

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rom8726/airo/config"
)

//go:embed files/realtime/api_ws_handler.go.tmpl
var apiWSHandlerTemplate string

//go:embed files/realtime/api_ws_middlewares_ws_auth.go.tmpl
var apiWSMiddlewaresWSAuthTemplate string

//go:embed files/realtime/api_ws_ws_conn.go.tmpl
var apiWSWSConnTemplate string

//go:embed files/realtime/context_context.go.tmpl
var contextContextTemplate string

//go:embed files/realtime/contract_realtime.go.tmpl
var contractRealtimeTemplate string

//go:embed files/realtime/domain_jwt.go.tmpl
var domainJWTTemplate string

//go:embed files/realtime/domain_realtime.go.tmpl
var domainRealtimeTemplate string

//go:embed files/realtime/domain_user.go.tmpl
var domainUserTemplate string

//go:embed files/realtime/services_realtime_conn_manager.go.tmpl
var servicesRealtimeConnManagerTemplate string

//go:embed files/realtime/services_realtime_service.go.tmpl
var servicesRealtimeServiceTemplate string

//go:embed files/realtime/services_tokenizer_tokenizer.go.tmpl
var servicesTokenizerTokenizerTemplate string

type RealtimeStep struct{}

func (RealtimeStep) Description() string {
	return "Generate WebSocket + JWT files"
}

func (RealtimeStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	if !cfg.UseRealtimeJWT {
		slog.Info("skipped")

		return nil
	}

	type renderData struct {
		Module string
	}
	data := renderData{
		Module: cfg.ModuleName,
	}

	// Map of file paths to templates
	files := map[string]string{
		"internal/api/ws/handler.go":                 apiWSHandlerTemplate,
		"internal/api/ws/middlewares/ws_auth.go":     apiWSMiddlewaresWSAuthTemplate,
		"internal/api/ws/ws_conn.go":                 apiWSWSConnTemplate,
		"internal/context/context.go":                contextContextTemplate,
		"internal/contract/realtime.go":              contractRealtimeTemplate,
		"internal/domain/jwt.go":                     domainJWTTemplate,
		"internal/domain/realtime.go":                domainRealtimeTemplate,
		"internal/domain/user.go":                    domainUserTemplate,
		"internal/services/realtime/conn_manager.go": servicesRealtimeConnManagerTemplate,
		"internal/services/realtime/service.go":      servicesRealtimeServiceTemplate,
		"internal/services/tokenizer/tokenizer.go":   servicesTokenizerTokenizerTemplate,
	}

	projectDir := cfg.ProjectName

	for filePath, tmplContent := range files {
		fullPath := filepath.Join(projectDir, filePath)
		dir := filepath.Dir(fullPath)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("mkdir failed for %q: %w", dir, err)
		}

		f, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("create file %q failed: %w", fullPath, err)
		}

		tmpl, err := template.New(filepath.Base(filePath)).Parse(tmplContent)
		if err != nil {
			_ = f.Close()

			return fmt.Errorf("parse template %q failed: %w", filePath, err)
		}

		if err := tmpl.Execute(f, data); err != nil {
			_ = f.Close()

			return fmt.Errorf("execute template %q failed: %w", filePath, err)
		}

		if err := f.Close(); err != nil {
			return fmt.Errorf("close file %q failed: %w", fullPath, err)
		}
	}

	return nil
}
