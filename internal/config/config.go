package config

type Config struct {
	Server ServerConfig  `yaml:"server"`
	Models []ModelConfig `yaml:"models"`
	Memory MemoryConfig  `yaml:"memory"`
	Storage StorageConfig `yaml:"storage"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type ModelConfig struct {
	Provider string `yaml:"provider"`
	Name	string `yaml:"display_name"`
	ModelName	string `yaml:"model_name"`
	BaseURL  string `yaml:"base_url"`
	APIKey   string `yaml:"api_key"`
}

type MemoryConfig struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}
