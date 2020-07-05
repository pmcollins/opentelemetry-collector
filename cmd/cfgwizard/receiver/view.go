package receiver

import "go.opentelemetry.io/collector/cmd/cfgwizard/printer"

type View struct {
	p printer.Printer
}

func NewView(p printer.Printer) View {
	return View{p}
}

func (v View) Menu() {
	v.p.Println("Choose a receiver")
}
