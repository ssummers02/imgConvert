package restmodel

type ImgOptions struct {
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	Quality      int    `json:"quality,omitempty"`
	OutputFormat string `json:"output_format,omitempty"`
	CropTop      int    `json:"crop_top,omitempty"`
	CropBottom   int    `json:"crop_bottom,omitempty"`
	CropLeft     int    `json:"crop_left,omitempty"`
	CropRight    int    `json:"crop_right,omitempty"`
}
