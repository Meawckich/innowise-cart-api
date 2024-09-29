package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DbURL string `mapstructure:"DB_URL"`
}

func InitConfig() (c Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("cfg")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	c.Port = viper.Get("PORT").(string)
	c.DbURL = viper.Get("DB_URL").(string)

	return
}
