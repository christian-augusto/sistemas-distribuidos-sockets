package config

import (
	"sistemas-distribuidos-sockets-server-golang/constants"

	"github.com/spf13/viper"
)

func LoadEnvironmentVars(serverEnv string) {
	if serverEnv == constants.PRODUCTION_ENV_VALUE {
		viper.AutomaticEnv()
	} else {
		viper.AddConfigPath("./config")

		viper.SetConfigName(constants.DEVELOPMENT_ENV_VALUE)
		viper.SetConfigType("env")

		viper.ReadInConfig()
	}
}

func IsDevelopmentEnv() bool {
	return viper.GetString(constants.ENV_NAME) == constants.DEVELOPMENT_ENV_VALUE
}

func IsProductionEnv() bool {
	return viper.GetString(constants.ENV_NAME) == constants.PRODUCTION_ENV_VALUE
}
