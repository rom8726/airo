# Generator Architecture

This document describes the architecture of the generator and how to extend it with new stages and technologies.

## Overview

The generator is responsible for creating a new project based on a configuration. It does this by executing a series of steps, each of which performs a specific task such as creating files, directories, or generating code.

The generator also supports various infrastructure components (databases and other services) that can be included in the generated project.

## Key Components

### Generator

The `Generator` is the main entry point for generating a project. It maintains a list of step providers and executes them in order.

```go
// Create a generator with default steps
generator := New(registry)

// Add a custom step
generator.AddStep(func(r *infra.Registry) Step {
    return &MyCustomStep{registry: r}
})

// Replace all steps with custom ones
generator.WithSteps([]StepProvider{
    func(r *infra.Registry) Step { return &MyFirstStep{} },
    func(r *infra.Registry) Step { return &MySecondStep{registry: r} },
})

// Generate the project
generator.GenerateProject(ctx, config)
```

### Step

A `Step` is a unit of work in the generation process. Each step implements the `Step` interface:

```go
type Step interface {
    Description() string
    Do(ctx context.Context, cfg *config.ProjectConfig) error
}
```

To create a new step, implement this interface and add it to the generator using `AddStep` or `WithSteps`.

### Registry

The `Registry` manages the available infrastructure components (databases and other services). It provides methods to register and retrieve components.

```go
// Create a registry with default components
registry := infra.NewRegistry(
    infra.WithPostgres(),
    infra.WithMySQL(),
    infra.WithRedis(),
)

// Register a new component at runtime
registry.RegisterInfra(
    "elasticsearch",
    "Elasticsearch",
    NewElasticsearchProcessor(),
    10,
)
```

### Processor

A `Processor` generates code for a specific infrastructure component. It implements the `Processor` interface, which defines methods for generating various parts of the code.

The `DefaultProcessor` provides default implementations for all methods of the `Processor` interface, making it easier to create new processors by only overriding the methods you need.

```go
// Create a processor for a new technology
func NewElasticsearchProcessor() Processor {
    return NewDefaultProcessor(
        WithImport(`"github.com/elastic/go-elasticsearch/v8"`),
        WithConfigField("Elasticsearch ElasticsearchConfig `envconfig:\"ELASTICSEARCH\"`"),
        WithStructField("ESClient *elasticsearch.Client"),
        // ... other options
    )
}
```

## How to Extend

### Adding a New Step

1. Create a new struct that implements the `Step` interface:

```go
type MyCustomStep struct {
    registry *infra.Registry
}

func (s MyCustomStep) Description() string {
    return "My custom step"
}

func (s MyCustomStep) Do(ctx context.Context, cfg *config.ProjectConfig) error {
    // Implement your step logic here
    return nil
}
```

2. Add the step to the generator:

```go
generator.AddStep(func(r *infra.Registry) Step {
    return &MyCustomStep{registry: r}
})
```

### Adding a New Technology

1. Create a processor for the new technology:

```go
func NewMyTechProcessor() Processor {
    return NewDefaultProcessor(
        WithImport(`"github.com/mytech/client"`),
        WithConfigField("MyTech MyTechConfig `envconfig:\"MYTECH\"`"),
        WithStructField("MyTechClient *mytech.Client"),
        // ... other options
    )
}
```

2. Register the technology with the registry:

```go
// As a registry option
func WithMyTech() RegistryOption {
    return WithInfra(
        "mytech",
        "My Technology",
        NewMyTechProcessor(),
        10,
    )
}

// Or directly
registry.RegisterInfra(
    "mytech",
    "My Technology",
    NewMyTechProcessor(),
    10,
)
```

3. Use the technology in the generator:

```go
registry := infra.NewRegistry(
    // ... other components
    WithMyTech(),
)
```

## Best Practices

1. Use the `DefaultProcessor` for new technologies to minimize boilerplate code.
2. Register components using `RegistryOption` functions for consistency.
3. Keep steps focused on a single responsibility.
4. Use the registry to access infrastructure components in steps.
5. Add detailed comments to your code to help others understand it.