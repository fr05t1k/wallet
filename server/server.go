package server

import (
	"errors"
	"github.com/fr05t1k/wallet/operation"
	"github.com/fr05t1k/wallet/wallet"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

type Wallet struct {
	Operations *operation.Manager
}

func (server *Wallet) GetTotal(ctx context.Context, e *empty.Empty) (*wallet.Total, error) {
	return server.Operations.GetTotal()
}

func (server *Wallet) AddOperation(ctx context.Context, operation *wallet.Operation) (*wallet.Total, error) {
	err := server.Operations.AddOperation(operation)
	if err != nil {
		return nil, errors.New("Can't create operation")
	}
	return server.Operations.GetTotal()
}
