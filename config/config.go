package config

//ConfStruct - defines the required environment variable for the application 
type ConfStruct struct {
	Port    string `env:"PORT"`
	DB_host string `env:"DB_HOST"`
	DB_user string `env:"DB_USER"`
	DB_pass string `env:"DB_PASS"`
	DB_port string `env:"DB_PORT"`
	DB_name string `env:"DB_NAME"`
}

var Elements ConfStruct
