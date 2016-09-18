package main

import (
	database "github.com/fr05t1k/wallet/db"
	_ "github.com/joho/godotenv"
)
import (
	"github.com/fr05t1k/wallet/config"
	"github.com/fr05t1k/wallet/operation"
	"github.com/fr05t1k/wallet/server"
)

func main() {
	db := database.Connect(config.GetConfig().MongoDbHost, config.GetConfig().MongoDbDatabase)
	operations := operation.Manager{DB: db}

	runner := &server.Wallet{Operations: &operations}
	runner.Run(config.GetConfig().GrpcPort)
}
