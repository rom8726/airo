package infra

import (
	_ "embed"

	"github.com/rom8726/airo/config"
)

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
	return NewDefaultProcessor("",
		// Define imports
		WithImport(func(*config.ProjectConfig) string { return `"github.com/elastic/go-elasticsearch/v8"` }),

		// Define config struct
		WithConfigField("Elasticsearch ElasticsearchConfig `envconfig:\"ELASTICSEARCH\"`"),

		// Define config field name
		WithConfigFieldName("Elasticsearch"),

		// Define app struct field
		WithStructField("ESClient *elasticsearch.Client"),

		// Define how to fill the struct field
		WithFillStructField("ESClient: esClient,"),

		// Define environment variables for docker-compose
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return `
# Elasticsearch
ELASTICSEARCH_URL=http://elasticsearch:9200
ELASTICSEARCH_USERNAME=elastic
ELASTICSEARCH_PASSWORD=changeme`
		}),

		// Define environment variables for local config
		WithConfigEnv(`
# Elasticsearch
ELASTICSEARCH_URL=http://localhost:9200
ELASTICSEARCH_USERNAME=elastic
ELASTICSEARCH_PASSWORD=changeme`),

		// Define docker-compose service
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return `elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.6.0
    environment:
      - discovery.type=single-node
      - xpack.security.enabled=false
    ports:
      - "9200:9200"
    volumes:
      - elasticsearch-data:/usr/share/elasticsearch/data`
		}),
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
