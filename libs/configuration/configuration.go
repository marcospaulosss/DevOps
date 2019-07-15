package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	v   *viper.Viper
	env string
}

var conf *Configuration

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	fmt.Println("Using env:", env)
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(os.Getenv("CONFIG_DIR"))
	v.AddConfigPath(".")
	v.SetEnvPrefix(env)
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Nao consegui ler o arquivo de configuracao:", err)
	}

	conf = &Configuration{v, env}
}

func Get() *Configuration {
	return conf
}

func (this *Configuration) GetEnvConfString(key string) string {
	return this.v.GetString(fmt.Sprintf("%s.%s", this.env, key))
}

func (this *Configuration) GetString(key string) string {
	return this.v.GetString(key)
}

func (this *Configuration) GetEnvConfInteger(key string) int {
	return this.v.GetInt(fmt.Sprintf("%s.%s", this.env, key))
}

func (this *Configuration) GetInteger(key string) int {
	return this.v.GetInt(key)
}
