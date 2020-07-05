package configwiz

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/cmd/cfgwizard/printer"
)

type View struct {
	pr printer.Printer
}

func NewView(p printer.Printer) View {
	return View{pr: p}
}

func (v View) StringLabel(label string, factoryDefault string, lvl int) {
	var s string
	if factoryDefault == "" {
		s = label + ": "
	} else {
		s = label + ": (default: [" + factoryDefault + "]"
	}
	v.pr.Indent(s, lvl)
}

func (v View) BoolLabel(label string, lvl int) {
	v.pr.Indent(label + ` ("true"/"false")`, lvl)
}

func (v View) IntLabel(label string, factoryDefault int64, lvl int) {
	var s string
	if factoryDefault == 0 {
		s = label + " (int): "
	} else {
		s = fmt.Sprintf("(int) (default: [%d])", factoryDefault)
	}
	v.pr.Indent(s, lvl)
}

func (v View) DurationLabel(label string, factoryDefault int64, lvl int) {
	var s string
	if factoryDefault == 0 {
		s = label + ` (duration e.g."1ms" "1s" "1m"): `
	} else {
		dur := time.Duration(factoryDefault)
		s = fmt.Sprintf("%s (duration e.g.\"1ms\" \"1s\" \"1m\") (default %v): ", label, dur)
	}
	v.pr.Indent(s, lvl)
}

func (v View) Label(label string, lvl int) {
	v.pr.Indent(label, lvl)
}

func (v View) Skip() {
	v.pr.Print(" skip? [y]/n: ")
}

func (v View) Println() {
	println()
}
