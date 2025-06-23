package steps

import "github.com/rom8726/airo/config"

type mockProcessor struct{}

func (m *mockProcessor) SetConfig(cfg *config.ProjectConfig) {}
func (m *mockProcessor) Import() string                      { return "" }
func (m *mockProcessor) Config() string                      { return "" }
func (m *mockProcessor) ConfigField() string                 { return "" }
func (m *mockProcessor) ConfigFieldName() string             { return "Postgres" }
func (m *mockProcessor) Constructor() string                 { return "" }
func (m *mockProcessor) InitInAppConstructor() string        { return "" }
func (m *mockProcessor) StructField() string                 { return "" }
func (m *mockProcessor) FillStructField() string             { return "" }
func (m *mockProcessor) Close() string                       { return "" }
func (m *mockProcessor) DockerCompose() string               { return "" }
func (m *mockProcessor) ComposeEnv() string                  { return "" }
func (m *mockProcessor) ConfigEnv() string                   { return "" }
func (m *mockProcessor) MigrateFileData() []byte             { return nil }
