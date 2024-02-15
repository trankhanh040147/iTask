package s3provider

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type S3Provider interface {
	PutObject(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

type s3Provider struct {
	bucket  string
	region  string
	apiKey  string
	secret  string
	domain  string
	session *session.Session
}

func NewS3Provider(config *config.Config) *s3Provider {
	provider := &s3Provider{
		bucket: config.AWS.S3Bucket,
		region: config.AWS.Region,
		apiKey: config.AWS.APIKey,
		secret: config.AWS.SecretKey,
		domain: config.AWS.S3Domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	provider.session = s3Session
	return provider
}
