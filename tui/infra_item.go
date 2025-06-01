package tui

import (
	"fmt"
)

type infraItem struct {
	title string
	code  string
	used  bool
}

func (i infraItem) Title() string {
	checked := "[ ]"
	if i.used {
		checked = "[✓]"
	}
	return fmt.Sprintf("%s %s", checked, i.title)
}
func (i infraItem) Description() string {
	switch i.code {
	case "kafka":
		return "Apache Kafka for event streaming and message queuing"
	case "redis":
		return "Redis for in-memory caching and data structure store"
	case "aerospike":
		return "Aerospike for high-performance NoSQL database"
	case "elasticsearch":
		return "Elasticsearch for full-text search and analytics"
	default:
		return "Infrastructure component"
	}
}
func (i infraItem) FilterValue() string { return i.title }
