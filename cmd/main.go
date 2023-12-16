package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/Nurka144/golang-service/internal/routes"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	logrus.SetOutput(file)

	errLoadEnv := godotenv.Load()

	if errLoadEnv != nil {
		logrus.Info("Ошибка загрузки переменных окружений : ", errLoadEnv)
	}

	srv := routes.InitRoutes()

	logrus.Info("Прослушивание и обслуживание HTTP на : " + os.Getenv("PORT"))
	srv.Run(":" + os.Getenv("PORT"))
}
