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
		return "Message broker"
	case "redis":
		return "In-memory cache"
	case "aerospike":
		return "NoSQL database"
	case "elasticsearch":
		return "Search engine"
	case "nats":
		return "Message broker"
	case "rabbitmq":
		return "Message broker"
	case "memcache":
		return "In-memory cache"
	case "etcd":
		return "Key-value store"
	case "mongo":
		return "NoSQL database"
	case "mysql":
		return "SQL database"
	case "postgres":
		return "SQL database"
	default:
		return "Infrastructure component"
	}
}
func (i infraItem) FilterValue() string { return i.title }
