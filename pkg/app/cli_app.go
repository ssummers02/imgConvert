package app

import (
	"flag"
	"fmt"
	"imgConverter/pkg/restmodel"
	"imgConverter/pkg/service"
	"path/filepath"
)

var (
	width = flag.Int("width", 0,
		"set width")
	height = flag.Int("height", 0,
		"set height")
	path = flag.String("path", "",
		"set path")
	quality = flag.Int("quality", 100,
		"set quality")
	format = flag.String("format", "",
		"set output format")
	top = flag.Int("crop_top", 0,
		"set crop_top")
	left = flag.Int("crop_left", 0,
		"set crop_left")
	bottom = flag.Int("crop_bottom", 0,
		"set crop_bottom")
	right = flag.Int("crop_right", 0,
		"set crop_right")
	maxSize = flag.Int("maxSize", 0,
		"set maxSize")
)

//nolint:forbidigo
func RunCli() {
	s := service.NewService()

	if *path == "" {
		fmt.Println("path is empty")

		return
	}

	var opt restmodel.ImgOptions
	if *width != 0 {
		opt.Width = *width
	}

	if *height != 0 {
		opt.Height = *height
	}

	if *quality != 100 {
		opt.Quality = *quality
	}

	if *top != 0 {
		opt.CropTop = *top
	}

	if *left != 0 {
		opt.CropLeft = *left
	}

	if *bottom != 0 {
		opt.CropBottom = *bottom
	}

	if *right != 0 {
		opt.CropRight = *right
	}

	if *maxSize != 0 {
		opt.MaxSize = *maxSize
	}

	if *format != "" {
		opt.OutputFormat = *format
	} else {
		opt.OutputFormat = filepath.Ext(*path)[1:]
	}

	converting, err := s.Img.ImageProcessing(opt, *path)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("converting:", converting)
}
