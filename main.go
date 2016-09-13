package main

import (
	"os"
	_ "github.com/joho/godotenv"
	"github.com/fr05t1k/wallet/db"
	"strconv"
)
import server "github.com/fr05t1k/wallet/grpc"

var (
	port, _ = strconv.Atoi(os.Getenv("WALLET_PORT"))
	mongoConnectionString = os.Getenv("WALLET_MONGODB_CONNECTION_STRING")
	mongoDBName = os.Getenv("WALLET_MONGODB_DB_NAME")
)

func init() {
	db.Connect(mongoConnectionString, mongoDBName)
	runner := &server.WalletServer{}

	runner.Run(uint16(port))
}

func main() {

}
