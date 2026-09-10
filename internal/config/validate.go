package config

import "fmt"

func (c *Config) Validate() error {
	if c.Server.Port <= 0 {
		return fmt.Errorf("server.port must be greater than 0")
	}

	if len(c.Models) == 0 {
		return fmt.Errorf("models cannot be empty")
	}

	for _, model := range c.Models {
		if model.Provider == "" {
			return fmt.Errorf("models.%s.provider is required", model.ModelName)
		}

		if model.ModelName == "" {
			return fmt.Errorf("models.%s.name is required", model.ModelName)
		}
	}

	return nil
}
