package infra

import (
	"strings"
	"text/template"
)

func render(tmplData string, name string, data any) string {
	tmpl, err := template.New(name).Parse(tmplData)
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	err = tmpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
