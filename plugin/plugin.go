package plugin

import (
	"context"
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/rom8726/airo/generator"
	"github.com/rom8726/airo/generator/infra"
)

// Plugin defines the interface that all Airo plugins must implement
type Plugin interface {
	// ID returns a unique identifier for the plugin
	ID() string

	// Name returns a human-readable name for the plugin
	Name() string

	// Version returns the plugin version
	Version() string

	// Description returns a description of the plugin
	Description() string

	// Initialize initializes the plugin with the given context and options
	Initialize(ctx context.Context, options map[string]interface{}) error

	// GetComponents returns the components provided by this plugin
	GetComponents() []Component
}

// Component defines the interface for a component provided by a plugin
type Component interface {
	// ID returns a unique identifier for the component
	ID() string

	// Name returns a human-readable name for the component
	Name() string

	// Type returns the type of the component (DB or Infra)
	Type() infra.ComponentType

	// GetProcessor returns the processor for this component
	GetProcessor() infra.Processor
}

// Manager handles plugin discovery, loading, and management
type Manager struct {
	plugins     map[string]Plugin
	components  map[string]Component
	pluginsDir  string
	initialized bool
}

// NewManager creates a new plugin manager
func NewManager(pluginsDir string) *Manager {
	return &Manager{
		plugins:    make(map[string]Plugin),
		components: make(map[string]Component),
		pluginsDir: pluginsDir,
	}
}

// Initialize discovers and loads plugins from the plugins directory
func (m *Manager) Initialize(ctx context.Context) error {
	if m.initialized {
		return nil
	}

	// Load plugins from the plugins directory
	if m.pluginsDir != "" {
		if err := m.loadPluginsFromDirectory(ctx, m.pluginsDir); err != nil {
			return err
		}
	}

	m.initialized = true
	return nil
}

// loadPluginsFromDirectory loads all plugins from the given directory
func (m *Manager) loadPluginsFromDirectory(ctx context.Context, dir string) error {
	// Find all .so files in the directory
	matches, err := filepath.Glob(filepath.Join(dir, "*.so"))
	if err != nil {
		return fmt.Errorf("failed to find plugins: %w", err)
	}

	// Load each plugin
	for _, path := range matches {
		if err := m.loadPlugin(ctx, path); err != nil {
			return fmt.Errorf("failed to load plugin %s: %w", path, err)
		}
	}

	return nil
}

// loadPlugin loads a plugin from the given path
func (m *Manager) loadPlugin(ctx context.Context, path string) error {
	// Open the plugin
	p, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open plugin: %w", err)
	}

	// Look up the plugin symbol
	sym, err := p.Lookup("Plugin")
	if err != nil {
		return fmt.Errorf("plugin does not export 'Plugin' symbol: %w", err)
	}

	// Assert that the symbol is a Plugin
	var plg Plugin
	switch s := sym.(type) {
	case *Plugin:
		plg = *s
	case Plugin:
		plg = s
	default:
		return fmt.Errorf("plugin symbol is not a Plugin: %T", sym)
	}

	// Initialize the plugin
	if err := plg.Initialize(ctx, nil); err != nil {
		return fmt.Errorf("failed to initialize plugin: %w", err)
	}

	// Register the plugin
	m.plugins[plg.ID()] = plg

	// Register the plugin's components
	for _, comp := range plg.GetComponents() {
		m.components[comp.ID()] = comp
	}

	return nil
}

// GetPlugins returns all loaded plugins
func (m *Manager) GetPlugins() []Plugin {
	plugins := make([]Plugin, 0, len(m.plugins))
	for _, p := range m.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

// GetComponents returns all components from all plugins
func (m *Manager) GetComponents() []Component {
	components := make([]Component, 0, len(m.components))
	for _, c := range m.components {
		components = append(components, c)
	}
	return components
}

// RegisterWithRegistry registers all components with the given registry
func (m *Manager) RegisterWithRegistry(registry *infra.Registry) {
	for _, comp := range m.components {
		info := infra.ComponentInfo{
			Code:      comp.ID(),
			Title:     comp.Name(),
			Type:      comp.Type(),
			Processor: comp.GetProcessor(),
		}
		registry.RegisterComponent(info)
	}
}

// GetStepProviders returns step providers for all plugins that implement the Step interface
func (m *Manager) GetStepProviders() []generator.StepProvider {
	providers := make([]generator.StepProvider, 0)

	for _, p := range m.plugins {
		// Check if the plugin implements the Step interface
		if step, ok := p.(generator.Step); ok {
			providers = append(providers, func(registry *infra.Registry) generator.Step {
				return step
			})
		}
	}

	return providers
}

// DefaultComponent is a basic implementation of the Component interface
type DefaultComponent struct {
	id        string
	name      string
	compType  infra.ComponentType
	processor infra.Processor
}

// NewComponent creates a new DefaultComponent
func NewComponent(id, name string, compType infra.ComponentType, processor infra.Processor) *DefaultComponent {
	return &DefaultComponent{
		id:        id,
		name:      name,
		compType:  compType,
		processor: processor,
	}
}

// ID implements Component.ID
func (c *DefaultComponent) ID() string {
	return c.id
}

// Name implements Component.Name
func (c *DefaultComponent) Name() string {
	return c.name
}

// Type implements Component.Type
func (c *DefaultComponent) Type() infra.ComponentType {
	return c.compType
}

// GetProcessor implements Component.GetProcessor
func (c *DefaultComponent) GetProcessor() infra.Processor {
	return c.processor
}
