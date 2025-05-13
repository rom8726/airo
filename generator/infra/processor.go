package infra

import (
	"github.com/rom8726/airo/config"
)

type Processor interface {
	SetConfig(cfg *config.ProjectConfig)

	Import() string
	Config() string
	ConfigField() string
	Constructor() string
	InitInAppConstructor() string
	StructField() string
	FillStructField() string
	Close() string
	DockerCompose() string
}
