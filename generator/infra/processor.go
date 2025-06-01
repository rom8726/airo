package infra

import (
	"github.com/rom8726/airo/config"
)

// Processor defines the interface for infrastructure component processors
type Processor interface {
	// SetConfig sets the project configuration
	SetConfig(cfg *config.ProjectConfig)

	// Import returns the import statements needed for this component
	Import() string

	// Config returns the configuration struct definition
	Config() string

	// ConfigField returns the field definition for the configuration struct
	ConfigField() string

	// ConfigFieldName returns the name of the configuration field
	ConfigFieldName() string

	// Constructor returns the constructor code
	Constructor() string

	// InitInAppConstructor returns the code to initialize this component in the app constructor
	InitInAppConstructor() string

	// StructField returns the field definition for the app struct
	StructField() string

	// FillStructField returns the code to fill the app struct field
	FillStructField() string

	// Close returns the code to close/cleanup this component
	Close() string

	// DockerCompose returns the docker-compose service definition
	DockerCompose() string

	// ComposeEnv returns the environment variables for docker-compose
	ComposeEnv() string

	// ConfigEnv returns the environment variables for the configuration
	ConfigEnv() string

	// MigrateFileData returns the migration file data
	MigrateFileData() []byte
}

// ProcessorOption is a function that configures a DefaultProcessor
type ProcessorOption func(*DefaultProcessor)

// DefaultProcessor provides default implementations for all Processor methods
type DefaultProcessor struct {
	BaseProcessor

	// Custom fields for overriding default behavior
	importsFn           func(cfg *config.ProjectConfig) string
	composeEnvFn        func(cfg *config.ProjectConfig) string
	configFieldCode     string
	configFieldNameCode string
	structFieldCode     string
	fillStructFieldCode string
	configEnvCode       string
	migrateFileData     []byte
}

// NewDefaultProcessor creates a new DefaultProcessor with the given options
func NewDefaultProcessor(tmpl string, opts ...ProcessorOption) *DefaultProcessor {
	p := &DefaultProcessor{
		BaseProcessor: BaseProcessor{
			tmpl: tmpl,
		},
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// WithImport sets the import code
func WithImport(fn func(cfg *config.ProjectConfig) string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.importsFn = fn
	}
}

// WithConfigField sets the config field code
func WithConfigField(code string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.configFieldCode = code
	}
}

// WithConfigFieldName sets the config field name code
func WithConfigFieldName(code string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.configFieldNameCode = code
	}
}

// WithStructField sets the struct field code
func WithStructField(code string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.structFieldCode = code
	}
}

// WithFillStructField sets the fill struct field code
func WithFillStructField(code string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.fillStructFieldCode = code
	}
}

// WithComposeEnv sets the compose env code
func WithComposeEnv(fn func(cfg *config.ProjectConfig) string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.composeEnvFn = fn
	}
}

// WithConfigEnv sets the config env code
func WithConfigEnv(code string) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.configEnvCode = code
	}
}

// WithMigrateFileData sets the migrate file data
func WithMigrateFileData(data []byte) ProcessorOption {
	return func(p *DefaultProcessor) {
		p.migrateFileData = data
	}
}

// Import implements Processor.Import
func (p *DefaultProcessor) Import() string {
	return p.importsFn(p.cfg)
}

// Config implements Processor.Config
func (p *DefaultProcessor) Config() string {
	return p.config()
}

// ConfigField implements Processor.ConfigField
func (p *DefaultProcessor) ConfigField() string {
	if p.configFieldCode != "" {
		return p.configFieldCode
	}
	return ""
}

// ConfigFieldName implements Processor.ConfigFieldName
func (p *DefaultProcessor) ConfigFieldName() string {
	if p.configFieldNameCode != "" {
		return p.configFieldNameCode
	}
	return ""
}

// Constructor implements Processor.Constructor
func (p *DefaultProcessor) Constructor() string {
	return p.constructor()
}

// InitInAppConstructor implements Processor.InitInAppConstructor
func (p *DefaultProcessor) InitInAppConstructor() string {
	return p.initInAppConstructor()
}

// StructField implements Processor.StructField
func (p *DefaultProcessor) StructField() string {
	if p.structFieldCode != "" {
		return p.structFieldCode
	}
	return ""
}

// FillStructField implements Processor.FillStructField
func (p *DefaultProcessor) FillStructField() string {
	if p.fillStructFieldCode != "" {
		return p.fillStructFieldCode
	}
	return ""
}

// Close implements Processor.Close
func (p *DefaultProcessor) Close() string {
	return p.close()
}

// DockerCompose implements Processor.DockerCompose
func (p *DefaultProcessor) DockerCompose() string {
	return p.dockerCompose()
}

// ComposeEnv implements Processor.ComposeEnv
func (p *DefaultProcessor) ComposeEnv() string {
	return p.composeEnvFn(p.cfg)
}

// ConfigEnv implements Processor.ConfigEnv
func (p *DefaultProcessor) ConfigEnv() string {
	if p.configEnvCode != "" {
		return p.configEnvCode
	}
	return ""
}

// MigrateFileData implements Processor.MigrateFileData
func (p *DefaultProcessor) MigrateFileData() []byte {
	return p.migrateFileData
}
