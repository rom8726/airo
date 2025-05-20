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

func WithMongo() Opt {
	return func(registry *Registry) {
		registry.addDB(config.DBTypeMongoDB, &DBInfo{
			Code:      config.DBTypeMongoDB,
			Title:     "MongoDB",
			Processor: &MongoProcessor{},
			order:     3,
		})
	}
}

//go:embed templates/mongodb.tmpl
var tmplMongo string

//go:embed files/cmd/server/migrate_mongodb_go
var tmplMigrateMongo []byte

type MongoProcessor struct {
	BaseProcessor
}

func (m *MongoProcessor) Import() string {
	return `"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"`
}

func (m *MongoProcessor) Config() string {
	return m.config(tmplMongo)
}

func (m *MongoProcessor) ConfigField() string {
	return "Mongo Mongo `envconfig:\"MONGO\"`\n"
}

func (m *MongoProcessor) ConfigFieldName() string {
	return "Mongo"
}

func (m *MongoProcessor) Constructor() string {
	return m.constructor(tmplMongo)
}

func (m *MongoProcessor) InitInAppConstructor() string {
	return m.initInAppConstructor(tmplMongo)
}

func (m *MongoProcessor) StructField() string {
	return "MongoClient *mongo.Client"
}

func (m *MongoProcessor) FillStructField() string {
	return "MongoClient: mongoClient,"
}

func (m *MongoProcessor) Close() string {
	return m.close(tmplMongo)
}

func (m *MongoProcessor) DockerCompose() string {
	return m.dockerCompose(tmplMongo)
}

func (m *MongoProcessor) ComposeEnv() string {
	host := m.cfg.ProjectName + "-mongodb"

	return fmt.Sprintf(mongoEnvFormat, host)
}

func (m *MongoProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(mongoEnvFormat, host)
}

func (m *MongoProcessor) MigrateFileData() []byte {
	return tmplMigrateMongo
}
