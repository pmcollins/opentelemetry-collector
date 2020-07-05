package menu

import (
	"fmt"

	"go.opentelemetry.io/collector/cmd/cfgwizard/printer"
)

type View struct {
	printer printer.Printer
}

func NewView(p printer.Printer) View {
	return View{printer: p}
}

func (v View) Display(names []string) {
	for i, name := range names {
		v.printer.Println(fmt.Sprintf("%d: %v", i+1, name))
	}
}
