package main

import (
	"os"

	"go.opentelemetry.io/collector/cmd/cfgwizard/configwiz"
	"go.opentelemetry.io/collector/cmd/cfgwizard/inputreader"
	"go.opentelemetry.io/collector/cmd/cfgwizard/menu"
	"go.opentelemetry.io/collector/cmd/cfgwizard/printer"
	"go.opentelemetry.io/collector/cmd/cfgwizard/receiver"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/service/defaultcomponents"
)

func main() {
	var factories, _ = defaultcomponents.Components()
	app(factories)
}

func app(factories config.Factories) {
	outpt := printer.New()
	inpt := inputreader.New(os.Stdin)
	menuController := menu.NewController(menu.NewView(outpt), inpt)
	receiverController := receiver.NewController(
		factories.Receivers,
		receiver.NewView(outpt),
		menuController,
		inpt,
	)
	rcfg := receiverController.Config()
	cfgController := configwiz.NewController(inpt, configwiz.NewView(outpt))
	cfgController.Populate(rcfg)
}
