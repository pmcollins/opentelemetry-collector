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
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/cmd/configschema/configschema"
	"testing"
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

func TestHandleField(t *testing.T) {
	writer := fakeWriter{}
	reader := fakeReader{}
	io := clio{writer.write, reader.read}
	p := indentingPrinter2{level: 0}
	p.write = io.write
	out := map[string]interface{}{}
	cfgField := configschema.Field{
		Name:    "testHandleField",
		Type:    "test",
		Kind:    "string",
		Default: nil,
		Doc:     "We are testing HandleField",
		Fields:  nil,
	}
	handleField(p, io, &cfgField, out)
	expected := "Field: " + cfgField.Name + "\nType: " + cfgField.Type + "\nDocs: " + cfgField.Doc + "\n> "
	assert.Equal(t, expected, writer.programOutput)
}

func TestParseCSV(t *testing.T) {
	expected := []string{"a", "b", "c"}
	assert.Equal(t, expected, parseCSV("a,b,c"))
	assert.Equal(t, expected, parseCSV("a, b, c"))
	assert.Equal(t, expected, parseCSV(" a , b , c "))
	assert.Equal(t, []string{"a"}, parseCSV(" a "))
}
