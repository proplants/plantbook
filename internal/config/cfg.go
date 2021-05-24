// Package config implements common config method and variable for plantbook project
package config

import (
	env "github.com/kaatinga/env_loader"
)

type Config struct {
	HTTPD struct {
		Host string `env:"PLANTBOOK_HTTPD_HOST"`
		Port string `env:"PLANTBOOK_HTTPD_PORT"`
	}
	DB struct {
		Provider string `env:"PLANTBOOK_DB_PROVIDER"`
		URL      string `env:"PLANTBOOK_DB_URL"`
	}
	LOG struct {
		Debug  bool   `env:"PLANTBOOK_LOG_DEBUG"`
		Format string `env:"PLANTBOOK_LOG_FORMAT"`
	}
	TokenSecret string `env:"PLANTBOOK_TOKEN_SECRET"`
}

// Defaults config default values
var Defaults Config = Config{
	HTTPD: struct {
		Host string "env:\"PLANTBOOK_HTTPD_HOST\""
		Port string "env:\"PLANTBOOK_HTTPD_PORT\""
	}{
		Host: "",
		Port: "8080",
	},
	DB: struct {
		Provider string "env:\"PLANTBOOK_DB_PROVIDER\""
		URL      string "env:\"PLANTBOOK_DB_URL\""
	}{
		Provider: "postgres",
		URL:      "postgres://plantbook_admin:mypassword@postgresql:5432/plantbook_admin?sslmode=disable",
	},
	LOG: struct {
		Debug  bool   "env:\"PLANTBOOK_LOG_DEBUG\""
		Format string "env:\"PLANTBOOK_LOG_FORMAT\""
	}{
		Debug:  false,
		Format: "console",
	},
	TokenSecret: "secret_password_for_token_sign",
}

// Read fills passed cfg by reading environment variables,
// if some parameter not passed default value will be used
func Read(defaults, cfg *Config) error {
	err := env.LoadUsingReflect(&cfg)
	if err != nil {
		return err
	}
	if defaults == nil {
		return nil
	}
	// set default values
	if cfg.HTTPD.Host == "" {
		cfg.HTTPD.Host = defaults.HTTPD.Host
	}
	if cfg.HTTPD.Port == "" {
		cfg.HTTPD.Port = defaults.HTTPD.Port
	}
	if cfg.DB.Provider == "" {
		cfg.DB.Provider = defaults.DB.Provider
	}
	if cfg.DB.URL == "" {
		cfg.DB.URL = defaults.DB.URL
	}
	if cfg.LOG.Format == "" {
		cfg.LOG.Format = defaults.LOG.Format
	}
	if cfg.TokenSecret == "" {
		cfg.TokenSecret = defaults.TokenSecret
	}
	return nil
}
