package service

import (
	"errors"
	"imgConverter/pkg/restmodel"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

const ImgDirectory = "uploads"
const ImgResult = "result"

type ImgService struct {
}

func NewImgService() *ImgService {
	return &ImgService{}
}

func (s *ImgService) ImageProcessing(options restmodel.ImgOptions, name string) (string, error) {
	opt := bimg.Options{
		Quality: options.Quality,
		Width:   options.Width,
		Height:  options.Height,
	}

	buffer, err := bimg.Read(filepath.Join(ImgDirectory, name))
	if err != nil {
		return "", err
	}

	newImage, err := convertExt(buffer, options.OutputFormat)
	if err != nil {
		return "", err
	}

	newImage, err = bimg.NewImage(buffer).Process(opt)
	if err != nil {
		return "", err
	}

	filename := strings.TrimSuffix(name, filepath.Ext(name)) + "." + options.OutputFormat

	path := filepath.Join(ImgResult, filename)
	bimg.Write(path, newImage)

	return path, nil
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
