package storage

import (
	"context"
	"fmt"
	"io"

    "github.com/voidSpir1t/madara/internal/config"
)

type Storage interface {
    Put(
        ctx context.Context,
        key string,
        body io.Reader,
        contentType string,
    ) error

    Get(
        ctx context.Context,
        key string,
    ) (io.ReadCloser, error)

    Delete(
        ctx context.Context,
        key string,
    ) error

    PresignGet(
        ctx context.Context,
        key string,
    ) (string, error)

    PresignPut(
        ctx context.Context,
        key string,
        contentType string,
    ) (string, error)
}


func NewStorage(cfg config.StorageConfig) (Storage, error) {
	switch cfg.Type {
	case "seaweedfs":
        s3cfg := S3Config{
        Endpoint:       cfg.Endpoint,
        AccessKey:      cfg.AccessKey,
        SecretKey:      cfg.SecretKey,
        Bucket:         cfg.Bucket,
        Expires:        cfg.Presign.Expire,
    }
		return NewS3Storage(s3cfg)
	default:
		return nil, fmt.Errorf(
			"unsupported storage type: %s",
			cfg.Type,
		)
	}
}
