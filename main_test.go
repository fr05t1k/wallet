package main

import (
	"github.com/fr05t1k/wallet/wallet"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestAddOperation(t *testing.T) {
	go RunWallet()
	time.Sleep(time.Second * 1)

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := wallet.NewWalletClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	total, err := client.GetTotal(ctx, &empty.Empty{})
	if err != nil {
		t.Errorf("Can't get total: %v", err)
		t.FailNow()
	}

	now := timestamp.Timestamp{time.Now().Unix(), 0}
	operations := wallet.Operation{
		Value: int64(1000),
		Note:  "Test",
		Time:  &now,
	}

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
		t.Errorf("Incorrect add operation difference is: %d Old: %d New: %d",
			difference, total.Value, newTotal.Value)
	}
}
