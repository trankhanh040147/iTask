package uploadusecase

import (
	"paradise-booking/config"
	s3provider "paradise-booking/provider/s3"
)

type uploadUseCase struct {
	cfg        *config.Config
	s3Provider s3provider.S3Provider
}

func NewUploadUseCase(cfg *config.Config, s3Provider s3provider.S3Provider) *uploadUseCase {
	return &uploadUseCase{cfg, s3Provider}
}
