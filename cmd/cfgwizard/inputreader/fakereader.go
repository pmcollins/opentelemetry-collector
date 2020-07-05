// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inputreader

import "io"

type FakeReader struct {
	inputNum int
	inputs   []string
}

var _ io.Reader = (*FakeReader)(nil)

func NewFakeIOReader(inputs ...string) *FakeReader {
	return &FakeReader{inputs: inputs}
}

func (r *FakeReader) Read(p []byte) (n int, err error) {
	i := copy(p, r.inputs[r.inputNum])
	r.inputNum++
	p[i] = '\n'
	return i + 1, nil
}
