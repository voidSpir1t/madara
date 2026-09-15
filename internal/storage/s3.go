package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	client  *s3.Client
	presign *s3.PresignClient
	bucket  string
	expires time.Duration
}

type S3Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	Region    string
	Expires   time.Duration
}

func NewS3Storage(cfg S3Config) (*S3Storage, error) {
	if cfg.Endpoint == "" {
		return nil, fmt.Errorf("s3 endpoint is required")
	}

	if cfg.Bucket == "" {
		return nil, fmt.Errorf("s3 bucket is required")
	}

	if cfg.Region == "" {
		cfg.Region = "cn"
	}

	if cfg.Expires <= 0 {
		cfg.Expires = 10 * time.Minute
	}

	awsCfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.Endpoint)
		o.UsePathStyle = true
	})
	presign := s3.NewPresignClient(client)

	return &S3Storage{
		client:  client,
		presign: presign,
		bucket:  cfg.Bucket,
		expires: cfg.Expires,
	}, nil
}

func (s *S3Storage) Put(ctx context.Context, key string, body io.Reader, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	return err
}

func (s *S3Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	output, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}

func (s *S3Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	return err
}

func (s *S3Storage) PresignGet(ctx context.Context, key string) (string, error) {
	req, err := s.presign.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = s.expires
		},
	)

	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func (s *S3Storage) PresignPut(ctx context.Context, key string, contentType string) (string, error) {
	req, err := s.presign.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:      aws.String(s.bucket),
			Key:         aws.String(key),
			ContentType: aws.String(contentType),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = s.expires
		},
	)
	if err != nil {
		return "", err
	}

	return req.URL, nil
}
