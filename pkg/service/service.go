package service

type Img interface {
	Create()
}

type Service struct {
	Img Img
}

func NewService() *Service {
	return &Service{
		Img: NewImgService(),
	}
}
