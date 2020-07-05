package menu

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/cmd/cfgwizard/inputreader"
	"go.opentelemetry.io/collector/cmd/cfgwizard/printer"
)

func TestController(t *testing.T) {
	p := printer.NewNop()
	v := NewView(p)
	in := inputreader.New(inputreader.NewFakeIOReader("1", "2", "3"), p)
	menuController := NewController(v, in)
	fruit := []string{"apple", "banana", "cherry"}
	selected := menuController.Select(fruit)
	require.Equal(t, "apple", selected)
	selected = menuController.Select(fruit)
	require.Equal(t, "banana", selected)
	selected = menuController.Select(fruit)
	require.Equal(t, "cherry", selected)
}
