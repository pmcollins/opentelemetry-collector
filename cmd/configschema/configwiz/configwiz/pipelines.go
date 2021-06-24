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
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
)

func pipelinesWizard(factories component.Factories) map[string]interface{} {
	out := map[string]interface{}{}
	for {
		fmt.Printf("Current pipelines: [%s]\n", strings.Join(keys(out), ", "))
		name, rpe := singlePipelineWizard(factories)
		if name == "" {
			break
		}
		out[name] = rpe
	}
	return out
}

func keys(p map[string]interface{}) []string {
	i := 0
	out := make([]string, len(p))
	for k := range p {
		out[i] = k
		i++
	}
	return out
}

func singlePipelineWizard(factories component.Factories) (string, rpe) {
	fmt.Print("Add pipeline (enter to skip)\n")
	fmt.Print("1: Metrics\n")
	fmt.Print("2: Traces\n")
	fmt.Print("> ")
	pipelineID := readline("")
	switch pipelineID {
	case "":
		return "", rpe{}
	case "1":
		return pipelineTypeWizard("metrics", receiverNames(factories, isMetricsReceiver), exporterNames(factories, isMetricsExporter))
	case "2":
		return pipelineTypeWizard("traces", receiverNames(factories, isTracesReceiver), exporterNames(factories, isTracesExporter))
	}
	fmt.Println("Invalid input. Try again.")
	return singlePipelineWizard(factories)
	//return "", rpe{}
}

// pipelineTypeWizard for a given pipelineType (e.g. "metrics", "traces")
func pipelineTypeWizard(
	pipelineType string,
	metricsReceiverNames []string,
	tracesReceiverNames []string,
) (string, rpe) {
	fmt.Printf("%s pipeline extended name (optional) > ", strings.Title(pipelineType))
	name := pipelineType
	nameExt := readline("")
	if nameExt != "" {
		name += "/" + nameExt
	}
	fmt.Printf("Pipeline %q\n", name)
	pr := indentingPrinter{level: 1}
	rpe := rpeWizard(pr, metricsReceiverNames, tracesReceiverNames)
	return name, rpe
}

func rpeWizard(pr indentingPrinter, receiverNames []string, exporterNames []string) rpe {
	out := rpe{}
	out.Receivers = componentListWizard(pr, "receiver", receiverNames)
	out.Exporters = componentListWizard(pr, "exporter", exporterNames)
	return out
}

type rpe struct {
	Receivers  []string
	Processors []string
	Exporters  []string
}

func componentListWizard(pr indentingPrinter, componentGroup string, componentNames []string) (out []string) {
	for {
		pr.println(fmt.Sprintf("Current %ss: [%s]", componentGroup, strings.Join(out, ", ")))
		key, name := componentNameWizard(pr, componentGroup, componentNames)
		if key == "" {
			break
		}
		if name != "" {
			key += "/" + name
		}
		out = append(out, key)
	}
	return
}

func componentNameWizard(pr indentingPrinter, componentType string, componentNames []string) (string, string) {
	pr.println(fmt.Sprintf("Add %s (enter to skip)", componentType))
	for i, name := range componentNames {
		pr.println(fmt.Sprintf("%d: %s", i, name))
	}
	pr.print("> ")
	choice := readline("")
	if choice == "" {
		return "", ""
	}
	i, _ := strconv.Atoi(choice)
	if i < 0 || i > len(componentNames)-1 {
		fmt.Println("Index passed in is out of range. Try again.")
		return componentNameWizard(pr, componentType, componentNames)
	}
	key := componentNames[i]
	pr.print(fmt.Sprintf("%s %s extended name (optional) > ", key, componentType))
	return key, readline("")
}

type receiverFactoryTest func(factory component.ReceiverFactory) bool

type exporterFactoryTest func(factory component.ExporterFactory) bool

func receiverNames(c component.Factories, test receiverFactoryTest) []string {
	var keys []string
	for k, v := range c.Receivers {
		if test(v) {
			keys = append(keys, string(k))
		}
	}
	sort.Strings(keys)
	return keys
}

func isTracesReceiver(f component.ReceiverFactory) bool {
	_, err := f.CreateTracesReceiver(
		context.Background(),
		component.ReceiverCreateParams{},
		f.CreateDefaultConfig(),
		nil,
	)
	return err != componenterror.ErrDataTypeIsNotSupported
}

func isMetricsReceiver(f component.ReceiverFactory) bool {
	_, err := f.CreateMetricsReceiver(
		context.Background(),
		component.ReceiverCreateParams{},
		f.CreateDefaultConfig(),
		nil,
	)
	return err != componenterror.ErrDataTypeIsNotSupported
}

func exporterNames(c component.Factories, test exporterFactoryTest) []string {
	var exporters []string
	for k, v := range c.Exporters {
		if test(v) {
			exporters = append(exporters, string(k))
		}
	}
	sort.Strings(exporters)
	return exporters
}

func isMetricsExporter(f component.ExporterFactory) bool {
	_, err := f.CreateMetricsExporter(context.Background(), component.ExporterCreateParams{}, f.CreateDefaultConfig())
	return err != componenterror.ErrDataTypeIsNotSupported
}

func isTracesExporter(f component.ExporterFactory) bool {
	_, err := f.CreateTracesExporter(context.Background(), component.ExporterCreateParams{}, f.CreateDefaultConfig())
	return err != componenterror.ErrDataTypeIsNotSupported
}