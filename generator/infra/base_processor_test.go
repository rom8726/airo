package infra

import (
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestBaseProcessor_SetConfig(t *testing.T) {
	b := &BaseProcessor{tmpl: "{{.ProjectName}}-{{.Module}}"}
	cfg := &config.ProjectConfig{ProjectName: "proj", ModuleName: "mod"}
	b.SetConfig(cfg)
	require.Equal(t, "proj", b.renderData.ProjectName)
	require.Equal(t, "mod", b.renderData.Module)
}

func TestBaseProcessor_RenderMethods(t *testing.T) {

	tmpl := `{{define "config"}}C-{{.ProjectName}}{{end}}
{{define "constructor"}}K-{{.Module}}{{end}}
{{define "init_in_app_constructor"}}I-{{.Module}}{{end}}
{{define "close"}}CL-{{.Module}}{{end}}
{{define "docker_compose"}}D-{{.ProjectName}}{{end}}`
	b := &BaseProcessor{tmpl: tmpl}
	cfg := &config.ProjectConfig{ProjectName: "proj", ModuleName: "mod"}
	b.SetConfig(cfg)

	require.Contains(t, b.config(), "C-proj")
	require.Contains(t, b.constructor(), "K-mod")
	require.Contains(t, b.initInAppConstructor(), "I-mod")
	require.Contains(t, b.close(), "CL-mod")
	require.Contains(t, b.dockerCompose(), "D-proj")
}

func TestBaseProcessor_MustParse_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic")
		}
	}()
	mustParse("bad", "{{bad")
}
