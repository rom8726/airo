package config

type ProjectConfig struct {
	ProjectName string
	ModuleName  string
	OpenAPIPath string
	UsePostgres bool
	UseRedis    bool
}
