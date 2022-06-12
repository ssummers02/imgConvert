package restmodel

type ImgOptions struct {
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	Quality      int    `json:"quality,omitempty"`
	OutputFormat string `json:"output_format,omitempty"`
}
