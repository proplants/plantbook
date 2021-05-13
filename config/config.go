package config

// ConfStruct - defines the required environment variable for the application 
type ConfStruct struct {
	Port    string `env:"PORT"`
	DBHost string `env:"DB_HOST"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBPort string `env:"DB_PORT"`
	DBName string `env:"DB_NAME"`
}

var Elements ConfStruct
