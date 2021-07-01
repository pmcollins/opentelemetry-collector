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
	"io/ioutil"
	"os"
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
	ip            indentingPrinter
}

func (w *fakeWriter) write(s string) {
	w.programOutput = s
}


func TestHandleField(t *testing.T) {
	writer := fakeWriter{ip: indentingPrinter{level: 0}}
	reader := fakeReader{}
	io := clio{writer.ip, reader.read}
	out := map[string]interface{}{}
	cfgField := configschema.Field{
		Name:    "testHandleField",
		Type:    "test",
		Kind:    "string",
		Default: nil,
		Doc:     "We are testing HandleField",
		Fields:  nil,
	}
	// piping stdout to program
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	handleField(io, &cfgField, out)

	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expectedField := "Field: " + cfgField.Name + "\nType: " + cfgField.Type + "\nDocs: " + cfgField.Doc + "\n> "
	assert.Equal(t, expectedField, string(output))

}

func TestParseCSV(t *testing.T) {
	expected := []string{"a", "b", "c"}
	assert.Equal(t, expected, parseCSV("a,b,c"))
	assert.Equal(t, expected, parseCSV("a, b, c"))
	assert.Equal(t, expected, parseCSV(" a , b , c "))
	assert.Equal(t, []string{"a"}, parseCSV(" a "))
}

