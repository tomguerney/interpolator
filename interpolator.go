package interpolator

import (
	"strings"
	"text/template"
)

func interpolate(tmpl string, values interface{}) (string, error) {
	builder := strings.Builder{}
	parsed, err := template.New("template").Parse(tmpl)
	if err != nil {
		return "", err
	}
	err = parsed.Execute(&builder, values)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}
