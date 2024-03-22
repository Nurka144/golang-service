package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/example/test/config"
	"github.com/example/test/internal/repository/postgres"
	"github.com/example/test/internal/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "staging"
	}

	if errLoadEnv := config.LoadEnv(appEnv); errLoadEnv != nil {
		logrus.Info("Ошибка загрузки переменных окружений : ", errLoadEnv.Error())
		panic("Ошибка загрузки переменных окружений : " + errLoadEnv.Error())
	}

	if err := config.InitConfig(appEnv); err != nil {
		logrus.Info("Ошибка загрузки конфига : ", err.Error())
		panic("Ошибка загрузки конфига : " + err.Error())
	}

	db, errDBPg := postgres.InitConnectDBPg()

	if errDBPg != nil {
		logrus.Info("Ошибка подключение к БД Postgres : ", errDBPg.Error())
		panic("Ошибка подключение к БД Postgres : " + errDBPg.Error())
	}

	srv := routes.InitRoutes(db)

	go func() {
		if errServer := srv.Run(viper.GetString("port")); errServer != nil {
			logrus.Info(" Cервер упал с ошибкой : ", errServer.Error())
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGTERM, syscall.SIGINT)

	strQuit := <-quite

	logrus.Info("Сервер остановлен. Сигнал остановки : ", strQuit)

	os.Exit(0)
}
