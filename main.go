package main

import (
	"os"
	"gopkg.in/mgo.v2"
	_ "github.com/joho/godotenv"
	"github.com/fr05t1k/wallet/db"
	"github.com/fr05t1k/wallet/wallet"
	"strconv"
)

var (
	port, _ = strconv.Atoi(os.Getenv("WALLET_PORT"))
	mongoConnectionString = os.Getenv("WALLET_MONGODB_CONNECTION_STRING")
	mongoDBName = os.Getenv("WALLET_MONGODB_DB_NAME")
)


func init() {
	db.Connect(mongoConnectionString, mongoDBName)
	runner := &wallet.WalletServerRunner{}

	runner.Run(uint16(port))
}

func main() {

}
