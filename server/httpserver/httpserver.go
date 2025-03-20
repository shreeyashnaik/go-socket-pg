package main

import (
	"chat-app/common/constants/enums"
	"chat-app/common/database"
	"chat-app/common/utils"
	"chat-app/httpserver/api"
	"context"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func main() {

	// Load environment variables
	utils.LoadEnv()

	// Init database
	database.InitRedisClient(context.Background(), enums.WRITE_ACCESS)

	// Init routes
	multiplexer := api.InitRoutes()

	// Create http server
	server := http.Server{
		Addr:         viper.GetString("HTTP_PORT"),
		Handler:      multiplexer,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start http server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
