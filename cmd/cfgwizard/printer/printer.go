package printer

import "strings"

type Printer struct {
	f func(string)
}

func New() Printer {
	return Printer{
		f: func(s string) { print(s) },
	}
}

func NewTest(lc *LineCapturer) *Printer {
	return &Printer{f: lc.capture}
}

func NewNop() Printer {
	return Printer{f: func(s string) {}}
}

func (p Printer) Indent(s string, lvl int) {
	const tabSize = 2
	pad := strings.Repeat(" ", lvl*tabSize)
	p.Print(pad + s)
}

func (p Printer) Println(s string) {
	p.Print(s + "\n")
}

func (p Printer) Print(s string) {
	p.f(s)
}
