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
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"text/template"

	"go.opentelemetry.io/collector/cmd/schemagen/configschema"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
)

func CLI(factories component.Factories) {
	dr := configschema.NewDefaultDirResolver()
	tmpl, err := tpl()
	if err != nil {
		panic(err)
	}

	componentType := os.Args[1]
	if componentType == "all" {
		allComponents(dr, tmpl, factories)
		return
	}

	singleComponent(dr, tmpl, factories, componentType, os.Args[2])
}

func allComponents(dr configschema.DirResolver, tmpl *template.Template, factories component.Factories) {
	configs := configschema.GetAllConfigs(factories)
	for _, cfg := range configs {
		writeConfigDoc(tmpl, dr, cfg)
	}
}

func singleComponent(dr configschema.DirResolver, tmpl *template.Template, factories component.Factories, componentType, componentName string) {
	cfg, err := configschema.GetConfig(factories, componentType, componentName)
	if err != nil {
		panic(err)
	}

	writeConfigDoc(tmpl, dr, cfg)
}

func writeConfigDoc(tmpl *template.Template, dr configschema.DirResolver, config config.NamedEntity) {
	v := reflect.ValueOf(config)
	f := configschema.ReadFields(v, dr)
	bytes, err := render(tmpl, f)
	if err != nil {
		panic(err)
	}

	dir := dr.PackageDir(v.Type().Elem())
	err = ioutil.WriteFile(path.Join(dir, "CONFIG.md"), bytes, 0600)
	if err != nil {
		panic(err)
	}
}
