package wallet

import (
	"net"
	"google.golang.org/grpc"
	"fmt"
	"log"
	"golang.org/x/net/context"
)

type WalletServerRunner struct {
}

func (server *WalletServerRunner)GetTotal(ctx context.Context, total *Total) (*Total, error) {
	return total, nil
}

func (server *WalletServerRunner)Run(port uint16) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	RegisterWalletServer(grpcServer, &WalletServerRunner{})

	grpcServer.Serve(lis)
}

