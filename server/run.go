package server

import (
	"fmt"
	"github.com/fr05t1k/wallet/wallet"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (server *Wallet) Run(port uint16) {
	log.Println("Starting GRPC server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	server.GrpcServer = grpc.NewServer()

	wallet.RegisterWalletServer(server.GrpcServer, server)
	log.Println(fmt.Sprintf("GRPC server listening on: :%d", port))
	err = server.GrpcServer.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("Failed to start GRPC server %v", err))
	}
}
