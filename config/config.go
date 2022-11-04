package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JwtSecret string   `mapstructure:"jwt_secret"`
	Server    Server   `mapstructure:"server"`
	Postgres  Postgres `mapstructure:"postgres"`
}

type Server struct {
	Port       string `mapstructure:"port"`
	SwaggerUrl string `mapstructure:"swagger_json"`
}

type Postgres struct {
	ConnString     string `mapstructure:"conn_string"`
	MaxConnections int32  `mapstructure:"max_connections"`
}

func New() (Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	var conf Config
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return conf, err
	}

	// unmarshal
	err = viper.Unmarshal(&conf)
	return conf, err
}
