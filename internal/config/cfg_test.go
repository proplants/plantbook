// Package config implements common config method and variable for plantbook project
package config

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	portEnv := map[string]string{"PLANTBOOK_HTTPD_PORT": "8080"}
	hostPortEnv := map[string]string{"PLANTBOOK_HTTPD_HOST": "127.0.0.1", "PLANTBOOK_HTTPD_PORT": "8000"}
	var configHostPort Config
	configHostPort = Defaults
	configHostPort.HTTPD.Port = "8000"
	configHostPort.HTTPD.Host = "127.0.0.1"

	tests := []struct {
		name     string
		defaults *Config
		cfg      *Config
		want     *Config
		env      map[string]string
		wantErr  bool
	}{
		{"empty_env_err", &Defaults, &Config{}, &Defaults, nil, false},
		{"port_env_ok", &Defaults, &Config{}, &Defaults, portEnv, false},
		{"hostport_env_ok", &Defaults, &Config{}, &configHostPort, hostPortEnv, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for envName, envValue := range tt.env {
				err := os.Setenv(envName, envValue)
				if err != nil {
					t.Fatalf("unexpected set env error, %s", err)
				}
			}
			if err := Read(tt.defaults, tt.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if tt.want != nil && tt.cfg != nil && !reflect.DeepEqual(*tt.want, *tt.cfg) {
				t.Errorf("expected cfg %+v, got %+v", *tt.want, *tt.cfg)
			}
		})
	}
}
