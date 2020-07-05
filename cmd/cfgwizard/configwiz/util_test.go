package configwiz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDestructure(t *testing.T) {
	a, b := splitTag("mapstructure:foo,bar")
	require.Equal(t, "foo", a)
	require.Equal(t, "bar", b)
}
