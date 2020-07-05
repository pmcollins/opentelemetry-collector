package configwiz

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"

	"go.opentelemetry.io/collector/cmd/cfgwizard/inputreader"
)

const (
	squash       = "squash"
	omitempty    = "omitempty"
	y            = "y"
	timeDuration = "time.Duration"
)

type strMap map[string]interface{}

type Controller struct {
	inpt inputreader.InputReader
	vw   View
}

func NewController(in inputreader.InputReader, v View) Controller {
	return Controller{inpt: in, vw: v}
}

func (c Controller) Populate(cfg interface{}) {
	m := strMap{}
	c.populateMap(reflect.ValueOf(cfg), m, 0)
}

func (c Controller) populateMap(input reflect.Value, m strMap, lvl int) {
	if input.Kind() == reflect.Ptr {
		input = input.Elem()
	}
	inputType := input.Type()
	for i := 0; i < input.NumField(); i++ {
		structField := inputType.Field(i)
		tag, meta := splitTag(structField.Tag)
		if tag == "-" {
			continue
		}
		label := tag
		if label == "" {
			label = structField.Name
		}
		if meta == "omitempty" {
			label += " (optional)"
		}
		fldVal := input.Field(i)
		switch fldVal.Kind() {
		case reflect.Map:
		case reflect.Slice:
		case reflect.Struct:
			c.handleStruct(m, fldVal, label, meta, lvl)
		case reflect.Ptr:
			c.handlePtr(m, fldVal, label, meta, lvl)
		case reflect.String:
			c.handleString(m, fldVal, label, lvl)
		case reflect.Bool:
			c.handleBool(m, label, lvl)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			c.handleSignedInt(m, structField, fldVal, label, lvl)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			c.handleUint(m, fldVal, label, lvl)
		}
	}
}

func (c Controller) handleStruct(m strMap, val reflect.Value, label string, meta string, lvl int) {
	if c.omitEmpty(label, lvl, meta) {
		return
	}
	if meta == squash {
		c.populateMap(val, m, lvl+1)
	} else {
		subMap := strMap{}
		m[label] = subMap
		c.populateMap(val, subMap, lvl+1)
	}
}

func (c Controller) handlePtr(m strMap, val reflect.Value, label string, meta string, lvl int) {
	if c.omitEmpty(label, lvl, meta) {
		return
	}
	// todo handle squash?
	newVal := reflect.New(val.Type().Elem())
	subMap := strMap{}
	m[label] = subMap
	c.populateMap(newVal, subMap, lvl+1)
}

func (c Controller) omitEmpty(label string, lvl int, meta string) bool {
	c.vw.Label(label, lvl)
	if meta == omitempty {
		c.vw.Skip()
		skip := c.inpt.ReadStringDefault(y)
		if skip == y {
			return true
		}
	} else {
		c.vw.Println()
	}
	return false
}

func (c Controller) handleString(m strMap, val reflect.Value, label string, lvl int) {
	factoryDefault := val.String()
	c.vw.StringLabel(label, factoryDefault, lvl)
	str := c.inpt.ReadStringDefault(factoryDefault)
	m[label] = str
}

func (c Controller) handleBool(m strMap, label string, lvl int) {
	// todo handle default
	c.vw.BoolLabel(label, lvl)
	str := c.inpt.ReadString()
	if str == "" {
		return
	}
	m[label] = str
}

func (c Controller) handleSignedInt(m strMap, field reflect.StructField, val reflect.Value, label string, lvl int) {
	// fixme?
	if field.Type.String() == timeDuration {
		c.handleDuration(m, val, label, lvl)
		return
	}
	c.handleInt(m, label, val.Int(), lvl)
}

func (c Controller) handleUint(m strMap, val reflect.Value, label string, lvl int) {
	c.handleInt(m, label, int64(val.Uint()), lvl)
}

func (c Controller) handleInt(m strMap, label string, factoryDefault int64, lvl int) {
	c.vw.IntLabel(label, factoryDefault, lvl)
	dfltStr := ""
	if factoryDefault != 0 {
		dfltStr = strconv.Itoa(int(factoryDefault))
	}
	str := c.inpt.ReadStringDefault(dfltStr)
	_, err := strconv.Atoi(str)
	if err != nil {
		// fixme
		println(err.Error())
		return
	}
	m[label] = str
}

func (c Controller) handleDuration(m strMap, val reflect.Value, label string, lvl int) {
	factoryDefault := val.Int()
	defaultStr := ""
	if factoryDefault != 0 {
		dur := time.Duration(factoryDefault)
		defaultStr = fmt.Sprintf("%v", dur)
	}
	c.vw.DurationLabel(label, factoryDefault, lvl)
	str := c.inpt.ReadStringDefault(defaultStr)
	if str == "" {
		return
	}
	_, err := time.ParseDuration(str)
	// fixme
	if err != nil {
		println(err.Error())
		return
	}
	m[label] = str
}

func mapToYaml(m strMap) string {
	yml, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(yml)
}
