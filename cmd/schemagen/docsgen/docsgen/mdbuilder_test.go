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

package docsgen

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/cmd/schemagen/configschema"
)

func TestBuildFields(t *testing.T) {
	jsonBytes, err := ioutil.ReadFile("testdata/opencensusreceiver.json")
	field := configschema.Field{}
	err = json.Unmarshal(jsonBytes, &field)
	require.NoError(t, err)

	tmpl, err := tableTemplate()
	require.NoError(t, err)

	buf := &bytes.Buffer{}
	err = renderTree(tmpl, &field, buf)
	require.NoError(t, err)

	require.NotNil(t, buf.Bytes())
}
