package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	etcdEnvFormat = `
# Etcd
ETCD_ENDPOINTS=%s:2379
ETCD_USERNAME=
ETCD_PASSWORD=`
)

// WithEtcd returns a registry option that adds Etcd support
func WithEtcd() RegistryOption {
	return WithInfra(
		"etcd",
		"Etcd",
		NewEtcdProcessor(),
		1,
	)
}

//go:embed templates/etcd.tmpl
var tmplEtcd string

// NewEtcdProcessor creates a new processor for Etcd
func NewEtcdProcessor() Processor {
	return NewDefaultProcessor(tmplEtcd,
		WithImport(func(*config.ProjectConfig) string {
			return `"go.etcd.io/etcd/client/v3"`
		}),
		WithConfigField("Etcd Etcd `envconfig:\"ETCD\"`"),
		WithConfigFieldName("Etcd"),
		WithStructField("EtcdClient *clientv3.Client"),
		WithFillStructField("EtcdClient: etcdClient,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(etcdEnvFormat, cfg.ProjectName+"-etcd")
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(etcdEnvFormat, "localhost") }()),
	)
}
