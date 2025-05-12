package infra

type Processor interface {
	Import() string
	Config() string
	ConfigField() string
	Constructor() string
	InitInAppConstructor() string
	StructField() string
	FillStructField() string
	Close() string
}
