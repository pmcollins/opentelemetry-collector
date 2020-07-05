package receiver

import (
	"go.opentelemetry.io/collector/cmd/cfgwizard/inputreader"
	"go.opentelemetry.io/collector/cmd/cfgwizard/menu"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
)

type Controller struct {
	view           View
	menuController menu.Controller
	receivers      map[configmodels.Type]component.ReceiverFactoryBase
	in             inputreader.InputReader
}

func NewController(
	receivers map[configmodels.Type]component.ReceiverFactoryBase,
	v View,
	mc menu.Controller,
	in inputreader.InputReader,
) Controller {
	return Controller{view: v, menuController: mc, receivers: receivers, in: in}
}

func (c Controller) Config() configmodels.Receiver {
	c.view.Menu()
	rcvrNames := menuItems(c.receivers)
	selected := c.menuController.Select(rcvrNames)
	receiverFactory := c.receivers[configmodels.Type(selected)]
	return receiverFactory.CreateDefaultConfig()
}

func menuItems(receivers map[configmodels.Type]component.ReceiverFactoryBase) (names []string) {
	for name, _ := range receivers {
		names = append(names, string(name))
	}
	return names
}
