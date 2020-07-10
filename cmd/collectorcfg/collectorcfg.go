package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/service/defaultcomponents"
)

var reader = bufio.NewReader(os.Stdin)
var factories, _ = defaultcomponents.Components()

func main() {
	bigYaml := ""
	receiverYaml, receiverName := receiverSection()
	bigYaml += receiverYaml

	procYaml, procName := processorSection()
	bigYaml += procYaml

	exporterYaml, exporterName := exporterSection()
	bigYaml += exporterYaml

	serviceYaml := serviceYaml(receiverName, procName, exporterName)
	bigYaml += serviceYaml

	const filename = "cfg.yaml"
	println("")
	println("~~~~~~~~~~~~" + filename + "~~~~~~~~~~~~")
	println(bigYaml)

	_ = ioutil.WriteFile(filename, []byte(bigYaml), 0644)
}

func receiverSection() (string, string) {
	println("Choose a receiver")
	var names []string
	for name, _ := range factories.Receivers {
		names = append(names, string(name))
	}
	key := menu(names)
	rcvrFact := factories.Receivers[configmodels.Type(key)]
	cfg := rcvrFact.CreateDefaultConfig()
	rcvrYaml := populateStruct(cfg)
	return fmt.Sprintf(
		`receivers:
  %s:
%s`, key, indent(rcvrYaml, 4)), key
}

func processorSection() (string, string) {
	println("Choose a processor ([enter] to skip)")
	var names []string
	for name, _ := range factories.Processors {
		names = append(names, string(name))
	}
	key := menu(names)
	if key == "" {
		return "", ""
	}
	procFactory := factories.Processors[configmodels.Type(key)]
	cfg := procFactory.CreateDefaultConfig()
	yml := populateStruct(cfg)
	return fmt.Sprintf(
		`processors:
  %s:
%s`, key, indent(yml, 4)), key
}

func exporterSection() (string, string) {
	println("Choose an exporter")
	var names []string
	for name, _ := range factories.Exporters {
		names = append(names, string(name))
	}
	key := menu(names)
	exporterFactory := factories.Exporters[configmodels.Type(key)]
	cfg := exporterFactory.CreateDefaultConfig()
	yml := populateStruct(cfg)
	return fmt.Sprintf(
		`exporters:
  %s:
%s`, key, indent(yml, 4)), key
}

