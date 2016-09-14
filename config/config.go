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

var configInstace config;

func init() {
	port, err := strconv.Atoi(os.Getenv("WALLET_PORT"))
	if err != nil {
		port = 50051
	}
	configInstace.GrpcPort = uint16(port)

	configInstace.MongoDbHost = os.Getenv("WALLET_MONGODB_HOST")
	configInstace.MongoDbDatabase = os.Getenv("WALLET_MONGODB_DB_NAME")

}

func GetConfig() config {
	return configInstace
}



