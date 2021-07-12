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
	"reflect"
	"strconv"
	"strings"

	"go.opentelemetry.io/collector/cmd/configschema/configschema"
	"go.opentelemetry.io/collector/component"
)

func serviceToComponentNames(service map[string]interface{}) map[string][]string {
	out := map[string][]string{}
	for _, v := range service {
		m := v.(map[string]interface{})
		for _, v2 := range m {
			r := v2.(rpe)
			for _, receiver := range r.Receivers {
				out["receiver"] = append(out["receiver"], receiver)
			}
			for _, processor := range r.Processors {
				out["processor"] = append(out["processor"], processor)
			}
			for _, exporter := range r.Exporters {
				out["exporter"] = append(out["exporter"], exporter)
			}
			for _, extension := range r.Extensions {
				out["extension"] = append(out["extension"], extension)
			}
		}
	}
	return out
}

func handleComponent(
	factories component.Factories,
	m map[string]interface{},
	componentGroup string,
	names []string,
	dr configschema.DirResolver,
) {
	typeMap := map[string]interface{}{}
	m[componentGroup+"s"] = typeMap
	for _, name := range names {
		cfgInfo, err := configschema.GetCfgInfo(factories, componentGroup, strings.Split(name, "/")[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %q\n", strings.Title(componentGroup), name)
		f := configschema.ReadFields(reflect.ValueOf(cfgInfo.CfgInstance), dr)
		typeMap[name] = componentWizard(0, f)
	}
}

func componentWizard(lvl int, f *configschema.Field) map[string]interface{} {
	out := map[string]interface{}{}
	io := clio{printLine, readline}
	p := io.newIndentingPrinter(lvl)
	for _, field := range f.Fields {
		if field.Name == "squash" {
			componentWizard(lvl, field)
		} else if field.Kind == "struct" {
			p.println(field.Name)
			out[field.Name] = componentWizard(lvl+1, field)
		} else if field.Kind == "ptr" {
			p.print(fmt.Sprintf("%s (optional) skip (Y/n)> ", field.Name))
			in := readline("")
			if in == "n" {
				out[field.Name] = componentWizard(lvl+1, field)
			}
		} else {
			handleField(io, p, field, out)
		}
	}
	return out
}

func handleField(io clio, p indentingPrinter2, field *configschema.Field, out map[string]interface{}) {
	p.println("Field: " + field.Name)
	typ := resolveType(field)
	if typ != "" {
		typString := "Type: " + typ
		if typ == "time.Duration" {
			typString += " (examples: 1h2m3s, 5m10s, 45s)"
		}
		p.println(typString)
	}
	if field.Doc != "" {
		p.println("Docs: " + strings.ReplaceAll(field.Doc, "\n", " "))
	}
	if field.Default != nil {
		p.println(fmt.Sprintf("Default (enter to accept): %v", field.Default))
	}
	p.print("> ")
	defaultVal := ""
	if field.Default != nil {
		defaultVal = fmt.Sprintf("%v", field.Default)
	}
	in := io.read(defaultVal)
	if in == "" {
		return
	}
	switch field.Kind {
	case "bool":
		b, _ := strconv.ParseBool(in)
		out[field.Name] = b
	case "int", "int8", "int16", "int32", "int64":
		atoi, _ := strconv.Atoi(in)
		out[field.Name] = atoi
	case "float32", "float64":
		f, _ := strconv.ParseFloat(in, 10)
		out[field.Name] = f
	case "[]string":
		out[field.Name] = parseCSV(in)
	default:
		out[field.Name] = in
	}
}

func parseCSV(in string) []string {
	a := strings.Split(in, ",")
	for i, s := range a {
		a[i] = strings.TrimSpace(s)
	}
	return a
}

func resolveType(f *configschema.Field) string {
	if f.Type != "" {
		return f.Type
	}
	return f.Kind
}
