package config

type ProjectConfig struct {
	Aborted     bool
	ProjectName string
	ModuleName  string
	OpenAPIPath string
	UsePostgres bool
	UseMySQL    bool
	UseRedis    bool
	UseKafka    bool
}
