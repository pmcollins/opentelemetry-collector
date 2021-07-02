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
	"strings"
)

type indentingPrinter struct {
	level int
	write func(s string, newLine bool)
}

func (p *indentingPrinter) ioWrite(s string, newLine bool) {
	if newLine {
		p.println(s)
	} else {
		p.print(s)
	}
}

func (p indentingPrinter) println(s string) {
	p.doPrint(s, "%s%s\n")
}

func (p indentingPrinter) print(s string) {
	p.doPrint(s, "%s%s")
}

func (p indentingPrinter) doPrint(s string, frmt string) {
	const tabSize = 4
	indent := p.level * tabSize
	fmt.Printf(frmt, strings.Repeat(" ", indent), s)
}
