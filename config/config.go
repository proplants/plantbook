package config

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

// ConfStruct - defines the required environment variable for the application
type ConfStruct struct {
	Port   string `env:"APP_PORT"`
	DBHost string `env:"APP_DB_HOST"`
	DBPort string `env:"APP_DB_PORT"`
	DBUser string `env:"APP_DB_USER"`
	DBPass string `env:"APP_DB_PASS" hide:"yes"`
	DBName string `env:"APP_DB_NAME"`
}

// New - this function returns the configuration structure
func New() ConfStruct {
	newConfif := &ConfStruct{}
	return *newConfif
}

// ConfLog - output of the configuration to the log. 
// If the field has a "hide" tag with value "yes"- this field is not displayed
func ConfLog(conf ConfStruct) {
	st := reflect.TypeOf(conf)
	sv := reflect.ValueOf(conf)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if field.Tag.Get("hide") != "yes" {
			logrus.Infof("%v: %v\n", field.Tag.Get("env"), sv.FieldByName(field.Name))
		}
	}
}
