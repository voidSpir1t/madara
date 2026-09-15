package config

import (
	"time"
)

type StorageConfig struct {
	Type      string        `yaml:"type"`
	Endpoint  string        `yaml:"endpoint"`
	Bucket    string        `yaml:"bucket"`
	Upload    UploadConfig  `yaml:"upload"`
	Presign   PresignConfig `yaml:"presign"`
	AccessKey string        `yaml:"accesskey"`
	SecretKey string        `yaml:"secretkey"`
}

type UploadConfig struct {
	MaxSize int64 `yaml:"max_size"`
}

type PresignConfig struct {
	Enabled bool          `yaml:"enabled"`
	Expire  time.Duration `yaml:"expire"`
}
