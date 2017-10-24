package config

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("eulerexec")
	viper.AddConfigPath(".")
	viper.AddConfigPath("~/.eulerexec")
	viper.AddConfigPath("./etc")
	viper.AddConfigPath("/etc/eulerexec/")
	logrus.Info("Loading Config")
	e := viper.ReadInConfig()
	if e != nil {
		logrus.Fatal(e)
	}
}
