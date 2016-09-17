package config

import (
	"strconv"
	"os"
)

type config struct {
	GrpcPort        uint16
	MongoDbHost     string
	MongoDbDatabase string
}

var configInstance config;

func init() {
	port, err := strconv.Atoi(os.Getenv("WALLET_PORT"))
	if err != nil {
		port = 50051
	}
	configInstance.GrpcPort = uint16(port)

	configInstance.MongoDbHost = os.Getenv("WALLET_MONGODB_HOST")
	configInstance.MongoDbDatabase = os.Getenv("WALLET_MONGODB_DB_NAME")

}

func GetConfig() config {
	return configInstance
}



