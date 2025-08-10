package assets

import _ "embed"

// ExampleServerYAML contains the built-in ping-pong server OpenAPI spec.
//
//go:embed server.yml
var ExampleServerYAML []byte

// EmbeddedOpenAPIPath is a special marker indicating to use the embedded ExampleServerYAML
// instead of a file path on disk.
const EmbeddedOpenAPIPath = "__embedded_openapi__"
