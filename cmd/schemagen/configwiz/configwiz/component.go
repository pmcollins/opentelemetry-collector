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

package configwiz

import (
	"fmt"
	"strconv"
	"strings"

	"go.opentelemetry.io/collector/cmd/schemagen/configschema"
)

func componentWizard(lvl int, f *configschema.Field) map[string]interface{} {
	out := map[string]interface{}{}
	p := indentingPrinter{lvl}
	for _, field := range f.Fields {
		if field.Name == "squash" {
			componentWizard(lvl, field)
		} else if field.Kind == "struct" || field.Kind == "ptr" {
			p.println(field.Name)
			out[field.Name] = componentWizard(lvl+1, field)
		} else {
			handleField(p, field, out)
		}
	}
	return out
}

func handleField(p indentingPrinter, field *configschema.Field, out map[string]interface{}) {
	p.println("Field: " + field.Name)
	if field.Doc != "" {
		p.print("Info:  " + strings.ReplaceAll(field.Doc, "\n", " "))
	}
	typ := resolveType(field)
	if typ != "" {
		p.println("Type:  " + typ)
	}
	if field.Default != nil {
		p.println(fmt.Sprintf("Default (enter to accept): %v", field.Default))
	}
	p.print("> ")
	defaultVal := ""
	if field.Default != nil {
		defaultVal = fmt.Sprintf("%v", field.Default)
	}
	in := readline(defaultVal)
	if field.Kind == "bool" {
		b, _ := strconv.ParseBool(in)
		out[field.Name] = b
	} else if in != "" {
		out[field.Name] = in
	}
}

func resolveType(f *configschema.Field) string {
	if f.Type != "" {
		return f.Type
	}
	return f.Kind
}

type indentingPrinter struct {
	level int
}

func (p indentingPrinter) println(s string) {
	p.doPrint(s, "%s%s\n")
}

func (p indentingPrinter) print(s string) {
	p.doPrint(s, "%s%s")
}

func (p indentingPrinter) doPrint(s string, frmt string) {
	const tabSize = 4
	indent := p.level * tabSize
	fmt.Printf(frmt, strings.Repeat(" ", indent), s)
}
