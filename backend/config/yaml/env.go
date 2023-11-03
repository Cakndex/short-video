package env

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
}

type Config struct {
}

func init() {
	viper.SetConfigName("./config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error reading config file, %s", err)
		panic(err)
	}

}
