package textbundle

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const fnamePrefix = "Content.textbundle/"

func fn(name string) string {
	return fnamePrefix + name
}

// Reader allows access the seperate Assets of a TextBundle, its content, and
// the associated metadata.
type Reader struct {
	Metadata Metadata
	Text     []byte
	Assets   []Asset
}

// Asset is a light wrapper around zip.File that trims the filename.
type Asset struct {
	Name string
	File *zip.File
}

// OpenReader creates a new Reader from the given file path.
func OpenReader(path string) (*Reader, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %w", path, err)
	}
	bundle := &Reader{}
	md := Metadata{}
	for _, f := range r.File {
		if f.FileHeader.Name == fn("info.json") {
			fp, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("could not open info.json for reading: %w", err)
			}
			if err := json.NewDecoder(fp).Decode(&md); err != nil {
				fp.Close()
				return nil, fmt.Errorf("could not decode info.json: %w", err)
			}
			fp.Close()
		} else if strings.HasPrefix(f.FileHeader.Name, fnamePrefix+"text.") {
			fp, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("could not open text item for reading: %w", err)
			}
			data, err := ioutil.ReadAll(fp)
			if err != nil {
				fp.Close()
				return nil, fmt.Errorf("could not read text item: %w", err)
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
