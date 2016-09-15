package wallet

import (
	"net"
	"google.golang.org/grpc"
	"fmt"
	"log"
	"golang.org/x/net/context"
	"github.com/fr05t1k/wallet/wallet"
)

type WalletServer struct {
}

func (server *WalletServer)GetTotal(ctx context.Context, total *wallet.Total) (*wallet.Total, error) {
	return total, nil
}

func (server *WalletServer)AddOperation(ctx context.Context, operation *wallet.Operation) (*wallet.Total, error) {
	return &wallet.Total{Value:int64(0)}, nil
}

func (server *WalletServer)Run(port uint16) {
	log.Println("Starting GRPC server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic("Can't run grpc server")
	}
	grpcServer := grpc.NewServer()

	wallet.RegisterWalletServer(grpcServer, &WalletServer{})
	log.Println(fmt.Sprintf("GRPC server listening on: :%d", port))
	err = grpcServer.Serve(lis)
	log.Fatalf("Failed start GRPC server %v", err)
	panic("Failed start GRPC server")
}

