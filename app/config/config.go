package config

import (
	"os"
	"reflect"

	"github.com/caarlos0/env/v11"
)

var Config struct {
	Env  string `env:"ENV,required"`
	Port string `env:"PORT,required"`
}

func init() {
	var err error

	if isUnitTest() {
		return
	}

	if err = env.ParseWithOptions(&Config, env.Options{FuncMap: map[reflect.Type]env.ParserFunc{}}); err != nil {
		panic(err)
	}
}

func isUnitTest() bool {
	return os.Getenv("UNIT_TEST") == "TRUE"
}
