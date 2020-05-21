package textbundle

import (
	"archive/zip"
	"encoding/json"
	"io"
)

type Writer struct {
	zipw     *zip.Writer
	Metadata Metadata
}

// NewWriter creates a new writer which should be persisted to the out.
func NewWriter(out io.Writer) *Writer {
	zipw := zip.NewWriter(out)
	return &Writer{
		zipw: zipw,
		Metadata: Metadata{
			Version: 2,
			Type:    "net.daringfireball.markdown",
		},
	}
}

func (w *Writer) writeInfo() error {
	out, err := w.zipw.Create("Content.textbundle/info.json")
	if err != nil {
		return err
	}
	return json.NewEncoder(out).Encode(&w.Metadata)
}

// Flush flushes the underlying writer.
func (w *Writer) Flush() error {
	return w.zipw.Flush()
}

// Close finalizes the bundles metadata and closes the underlying writer.
func (w *Writer) Close() error {
	if err := w.writeInfo(); err != nil {
		return err
	}
	return w.zipw.Close()
}

// CreateAsset creates a new asset file with the given name.
func (w *Writer) CreateAsset(name string) (io.Writer, error) {
	return w.zipw.Create("Content.textbundle/assets/" + name)
}

// SetText writes Content.textbundle/text.$typ.
func (w *Writer) SetText(typ string, content string) error {
	fp, err := w.zipw.Create("Content.textbundle/text." + typ)
	if err != nil {
		return err
	}
	_, err = fp.Write([]byte(content))
	return err
}
