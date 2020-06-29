package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"

	"go.opentelemetry.io/collector/receiver/otlpreceiver"
)

func TestMapstructure(t *testing.T) {
	cfg := otlpreceiver.Config{}
	m := map[string]interface{}{}
	err := mapstructure.Decode(cfg, &m)
	if err != nil {
		println(err.Error())
		return
	}
	yml, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	println(string(yml))
}

func TestIndenter(t *testing.T) {
	indented := indent(`
a
b
c
d
`, 4)
	println(indented)
}

func TestMTag(t *testing.T) {
	s, meta := mTag(`mapstructure:"ca_file"`)
	fmt.Printf("[%s] [%s]\n", s, meta)
	s, meta = mTag(`mapstructure:"ca_file,omitempty"`)
	fmt.Printf("[%s] [%s]\n", s, meta)
}

func TestDur(t *testing.T) {
	dur, _ := time.ParseDuration("1m22s")
	fmt.Printf("%v", dur)
}
