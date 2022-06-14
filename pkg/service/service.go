package service

import "imgConverter/pkg/restmodel"

type Img interface {
	ImageProcessing(options restmodel.ImgOptions, name string) (string, error)
	ImageResize(options restmodel.ImgOptions, name string) (string, error)
}

type Service struct {
	Img Img
}

func NewService() *Service {
	return &Service{
		Img: NewImgService(),
	}
}
