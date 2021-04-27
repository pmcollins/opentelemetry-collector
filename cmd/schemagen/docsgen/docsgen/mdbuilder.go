// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docsgen

import (
	"bytes"
	"strings"
	"text/template"

	"go.opentelemetry.io/collector/cmd/schemagen/configschema"
)

func render(tmpl *template.Template, field *configschema.Field) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := renderTree(tmpl, field, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

const templateDir = "cmd/schemagen/docsgen/docsgen/templates/"

var tpl = tableTemplate

func tableTemplate() (*template.Template, error) {
	const basename = "table.tmpl"
	tmpl := template.New(basename)
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"join":      join,
			"cleanType": cleanType,
			"hasAnchor": hasAnchor,
		},
	)
	return tmpl.ParseFiles(templateDir + basename)
}

func hasAnchor(kind string) bool {
	return kind == "struct" || kind == "ptr"
}

func join(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}

func cleanType(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "*", ""), ".", "-")
}

func renderTree(tmpl *template.Template, field *configschema.Field, buf *bytes.Buffer) error {
	err := renderField(tmpl, field, buf)
	if err != nil {
		return err
	}
	for _, subField := range field.Fields {
		if subField.Fields == nil {
			continue
		}
		err = renderTree(tmpl, subField, buf)
		if err != nil {
			return err
		}
	}
	return nil
}

func renderField(tmpl *template.Template, field *configschema.Field, buf *bytes.Buffer) error {
	err := tmpl.Execute(buf, field)
	if err != nil {
		return err
	}
	return nil
}
