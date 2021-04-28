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
	"strconv"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
)

func pipelinesWizard(factories component.Factories) map[string]interface{} {
	out := map[string]interface{}{}
	for {
		name, rpe := pipelineItemWizard(factories)
		if name == "" {
			break
		}
		out[name] = rpe
	}
	return out
}

func pipelineItemWizard(factories component.Factories) (string, map[string]interface{}) {
	println("Add pipeline:")
	println("1: Metrics")
	println("2: Traces")
	print("> ")
	pipeline := readline("")
	if pipeline == "1" {
		return pipelineTypeWizard(factories, isMetricsReceiver, isMetricsExporter, "metrics")
	} else if pipeline == "2" {
		return pipelineTypeWizard(factories, isTracesReceiver, isTracesExporter, "traces")
	}
	return "", nil
}

func pipelineTypeWizard(factories component.Factories, recevierTest receiverFactoryTest, exporterTest exporterFactoryTest, pipelineType string) (string, map[string]interface{}) {
	fmt.Printf("Setting up %s pipeline\n", pipelineType)
	print("Set extended name (optional): ")
	name := pipelineType
	nameExt := readline("")
	if nameExt != "" {
		name += "/" + nameExt
	}
	rpe := rpeWizard(factories, recevierTest, exporterTest)
	return name, rpe
}

func rpeWizard(
	factories component.Factories,
	rt receiverFactoryTest,
	et exporterFactoryTest,
) map[string]interface{} {
	out := map[string]interface{}{}
	out["receivers"] = receiverListWizard(factories, rt)
	out["exporters"] = exporterListWizard(factories, et)
	return out
}

func receiverListWizard(factories component.Factories, t receiverFactoryTest) (out []string) {
	receivers := receiverNames(factories, t)
	for {
		key, name := componentNameWizard("receiver", receivers)
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

func exporterListWizard(factories component.Factories, t exporterFactoryTest) (out []string) {
	receivers := exporterNames(factories, t)
	for {
		key, name := componentNameWizard("exporter", receivers)
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

func componentNameWizard(componentType string, componentNames []string) (string, string) {
	fmt.Printf("Add %s:\n", componentType)
	for i, name := range componentNames {
		fmt.Printf("%d: %s\n", i, name)
	}
	printPrompt()
	choice := readline("")
	if choice == "" {
		return "", ""
	}
	i, _ := strconv.Atoi(choice)
	key := componentNames[i]
	fmt.Printf("Set %s %s extended name (optional):", key, componentType)
	return key, readline("")
}

func printPrompt() {
	print("> ")
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
	var keys []string
	for k, v := range c.Exporters {
		if test(v) {
			keys = append(keys, string(k))
		}
	}
	return keys
}

func isMetricsExporter(f component.ExporterFactory) bool {
	_, err := f.CreateMetricsExporter(context.Background(), component.ExporterCreateParams{}, f.CreateDefaultConfig())
	return err != componenterror.ErrDataTypeIsNotSupported
}

func isTracesExporter(f component.ExporterFactory) bool {
	_, err := f.CreateTracesExporter(context.Background(), component.ExporterCreateParams{}, f.CreateDefaultConfig())
	return err != componenterror.ErrDataTypeIsNotSupported
}
