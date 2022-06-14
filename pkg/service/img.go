package service

import (
	"errors"
	"imgConverter/pkg/restmodel"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

type ImgService struct{}

func NewImgService() *ImgService {
	return &ImgService{}
}

func (s *ImgService) ImageProcessing(options restmodel.ImgOptions, name string) (string, error) {
	opt := bimg.Options{
		Quality: options.Quality,
		Width:   options.Width,
		Height:  options.Height,
	}

	buffer, err := bimg.Read(name)
	if err != nil {
		return "", err
	}

	imgMD, err := bimg.NewImage(buffer).Metadata()
	if err != nil {
		return "", err
	}

	newImage, err := convertExt(buffer, options.OutputFormat)
	if err != nil {
		return "", err
	}

	newImage, err = bimg.NewImage(newImage).Extract(options.CropTop, options.CropLeft, imgMD.Size.Width-options.CropRight-options.CropLeft, imgMD.Size.Height-options.CropBottom-options.CropTop)
	if err != nil {
		return "", err
	}

	newImage, err = bimg.NewImage(newImage).Process(opt)
	if err != nil {
		return "", err
	}

	filename := strings.TrimSuffix(name, filepath.Ext(name)) + "-res." + options.OutputFormat

	err = bimg.Write(filename, newImage)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *ImgService) ImageResize(options restmodel.ImgOptions, name string) (string, error) {
	buffer, err := bimg.Read(name)
	if err != nil {
		return "", err
	}

	imgMD, err := bimg.NewImage(buffer).Metadata()
	if err != nil {
		return "", err
	}

	imgWidth := imgMD.Size.Width

	w := 100
	if imgWidth < 1000 {
		w = 10
	}

	for len(buffer) > options.MaxSize {
		imgWidth -= w

		buffer, err = bimg.NewImage(buffer).Resize(imgWidth, 0)
		if err != nil {
			return "", err
		}
	}

	filename := strings.TrimSuffix(name, filepath.Ext(name)) + "-res." + options.OutputFormat

	err = bimg.Write(filename, buffer)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func convertExt(buf []byte, ext string) ([]byte, error) {
	switch strings.ToLower(ext) {
	case "png":
		return bimg.NewImage(buf).Convert(bimg.PNG)
	case "svg":
		return bimg.NewImage(buf).Convert(bimg.SVG)
	case "jpg", "jpeg":
		return bimg.NewImage(buf).Convert(bimg.JPEG)
	}

	return nil, errors.New("ext error")
}
