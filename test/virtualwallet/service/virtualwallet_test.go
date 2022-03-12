package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	walletClientV1 "sign-in/api/virtualwallet/v1"
	"testing"
)

var (
	ctx = context.TODO()
	client = newWalletClient("127.0.0.1:9002")
)

func TestGetBalance(t *testing.T) {
	resp, err := client.GetBalance(ctx, &walletClientV1.GetBalanceRequest{
		User: 2,
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Get user balance reply: %v", resp.Balance)
	}
}

func TestDebit(t *testing.T) {
	resp, err := client.Debit(ctx, &walletClientV1.DebitRequest{
		User: 2,
		Amount: 100.0,
	})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("User debit balance of wallet reply: %v", resp)
}

func TestCredit(t *testing.T) {
	resp, err := client.Credit(ctx, &walletClientV1.CreditRequest{
		User: 2,
		Amount: 100.0,
	})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("User credit banlance of wallet reply: %v", resp)
}

func newWalletClient(addr string) walletClientV1.VirtualWalletClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return walletClientV1.NewVirtualWalletClient(conn)
}