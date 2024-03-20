package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Avinashanakal/grocery-delivery-service/api"
	"github.com/Avinashanakal/grocery-delivery-service/api/server/http"
	"github.com/Avinashanakal/grocery-delivery-service/platform/config"
	"github.com/Avinashanakal/grocery-delivery-service/platform/database"
)

func main() {

	// Initialize logrus
	logger := initLogger()

	// Fetching the configurations
	configs, err := config.New()
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	// Connecting to db
	dbConfig, _ := configs.Database()
	db, err := database.New(dbConfig)
	if err != nil {
		dbConfig.Password = "***"
		return
	}
	connection, _ := db.DB()
	defer connection.Close()

	// Create API instance and pass the logger
	apiInstance := api.New(configs, db, logger)

	// Create HTTP server and pass the API instance
	router := http.New(apiInstance)
	router.Run(":90")
}

// initLogger initializes logrus with desired configurations.
func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