func menu(names []string) string {
	sort.Strings(names)
	for i, name := range names {
		fmt.Printf("%d: %v\n", i+1, name)
	}
	print("> ")
	numStr := readline("")
	if numStr == "" {
		return ""
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	key := names[num-1]
	println("Selected " + key)
	return key
}

type strMap map[string]interface{}

func populateStruct(in interface{}) string {
	m := strMap{}
	populateMap(reflect.ValueOf(in), m, 0)
	return mapToYaml(m)
}

func populateMap(input reflect.Value, m strMap, lvl int) {
	prntr := printer{lvl}
	if input.Kind() == reflect.Ptr {
		input = input.Elem()
	}
	inputType := input.Type()
	for i := 0; i < input.NumField(); i++ {
		structField := inputType.Field(i)
		tag, meta := mTag(structField.Tag)
		if tag == "-" {
			continue
		}
		label := tag
		if label == "" {
			label = structField.Name
		}
		mapKey := label // we don't want (optional) in the map key
		if meta == "omitempty" {
			label += " (optional)"
		}
		fldVal := input.Field(i)
		switch fldVal.Kind() {
		case reflect.Map:
			prntr.println(label + ": map type not implemented yet")
			continue
		case reflect.Slice:
			prntr.println(label + ": slice type not implemented yet")
			continue
		case reflect.Struct:
			prntr.print(label)

			if meta == "omitempty" {
				print(" skip? [y]/n: ")
				skip := readline("y")
				if skip == "y" {
					continue
				}
			} else {
				println()
			}

			if meta == "squash" {
				populateMap(fldVal, m, lvl+1)
			} else {
				subMap := strMap{}
				m[mapKey] = subMap
				populateMap(fldVal, subMap, lvl+1)
			}
		case reflect.Ptr:
			prntr.print(label)

			if meta == "omitempty" {
				print(" skip? [y]/n: ")
				skip := readline("y")
				if skip == "y" {
					continue
				}
			} else {
				println()
			}

			// todo handle squash?

			newVal := reflect.New(fldVal.Type().Elem())

			subMap := strMap{}
			m[mapKey] = subMap

			populateMap(newVal, subMap, lvl+1)
		case reflect.String:
			factoryDefault := fldVal.String()
			if factoryDefault == "" {
				prntr.print(label + ": ")
			} else {
				prntr.print(label + " (default: [" + factoryDefault + "]): ")
			}
			str := readline(factoryDefault)
			m[mapKey] = str
		case reflect.Bool:
			prntr.print(label + ` ("true"/"false"): `)
			str := readline("")
			if str == "" {
				continue
			}
			m[mapKey] = str
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			factoryDefault := fldVal.Int()

			if structField.Type.String() == "time.Duration" {
				defaultStr := ""
				if factoryDefault == 0 {
					prntr.print(label + ` (duration e.g."1ms" "1s" "1m"): `)
				} else {
					dur := time.Duration(factoryDefault)
					defaultStr = fmt.Sprintf("%v", dur)
					prntr.print(label + ` (duration e.g."1ms" "1s" "1m") (default ` + defaultStr + `): `)
				}
				str := readline(defaultStr)
				if str == "" {
					continue
				}
				_, err := time.ParseDuration(str)
				if err != nil {
					println(err.Error())
					continue
				}
				m[mapKey] = str
				continue
			}

			dfltStr := ""
			if factoryDefault == 0 {
				prntr.print(label + " (int): ")
			} else {
				prntr.print(label + " (int) (default: [" + fmt.Sprintf("%d", factoryDefault) + "]): ")
				dfltStr = strconv.Itoa(int(factoryDefault))
			}

			str := readline(dfltStr)
			num, err := strconv.Atoi(str)
			if err != nil {
				println(err.Error())
				continue
			}
			m[mapKey] = num
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			factoryDefault := fldVal.Uint()
			dfltStr := ""
			if factoryDefault == 0 {
				prntr.print(label + " (uint): ")
			} else {
				prntr.print(label + " (uint) (default: [" + fmt.Sprintf("%d", factoryDefault) + "]): ")
				dfltStr = strconv.Itoa(int(factoryDefault))
			}
			str := readline(dfltStr)
			if str == "" {
				continue
			}
			num, err := strconv.Atoi(str)
			if err != nil {
				println(err.Error())
				continue
			}
			m[mapKey] = num
		}
	}
}

func mapToYaml(m strMap) string {
	yml, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(yml)
}

func mTag(s reflect.StructTag) (string, string) {
	trimmed := strings.Trim(strings.TrimPrefix(string(s), "mapstructure:"), `"`)
	split := strings.Split(trimmed, ",")
	meta := ""
	if len(split) == 2 {
		meta = strings.TrimSpace(split[1])
	}
	return split[0], meta
}

func readline(dflt string) string {
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str = str[:len(str)-1]
	if str == "" {
		return dflt
	}
	return str
}

type printer struct {
	level int
}

func (p printer) println(s string) {
	p.doPrint(s, "%s%s\n")
}

func (p printer) print(s string) {
	p.doPrint(s, "%s%s")
}

func (p printer) doPrint(s string, frmt string) {
	const tabSize = 4
	indent := p.level * tabSize
	fmt.Printf(frmt, strings.Repeat(" ", indent), s)
}

func indent(s string, lvl int) string {
	out := ""
	const lf = "\n"
	lines := strings.Split(s, lf)
	for _, line := range lines {
		spaces := strings.Repeat(" ", lvl)
		out += spaces + line + lf
	}
	return out
}

func serviceYaml(receiverName string, procName string, exporterName string) string {
	m := strMap{}
	metricsMap := strMap{
		"receivers": []string{receiverName},
		"exporters": []string{exporterName},
	}
	if procName != "" {
		metricsMap["processors"] = []string{procName}
	}
	m["service"] = strMap{
		"pipelines": strMap{
			"metrics": metricsMap,
		},
	}
	return mapToYaml(m)
}
