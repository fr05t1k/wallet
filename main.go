package main

import (
	_ "github.com/joho/godotenv"
	"github.com/fr05t1k/wallet/db"
)
import (
	server "github.com/fr05t1k/wallet/grpc"
	"github.com/fr05t1k/wallet/config"
)


func init() {
	db.Connect(config.GetConfig().MongoDbHost, config.GetConfig().MongoDbDatabase)
	runner := &server.WalletServer{}

	runner.Run(config.GetConfig().GrpcPort)
}

func main() {

}
