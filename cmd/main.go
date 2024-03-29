package main

import (
	"medods"
	"medods/pkg/database"
	"medods/pkg/handler"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	err := database.InitDB(os.Getenv("MONGO_URI"), "users")
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	srv := new(medods.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Print("Server started")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("local")
	return viper.ReadInConfig()
}
