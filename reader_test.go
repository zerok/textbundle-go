package textbundle_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zerok/textbundle-go"
)

func TestReader(t *testing.T) {
	r, err := textbundle.OpenReader("testdata/test.textpack")
	require.NoError(t, err)
	require.NotNil(t, r)
	require.Equal(t, 2, r.Metadata.Version)
	require.True(t, bytes.HasPrefix(r.Text, []byte("# Getting")))
	require.Len(t, r.Assets, 1)
}
