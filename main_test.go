package main

import (
	"testing"
	"github.com/fr05t1k/wallet/wallet"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

func TestAddOperation(t *testing.T) {

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := wallet.NewWalletClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond * 500)
	total, err := client.GetTotal(ctx, &empty.Empty{})
	if err != nil {
		t.Errorf("Can't get total: %v", err)
		t.FailNow()
	}

	operations := wallet.Operation{Value:int64(1000), Note:"Test"}

	_, err = client.AddOperation(ctx, &operations)

	if err != nil {
		t.Errorf("Can't add operation %v", err)
		t.FailNow()
	}

	newTotal, err := client.GetTotal(ctx, &empty.Empty{})

	if err != nil {
		t.Errorf("Wrong response %v", err)
		t.FailNow()
	}

	difference := newTotal.Value - total.Value
	if difference != 1000 {
		t.Errorf("Incorrect add operation difference is: %d", difference)
	}
}
