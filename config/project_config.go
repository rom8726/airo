package config

const (
	DBTypePostgres = "postgresql"
	DBTypeMySQL    = "mysql"
	DBTypeMongoDB  = "mongodb"
)

type ProjectConfig struct {
	Aborted        bool
	ProjectName    string
	ModuleName     string
	OpenAPIPath    string
	DB             string
	UseInfra       []string
	UseTesty       bool
	UseRealtimeJWT bool
}
