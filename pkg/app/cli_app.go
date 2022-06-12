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
		"set width",
	)
	height = flag.Int("height", 0,
		"set height")
	path = flag.String("path", "",
		"set path")
	quality = flag.Int("quality", 100,
		"set quality")
	format = flag.String("format", "",
		"set output format")
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
