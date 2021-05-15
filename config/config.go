package config

// ConfStruct - defines the required environment variable for the application 
type ConfStruct struct {
	Port    string `env:"APP_PORT"`
	DBHost string `env:"APP_DB_HOST"`
	DBUser string `env:"APP_DB_USER"`
	DBPass string `env:"APP_DB_PASS"`
	DBPort string `env:"APP_DB_PORT"`
	DBName string `env:"APP_DB_NAME"`
}

var Elements ConfStruct


