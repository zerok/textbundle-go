package textbundle_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zerok/textbundle-go"
)

func TestWriter(t *testing.T) {
	fp, err := ioutil.TempFile(os.TempDir(), "bundle")
	require.NoError(t, err)
	defer os.RemoveAll(fp.Name())
	w := textbundle.NewWriter(fp)
	require.NoError(t, w.SetText("md", "Hello world"))
	_, err = w.CreateAsset("test.dat")
	require.NoError(t, err)
	require.NoError(t, w.Close())
	require.NoError(t, fp.Close())

	r, err := textbundle.OpenReader(fp.Name())
	require.NotEmpty(t, r.Assets)
	require.NoError(t, err)
	require.NotNil(t, r)
	require.Equal(t, "Hello world", string(r.Text))
}
