package s3provider

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"paradise-booking/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (p *s3Provider) PutObject(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.bucket),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	imgRes := &common.Image{
		Url: fmt.Sprintf("%s/%s", p.domain, dst),
	}
	return imgRes, nil
}
