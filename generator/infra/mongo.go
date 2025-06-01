package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const mongoEnvFormat = `
# MongoDB
MONGO_HOST=%s
MONGO_PORT=27017
MONGO_DATABASE=db
MONGO_PASSWORD=password
MONGO_USER=user`

//go:embed templates/mongodb.tmpl
var tmplMongo string

//go:embed files/cmd/server/migrate_mongodb_go
var tmplMigrateMongo []byte

// WithMongo returns a registry option that adds MongoDB support
func WithMongo() RegistryOption {
	return WithDB(
		config.DBTypeMongoDB,
		"MongoDB",
		NewMongoDBProcessor(),
		3,
	)
}

func NewMongoDBProcessor() Processor {
	return NewDefaultProcessor(tmplMongo,
		WithImport(func(cfg *config.ProjectConfig) string {
			return `"go.mongodb.org/mongo-driver/mongo"
			"go.mongodb.org/mongo-driver/mongo/options"`
		}),
		WithConfigField("Mongo Mongo `envconfig:\"MONGO\"`"),
		WithConfigFieldName("Mongo"),
		WithStructField("MongoClient *mongo.Client"),
		WithFillStructField("MongoClient: mongoClient,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(mongoEnvFormat, cfg.ProjectName+"-mongodb")
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(mongoEnvFormat, "localhost") }()),
		WithMigrateFileData(tmplMigrateMongo),
	)
}
