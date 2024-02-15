package uploadusecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"iTask/common"
)

func (uc *uploadUseCase) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*common.Image, error) {
	fileName := fileHeader.Filename

	file, err := fileHeader.Open()
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	defer file.Close()
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	dataBytes := make([]byte, fileHeader.Size)
	if _, err := file.Read(dataBytes); err != nil {
		panic(common.ErrBadRequest(err))
	}

	pathFile := fmt.Sprintf("%s/%s", uc.cfg.AWS.S3FolderImages, fileName)
	img, err := uc.s3Provider.PutObject(ctx, dataBytes, pathFile)
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	return img, nil
}
