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
	"testing"

	"go.opentelemetry.io/collector/cmd/configschema/configschema"

	"github.com/stretchr/testify/assert"
)

type fakeReader struct {
	userInput string
}

func (r fakeReader) read(defaultVal string) string {
	return r.userInput
}

type fakeWriter struct {
	programOutput string
}

func (w *fakeWriter) write(s string) {
	w.programOutput += s
}

func buildTestCFGFields(name string, typ string, kind string, defaultStr string, doc string) configschema.Field {
	var fields []*configschema.Field
	cfgField2 := configschema.Field{
		Name: name + "1",
		Type: typ + "1",
		Kind: kind + "1",
		Default: defaultStr + "1",
		Doc: doc + "1",
	}
	fields = append(fields, &cfgField2)
	cfgField := configschema.Field{
		Name:    name,
		Type:    typ,
		Kind:    kind,
		Default: defaultStr,
		Doc:     doc,
		Fields: fields,
	}
	return cfgField
}

func TestHandleField(t *testing.T) {
	writer := fakeWriter{}
	reader := fakeReader{}
	io := clio{writer.write, reader.read}
	p := io.newIndentingPrinter(0)
	out := map[string]interface{}{}
	cfgField := buildTestCFGFields(
		"testHandleField",
		"test",
		"[]string",
		"defaultStr1",
		"we are testing handleField",
	)
	handleField(io, p, &cfgField, out)
	expected := fmt.Sprintf("Field: %s\n", cfgField.Name)
	expected += fmt.Sprintf("Type: %s\n", cfgField.Type)
	expected += fmt.Sprintf("Docs: %s\n", cfgField.Doc)
	expected += "Default (enter to accept): defaultStr1\n"
	expected += "> "
	assert.Equal(t, expected, writer.programOutput)
}

func TestParseCSV(t *testing.T) {
	expected := []string{"a", "b", "c"}
	assert.Equal(t, expected, parseCSV("a,b,c"))
	assert.Equal(t, expected, parseCSV("a, b, c"))
	assert.Equal(t, expected, parseCSV(" a , b , c "))
	assert.Equal(t, []string{"a"}, parseCSV(" a "))
}
