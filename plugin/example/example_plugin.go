package main

import (
	"context"
	"fmt"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
	"github.com/rom8726/airo/plugin"
)

// ExamplePlugin is a simple plugin that demonstrates how to create a plugin for Airo
type ExamplePlugin struct {
	components []plugin.Component
}

// Exported symbol that Airo will look for
var Plugin ExamplePlugin

func init() {
	// Create the plugin instance
	Plugin = ExamplePlugin{
		components: []plugin.Component{
			plugin.NewComponent(
				"example",
				"Example Component",
				infra.InfraComponent,
				NewExampleProcessor(),
			),
		},
	}
}

// ID implements plugin.Plugin.ID
func (p ExamplePlugin) ID() string {
	return "example-plugin"
}

// Name implements plugin.Plugin.Name
func (p ExamplePlugin) Name() string {
	return "Example Plugin"
}

// Version implements plugin.Plugin.Version
func (p ExamplePlugin) Version() string {
	return "1.0.0"
}

// Description implements plugin.Plugin.Description
func (p ExamplePlugin) Description() string {
	return "A simple example plugin for Airo"
}

// Initialize implements plugin.Plugin.Initialize
func (p ExamplePlugin) Initialize(ctx context.Context, options map[string]interface{}) error {
	fmt.Println("Initializing Example Plugin")
	return nil
}

// GetComponents implements plugin.Plugin.GetComponents
func (p ExamplePlugin) GetComponents() []plugin.Component {
	return p.components
}

// ExampleProcessor is a simple processor that demonstrates how to create a custom component
type ExampleProcessor struct {
	infra.BaseProcessor
}

// NewExampleProcessor creates a new ExampleProcessor
func NewExampleProcessor() *ExampleProcessor {
	return &ExampleProcessor{}
}

// SetConfig implements infra.Processor.SetConfig
func (p *ExampleProcessor) SetConfig(cfg *config.ProjectConfig) {
	p.BaseProcessor.SetConfig(cfg)
}

// Import implements infra.Processor.Import
func (p *ExampleProcessor) Import() string {
	return `"github.com/example/example-client"`
}

// Config implements infra.Processor.Config
func (p *ExampleProcessor) Config() string {
	return `
type ExampleConfig struct {
	Host     string ` + "`envconfig:\"HOST\" required:\"true\"`" + `
	Port     int    ` + "`envconfig:\"PORT\" required:\"true\"`" + `
	Username string ` + "`envconfig:\"USERNAME\" required:\"true\"`" + `
	Password string ` + "`envconfig:\"PASSWORD\" required:\"true\"`" + `
}
`
}

// ConfigField implements infra.Processor.ConfigField
func (p *ExampleProcessor) ConfigField() string {
	return "Example ExampleConfig `envconfig:\"EXAMPLE\"`"
}

// ConfigFieldName implements infra.Processor.ConfigFieldName
func (p *ExampleProcessor) ConfigFieldName() string {
	return "Example"
}

// Constructor implements infra.Processor.Constructor
func (p *ExampleProcessor) Constructor() string {
	return `
func newExampleClient(ctx context.Context, cfg *config.ExampleConfig) (*example.Client, error) {
	client, err := example.NewClient(
		example.WithHost(cfg.Host),
		example.WithPort(cfg.Port),
		example.WithCredentials(cfg.Username, cfg.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("create example client: %w", err)
	}

	return client, nil
}
`
}

// InitInAppConstructor implements infra.Processor.InitInAppConstructor
func (p *ExampleProcessor) InitInAppConstructor() string {
	return `
	exampleClient, err := newExampleClient(ctx, &cfg.Example)
	if err != nil {
		return nil, fmt.Errorf("create example client: %w", err)
	}
`
}

// StructField implements infra.Processor.StructField
func (p *ExampleProcessor) StructField() string {
	return "ExampleClient *example.Client"
}

// FillStructField implements infra.Processor.FillStructField
func (p *ExampleProcessor) FillStructField() string {
	return "ExampleClient: exampleClient,"
}

// Close implements infra.Processor.Close
func (p *ExampleProcessor) Close() string {
	return `
	if app.ExampleClient != nil {
		app.ExampleClient.Close()
	}
`
}

// DockerCompose implements infra.Processor.DockerCompose
func (p *ExampleProcessor) DockerCompose() string {
	return `
  {{ .ProjectName }}-example:
    image: 'example/example-service'
    container_name: {{ .ProjectName }}-example
    environment:
      - EXAMPLE_USERNAME=example
      - EXAMPLE_PASSWORD=example
    ports:
      - '8080:8080'
`
}

// ComposeEnv implements infra.Processor.ComposeEnv
func (p *ExampleProcessor) ComposeEnv() string {
	return `
# Example
EXAMPLE_HOST={{ .ProjectName }}-example
EXAMPLE_PORT=8080
EXAMPLE_USERNAME=example
EXAMPLE_PASSWORD=example
`
}

// ConfigEnv implements infra.Processor.ConfigEnv
func (p *ExampleProcessor) ConfigEnv() string {
	return `
# Example
EXAMPLE_HOST=localhost
EXAMPLE_PORT=8080
EXAMPLE_USERNAME=example
EXAMPLE_PASSWORD=example
`
}

// MigrateFileData implements infra.Processor.MigrateFileData
func (p *ExampleProcessor) MigrateFileData() []byte {
	return nil
}
