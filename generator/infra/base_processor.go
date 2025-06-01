package infra

import (
	"bytes"
	"text/template"

	"github.com/rom8726/airo/config"
)

type renderData struct {
	ProjectName string
	Module      string
}

type BaseProcessor struct {
	cfg        *config.ProjectConfig
	renderData renderData
	tmpl       string
}

func (b *BaseProcessor) SetConfig(cfg *config.ProjectConfig) {
	b.cfg = cfg
	b.renderData = renderData{
		ProjectName: cfg.ProjectName,
		Module:      cfg.ModuleName,
	}
}

func (b *BaseProcessor) config() string {
	return b.mustRender("config", b.tmpl, b.renderData)
}

func (b *BaseProcessor) constructor() string {
	return b.mustRender("constructor", b.tmpl, b.renderData)
}

func (b *BaseProcessor) initInAppConstructor() string {
	return b.mustRender("init_in_app_constructor", b.tmpl, b.renderData)
}

func (b *BaseProcessor) close() string {
	return b.mustRender("close", b.tmpl, b.renderData)
}

func (b *BaseProcessor) dockerCompose() string {
	return b.mustRender("docker_compose", b.tmpl, b.renderData)
}

func (b *BaseProcessor) mustRender(name, template string, data any) string {
	var buf bytes.Buffer
	err := mustParse(name, template).Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func mustParse(name, templateData string) *template.Template {
	tmpl, err := template.New(name).Parse(templateData)
	if err != nil {
		panic(err)
	}

	return tmpl
}
