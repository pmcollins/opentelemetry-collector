package receiver

import (
	"testing"

	"go.opentelemetry.io/collector/cmd/cfgwizard/inputreader"
	"go.opentelemetry.io/collector/cmd/cfgwizard/menu"
	"go.opentelemetry.io/collector/cmd/cfgwizard/printer"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
)

func TestReceiverController(t *testing.T) {
	m := map[configmodels.Type]component.ReceiverFactoryBase{}
	factory := fakeReceiverFactory{}
	m[factory.Type()] = factory
	out := printer.NewNop()
	in := inputreader.New(inputreader.NewFakeIOReader("1"), out)
	mc := menu.NewController(menu.NewView(out), in)
	rc := NewController(m, NewView(out), mc, in)
	rc.Config()
}

type fakeReceiverFactory struct {
}

func (f fakeReceiverFactory) Type() configmodels.Type {
	return "fakeReceiverFactory"
}

func (f fakeReceiverFactory) CreateDefaultConfig() configmodels.Receiver {
	return &configmodels.ReceiverSettings{}
}

func (f fakeReceiverFactory) CustomUnmarshaler() component.CustomUnmarshaler {
	return nil
}
