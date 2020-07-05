package printer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndenter(t *testing.T) {
	lc := NewLineCapturer()
	p := NewTest(lc)
	p.Print("foo")
	require.Equal(t, "foo", lc.Lines[0])
	i := p.NewIndenter(1)
	i.Print("bar")
	require.Equal(t, "  bar", lc.Lines[1])
}
