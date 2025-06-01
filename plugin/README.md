# Airo Plugin System

The Airo plugin system allows you to extend the functionality of Airo with custom components without modifying the core codebase. This document explains how to create and use plugins.

## Using Plugins

To use plugins with Airo, specify the plugins directory using the `--plugins-dir` flag:

```bash
airo generate --plugins-dir /path/to/plugins
```

Airo will load all plugins from the specified directory and register their components with the generator.

## Creating Plugins

Plugins are Go shared libraries (.so files) that implement the `Plugin` interface. Here's how to create a plugin:

1. Create a new Go module for your plugin:

```bash
mkdir my-plugin
cd my-plugin
go mod init github.com/yourusername/my-plugin
```

2. Create a main.go file that implements the `Plugin` interface:

```go
package main

import (
	"context"
	"fmt"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
	"github.com/rom8726/airo/plugin"
)

// MyPlugin is a plugin for Airo
type MyPlugin struct {
	components []plugin.Component
}

// Exported symbol that Airo will look for
var Plugin MyPlugin

func init() {
	// Create the plugin instance
	Plugin = MyPlugin{
		components: []plugin.Component{
			plugin.NewComponent(
				"my-component",
				"My Component",
				infra.InfraComponent,
				NewMyProcessor(),
			),
		},
	}
}

// ID implements plugin.Plugin.ID
func (p MyPlugin) ID() string {
	return "my-plugin"
}

// Name implements plugin.Plugin.Name
func (p MyPlugin) Name() string {
	return "My Plugin"
}

// Version implements plugin.Plugin.Version
func (p MyPlugin) Version() string {
	return "1.0.0"
}

// Description implements plugin.Plugin.Description
func (p MyPlugin) Description() string {
	return "My custom plugin for Airo"
}

// Initialize implements plugin.Plugin.Initialize
func (p MyPlugin) Initialize(ctx context.Context, options map[string]interface{}) error {
	fmt.Println("Initializing My Plugin")
	return nil
}

// GetComponents implements plugin.Plugin.GetComponents
func (p MyPlugin) GetComponents() []plugin.Component {
	return p.components
}

// MyProcessor is a custom processor for my component
type MyProcessor struct {
	infra.BaseProcessor
}

// NewMyProcessor creates a new MyProcessor
func NewMyProcessor() *MyProcessor {
	return &MyProcessor{}
}

// Implement the infra.Processor interface methods...
```

3. Build the plugin as a shared library:

```bash
go build -buildmode=plugin -o my-plugin.so main.go
```

4. Copy the shared library to the plugins directory:

```bash
cp my-plugin.so /path/to/plugins/
```

## Plugin Interface

Plugins must implement the `Plugin` interface:

```go
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
```

## Component Interface

Components provided by plugins must implement the `Component` interface:

```go
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
```

The `DefaultComponent` implementation can be used as a base for custom components:

```go
plugin.NewComponent(
	"my-component",
	"My Component",
	infra.InfraComponent,
	NewMyProcessor(),
)
```

## Processor Interface

Components must provide a processor that implements the `infra.Processor` interface. The processor generates code for the component. See the `infra.Processor` interface for details.

## Example Plugin

See the [example plugin](example/example_plugin.go) for a complete example of how to create a plugin.