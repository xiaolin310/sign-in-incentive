package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	recordV1 "sign-in/api/record/v1"
	"testing"
)

var (
	ctx = context.TODO()
	client = newRecordClient("127.0.0.1:9003")

)

func TestGetSignInInfo(t *testing.T) {
	resp, err := client.GetSignInInfo(ctx, &recordV1.GetSignInInfoRequest{
		User: 9,
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Get SignIn info reply: %v", resp)
	}
}

func TestSignIn(t *testing.T) {
	resp, err := client.SignIn(ctx, &recordV1.SignInRequest{
		User: 9,
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("User Sign In reply: %v", resp)
	}
}

func TestGetUserSignInRecord(t *testing.T) {
	resp, err := client.GetUserSignInRecord(ctx, &recordV1.GetUserSignInRecordRequest{
		User: 9,
		Day: []string{"20220311", "20220312", "20220313"},
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Get User SignIn Record reply: %v", resp.UserRecord)
	}
}

func newRecordClient(addr string) recordV1.RecordServiceClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return recordV1.NewRecordServiceClient(conn)
}