package menu

import (
	"sort"

	"go.opentelemetry.io/collector/cmd/cfgwizard/inputreader"
)

type Controller struct {
	view View
	in   inputreader.InputReader
}

func NewController(v View, in inputreader.InputReader) Controller {
	return Controller{view: v, in: in}
}

func (c Controller) Select(names []string) string {
	sort.Strings(names)
	c.view.Display(names)
	for {
		i := c.in.ReadInt()
		if i != nil {
			if *i > len(names) {
				continue
			}
			return names[*i-1]
		}
	}
}
