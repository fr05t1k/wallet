package main

import (
	_ "github.com/joho/godotenv"
	database "github.com/fr05t1k/wallet/db"
)
import (
	"github.com/fr05t1k/wallet/server"
	"github.com/fr05t1k/wallet/config"
	"github.com/fr05t1k/wallet/operation"
)

func main() {
	db := database.Connect(config.GetConfig().MongoDbHost, config.GetConfig().MongoDbDatabase)
	operations := operation.Manager{DB:db}

	runner := &server.Wallet{Operations: &operations}
	runner.Run(config.GetConfig().GrpcPort)
}
