package main

import (
	"medods"
	"medods/pkg/database"
	"medods/pkg/handler"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	err := database.InitDB(os.Getenv("MONGO_URI"), os.Getenv("DB_NAME"), os.Getenv("DB_COLLECTION_NAME"))
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	srv := new(medods.Server)
	if err := srv.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Print("Server started")
}
