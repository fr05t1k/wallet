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
	RunWallet()

}

func RunWallet() {
	runner := GetWalletServer()
	runner.Run(config.GetConfig().GrpcPort)
}

func GetWalletServer() *server.Wallet {
	db := database.Connect(config.GetConfig().MongoDbHost, config.GetConfig().MongoDbDatabase)
	operations := operation.Manager{DB: db}

	return &server.Wallet{Operations: &operations}
}
