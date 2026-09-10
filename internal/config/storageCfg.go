package config

import (
    "time"
)

type StorageConfig struct {
    Type     string         `yaml:"type"`
    Endpoint string         `yaml:"endpoint"`
    Timeout  time.Duration  `yaml:"timeout"`
    Bucket   string         `yaml:"bucket"`
    Upload   UploadConfig   `yaml:"upload"`
    Presign  PresignConfig  `yaml:"presign"`
    AccessKey string         `yaml:"access_key"`
    SecretKey string         `yaml:"secret_key"`
}

type UploadConfig struct {
    MaxSize int64 `yaml:"max_size"`
}

type PresignConfig struct {
    Enabled bool          `yaml:"enabled"`
    Expire  time.Duration `yaml:"expire"`
}