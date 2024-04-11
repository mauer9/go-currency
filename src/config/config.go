package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Config constants
const DefaultChanCapacity = 2

// MainConfig - global config variable
var MainConfig *Config

// Main config struct
type Config struct {
	MicrosoftSQLConfig
	Port string `json:"app-port"`
}

// Remote DB config struct
type MicrosoftSQLConfig struct {
	MSSQLHost string `json:"mssql-host"`
	MSSQLPort int    `json:"mssql-port"`
	MSSQLDB   string `json:"mssql-db"`
	MSSQLUser string `json:"mssql-user"`
	MSSQLPass string `json:"mssql-pass"`
}

// Initialize service configuration from ENV
func InitConfigs() error {
	var cfg Config

	dir, err := os.Getwd()
	if err != nil {
		return errors.New("failed to get current working directory: " + err.Error())
	}

	jsonFilePath := filepath.Join(dir, "config.json")

	file, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return errors.New("failed to read config file: " + err.Error())
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return err
	}

	MainConfig = &cfg
	return nil
}
