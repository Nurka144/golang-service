package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig(env string) error {

	viper.AddConfigPath("./config")

	configFile := fmt.Sprintf("config.%s", env)

	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.Set("keycloak.url", os.Getenv("KEYCLOAK_SERVER_URL"))
	viper.Set("keycloak.realm", os.Getenv("KEYCLOAK_REALM_NAME"))

	if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		return err
	}

	return nil
}

func LoadEnv(env string) error {
	envFile := fmt.Sprintf(".env.%s", env)

	err := godotenv.Load(envFile)
	if err != nil {
		return errors.New("Ошибка загрузки файла переменных окружений " + envFile + " : " + err.Error())
	}

	return nil
}
