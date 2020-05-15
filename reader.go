package textbundle

import (
	"archive/zip"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

const fnamePrefix = "Content.textbundle/"

func fn(name string) string {
	return fnamePrefix + name
}

type Reader struct {
	Metadata Metadata
	Text     []byte
	Assets   []Asset
}

type Asset struct {
	Name string
	File *zip.File
}

func NewReader(r io.ReaderAt, size int64) (*Reader, error) {
	zip.NewReader(r, size)
	return nil, nil
}

func OpenReader(path string) (*Reader, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	bundle := &Reader{}
	md := Metadata{}
	for _, f := range r.File {
		if f.FileHeader.Name == fn("info.json") {
			fp, err := f.Open()
			if err != nil {
				return nil, err
			}
			if err := json.NewDecoder(fp).Decode(&md); err != nil {
				fp.Close()
				return nil, err
			}
			fp.Close()
		} else if strings.HasPrefix(f.FileHeader.Name, fnamePrefix+"text.") {
			fp, err := f.Open()
			if err != nil {
				return nil, err
			}
			data, err := ioutil.ReadAll(fp)
			if err != nil {
				fp.Close()
				return nil, err
			}
			bundle.Text = data
		} else if strings.HasPrefix(f.FileHeader.Name, fnamePrefix+"assets/") && f.FileHeader.Name != fn("assets/") {
			a := Asset{
				Name: strings.TrimPrefix(f.Name, fnamePrefix+"assets/"),
				File: f,
			}
			bundle.Assets = append(bundle.Assets, a)
		}
	}
	bundle.Metadata = md
	return bundle, nil
}
