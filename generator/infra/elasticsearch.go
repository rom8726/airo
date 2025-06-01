package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const elasticsearchEnvFormat = `
# Elasticsearch
ELASTICSEARCH_URL=http://%s:9200
ELASTICSEARCH_USERNAME=elastic
ELASTICSEARCH_PASSWORD=changeme`

//go:embed templates/elasticsearch.tmpl
var tmplElasticsearch string

// WithElasticsearch returns a registry option that adds Elasticsearch support
func WithElasticsearch() RegistryOption {
	return WithInfra(
		"elasticsearch", // Code used to identify this component
		"Elasticsearch", // Human-readable title
		NewElasticsearchProcessor(),
		3, // Order in the list
	)
}

// NewElasticsearchProcessor creates a new processor for Elasticsearch
func NewElasticsearchProcessor() Processor {
	// Create a DefaultProcessor with custom options
	return NewDefaultProcessor(tmplElasticsearch,
		// Define imports
		WithImport(func(*config.ProjectConfig) string {
			return `"github.com/elastic/go-elasticsearch/v8"
	"net/http"`
		}),

		// Define config struct
		WithConfigField("Elasticsearch Elasticsearch `envconfig:\"ELASTICSEARCH\"`"),

		// Define config field name
		WithConfigFieldName("Elasticsearch"),

		// Define app struct field
		WithStructField("ESClient *elasticsearch.Client"),

		// Define how to fill the struct field
		WithFillStructField("ESClient: esClient,"),

		// Define environment variables for docker-compose
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(elasticsearchEnvFormat, cfg.ProjectName+"-elasticsearch")
		}),

		// Define environment variables for local config
		WithConfigEnv(func() string { return fmt.Sprintf(elasticsearchEnvFormat, "localhost") }()),
	)
}

// RegisterElasticsearch demonstrates how to register a component at runtime
func RegisterElasticsearch(registry *Registry) {
	registry.RegisterInfra(
		"elasticsearch",
		"Elasticsearch",
		NewElasticsearchProcessor(),
		10,
	)
}
