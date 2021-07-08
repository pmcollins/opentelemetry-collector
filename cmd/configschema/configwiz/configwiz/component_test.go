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
	"strings"
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

func buildTestCFGFields(name string, typ string, kind string, defaultStr string, doc string) (cfgField configschema.Field) {
	var fields []*configschema.Field
	cfgField2 := configschema.Field{
		Name: name + "1",
		Type: typ + "1",
		Kind: kind + "1",
		Default: defaultStr + "1",
		Doc: doc + "1",
		Fields:  nil,
	}
	fields = append(fields, &cfgField2)

	cfgField = configschema.Field{
		Name:    name,
		Type:    typ,
		Kind:    kind,
		Default: defaultStr,
		Doc:     doc,
		Fields: fields,
	}
	return cfgField
}

func runCompWizardSquash(io clio) configschema.Field {
	cfgField := buildTestCFGFields(
		"squashTest",
		"test",
		"[]string",
		"defaultStr",
		"testing CompWizard squash",
	)

	newField := buildTestCFGFields(
		"squash",
		"test",
		"[]string",
		"defaultStr",
		"testing Compwizard squash after squash",
	)
	var fields []*configschema.Field
	fields = append(fields, &newField)
	cfgField.Fields = fields
	componentWizard(io, 0, &cfgField)
	return cfgField
}

func runCompWizardStruct(io clio) configschema.Field {
	cfgField := buildTestCFGFields(
		"structTest",
		"test",
		"helper",
		"defaultStr",
		"testing CompWizard struct",
	)
	newField := buildTestCFGFields(
		"testCompWizStruct",
		"test",
		"struct",
		"defaultStr",
		"Testing comp Wizard struct case",
	)

	var fields []*configschema.Field
	fields = append(fields, &newField)
	cfgField.Fields = fields
	componentWizard(io, 0, &cfgField)
	return cfgField
}

func runCompWizardPtr(io clio) configschema.Field {
	cfgField := buildTestCFGFields(
		"ptrTest",
		"test",
		"helper",
		"defaultStr",
		"testing CompWizard Ptr",
		)

	newField := buildTestCFGFields(
		"testPtr",
		"test",
		"ptr",
		"defaultPtr",
		"testing compWizard ptr case",
		)
	var fields []*configschema.Field
	fields = append(fields, &newField)
	cfgField.Fields = fields
	componentWizard(io, 0, &cfgField)
	return cfgField
}

func runCompWizardHandleField(io clio) configschema.Field {
	cfgField := buildTestCFGFields(
		"testCompWizard",
		"test",
		"[]string",
		"default",
		"testing CompWizard handleField",
	)
	componentWizard(io, 0, &cfgField)
	return cfgField
}



func TestComponentWizard(t *testing.T) {
	//if field.name == squash
	writerSquash := fakeWriter{}
	ioSquash := clio{writerSquash.write, fakeReader{}.read}
	cfgSquash := runCompWizardSquash(ioSquash)
	currSquash := cfgSquash.Fields[0].Fields[0]
	expectedSquash := "Field: " + currSquash.Name + "\nType: " + currSquash.Type + "\n"
	expectedSquash += "Docs: " + currSquash.Doc + "\nDefault (enter to accept): defaultStr1\n> "
	//assert.Equal(t, expectedSquash, writerSquash.programOutput)

	//else if field.kind == "struct"
	writerStruct := fakeWriter{}
	ioStruct := clio{writerStruct.write, fakeReader{}.read}
	cfgStruct := runCompWizardStruct(ioStruct)
	tab := strings.Repeat(" ", 4)
	currStruct := cfgStruct.Fields[0].Fields[0]
	expectedStruct := string(cfgStruct.Fields[0].Name) + "\n" + tab + "Field: " + currStruct.Name + "\n"
	expectedStruct += tab + "Type: " + currStruct.Type + "\n" + tab + "Docs: " + currStruct.Doc + "\n"
	expectedStruct += tab + "Default (enter to accept): defaultStr1\n" + tab + "> "
	//assert.Equal(t, expectedStruct, writerStruct.programOutput)

	//else if field.kind == "ptr" FIXME
	writerPtr := fakeWriter{}
	ioPtr := clio{writerPtr.write, fakeReader{"n"}.read}
	//cfgPtr := runCompWizardPtr(ioPtr)
	runCompWizardPtr(ioPtr)
	assert.Equal(t, "", writerPtr.programOutput)
	//else
	writerHandle := fakeWriter{}
	ioHandle := clio{writerHandle.write, fakeReader{}.read}
	cfgHandle := runCompWizardHandleField(ioHandle)
	expectedHandle := "Field: " + cfgHandle.Fields[0].Name + "\nType: " + cfgHandle.Fields[0].Type + "\n"
	expectedHandle += "Docs: " + cfgHandle.Fields[0].Doc + "\nDefault (enter to accept): default1\n> "
	//assert.Equal(t, expectedHandle, writerHandle.programOutput)

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
		"default",
		"we are testing handleField",
		)
	handleField(p, io, &cfgField, out)
	expected := "Field: " + cfgField.Name + "\nType: " + cfgField.Type + "\nDocs: " + cfgField.Doc
	expected += "\nDefault (enter to accept): default" + "\n> "
	assert.Equal(t, expected, writer.programOutput)
}

func TestParseCSV(t *testing.T) {
	expected := []string{"a", "b", "c"}
	assert.Equal(t, expected, parseCSV("a,b,c"))
	assert.Equal(t, expected, parseCSV("a, b, c"))
	assert.Equal(t, expected, parseCSV(" a , b , c "))
	assert.Equal(t, []string{"a"}, parseCSV(" a "))
}
