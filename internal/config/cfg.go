// Package config implements common config method and variable for plantbook project
package config

type HTTPD struct {
	Host       string `env:"PLANTBOOK_HTTPD_HOST"`
	Port       string `env:"PLANTBOOK_HTTPD_PORT" default:"8081"`
	MetricPort string `env:"PLANTBOOK_HTTPD_METRICPORT" default:"8082"`
}

type DB struct {
	Provider string `env:"PLANTBOOK_DB_PROVIDER" default:"postgres"`
	URL      string `env:"PLANTBOOK_DB_URL" default:"postgres://plantbook_admin:mypassword@postgresql:5432/plantbook_admin?sslmode=disable"`
}

type Log struct {
	Debug  bool   `env:"PLANTBOOK_LOG_DEBUG"`
	Format string `env:"PLANTBOOK_LOG_FORMAT" default:"console"`
}

type Config struct {
	HTTPD       HTTPD
	DB          DB
	Log         Log
	TokenSecret string `env:"PLANTBOOK_TOKEN_SECRET" default:"secret_password_for_token_sign"`
}
