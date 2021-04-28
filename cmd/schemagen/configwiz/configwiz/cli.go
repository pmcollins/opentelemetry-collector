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
	"sort"

	"gopkg.in/yaml.v2"

	"go.opentelemetry.io/collector/cmd/schemagen/configschema"
	"go.opentelemetry.io/collector/component"
)

func CLI(factories component.Factories) {
	m := map[string]interface{}{}
	services := map[string]interface{}{
		"pipelines": pipelinesWizard(factories),
	}
	m["service"] = services
	componentGroups := getComponents(services)

	dr := configschema.NewDefaultDirResolver()
	for componentType, names := range componentGroups {
		handleComponent(factories, m, componentType, names, dr)
	}

	bytes, _ := yaml.Marshal(m)
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	println(string(bytes))
}

func handleComponent(
	factories component.Factories,
	m map[string]interface{},
	componentType string,
	names []string,
	dr configschema.DirResolver,
) {
	typeMap := map[string]interface{}{}
	m[componentType] = typeMap
	for _, name := range names {
		config, err := configschema.GetConfig(factories, componentType, name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Configuring %s: %q\n", componentType[:len(componentType)-1], name)
		f := configschema.ReadFields(reflect.ValueOf(config), dr)
		typeMap[name] = componentWizard(0, f)
	}
}

func getComponents(services map[string]interface{}) map[string][]string {
	out := map[string][]string{}
	for _, v := range services {
		m := v.(map[string]interface{})
		for _, v2 := range m {
			m2 := v2.(map[string]interface{})
			for componentType, v3 := range m2 {
				l := v3.([]string)
				for _, cmp := range l {
					out[componentType] = append(out[componentType], cmp)
				}
			}
		}
	}
	for _, strings := range out {
		sort.Strings(strings)
	}
	return out
}
