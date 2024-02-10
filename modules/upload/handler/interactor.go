package uploadhandler

import (
	"context"
	"mime/multipart"
	"paradise-booking/common"
	"paradise-booking/config"
)

type uploadUseCase interface {
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*common.Image, error)
}

type uploadHandler struct {
	uploadUC uploadUseCase
	cfg      *config.Config
}

func NewUploadHandler(cfg *config.Config, uploadUseCase uploadUseCase) *uploadHandler {
	return &uploadHandler{uploadUC: uploadUseCase, cfg: cfg}
}
