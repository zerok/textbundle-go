package textbundle

type Metadata struct {
	Version           int    `json:"version"`
	Type              string `json:"type,omitempty"`
	Transient         bool   `json:"transient,omitempty"`
	CreatorURL        string `json:"creatorURL,omitempty"`
	CreatorIdentifier string `json:"creatorIdentifier,omitempty"`
	SourceURL         string `json:"sourceURL,omitempty"`
}
