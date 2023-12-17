package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Nurka144/golang-service/config"
	"github.com/Nurka144/golang-service/internal/routes"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	logrus.SetOutput(file)

	if err := config.InitConfig(); err != nil {
		logrus.Info("Ошибка загрузки переменных окружений : ", err.Error())
	}

	srv := routes.InitRoutes()

	logrus.Info("Прослушивание и обслуживание HTTP на : " + viper.GetString("port"))
	srv.Run(":" + viper.GetString("port"))
}
