package handlers

import (
	"golang.org/x/net/context"
	"github.com/fr05t1k/wallet/wallet"
	"log"
)

type Server struct {

}

func (s *Server) GetTotal(ctx context.Context, in *wallet.Total) (*wallet.Total, error) {
	log.Println("I recieve request GetTotal")

	value := int64(1)
	return &wallet.Total{Value: &value}, nil
}
