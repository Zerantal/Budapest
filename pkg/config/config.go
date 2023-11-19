package config

import (
	"encoding/json"
	"io/ioutil"

	"budapest/pkg/repo"
)

// Config represents the configuration structure for the application.
type Config struct {
	DatabaseConfig repo.DatabaseConfig `json:"databaseConfig"`
	DatabaseName   string              `json:"databaseName"`
}

// LoadConfig reads the configuration from a JSON file.
func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
