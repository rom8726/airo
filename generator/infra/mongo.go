package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const mongoEnvFormat = "MONGO_HOST=%s\nMONGO_PORT=27017\nMONGO_DATABASE=db\nMONGO_PASSWORD=password\nMONGO_USER=user"

func init() {
	addDB(config.DBTypeMongoDB, DBInfo{
		Code:      config.DBTypeMongoDB,
		Title:     "MongoDB",
		Processor: &MongoProcessor{},
		order:     3,
	})
}

//go:embed templates/mongodb.tmpl
var tmplMongo string

type MongoProcessor struct {
	cfg *config.ProjectConfig
}

func (m *MongoProcessor) SetConfig(cfg *config.ProjectConfig) {
	m.cfg = cfg
}

func (m *MongoProcessor) Import() string {
	return `"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"`
}

func (m *MongoProcessor) Config() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMongo, "config", renderData)
}

func (m *MongoProcessor) ConfigField() string {
	return "Mongo Mongo `envconfig:\"MONGO\"`\n"
}

func (m *MongoProcessor) Constructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMongo, "constructor", renderData)
}

func (m *MongoProcessor) InitInAppConstructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMongo, "init_in_app_constructor", renderData)
}

func (m *MongoProcessor) StructField() string {
	return "MongoClient *mongo.Client"
}

func (m *MongoProcessor) FillStructField() string {
	return "MongoClient: mongoClient,"
}

func (m *MongoProcessor) Close() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMongo, "close", renderData)
}

func (m *MongoProcessor) DockerCompose() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMongo, "docker_compose", renderData)
}

func (m *MongoProcessor) ComposeEnv() string {
	host := m.cfg.ProjectName + "-mongodb"

	return fmt.Sprintf(mongoEnvFormat, host)
}

func (m *MongoProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(mongoEnvFormat, host)
}
