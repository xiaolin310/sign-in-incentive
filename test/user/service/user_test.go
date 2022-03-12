package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	userV1 "sign-in/api/user/v1"
	"testing"
)

var (
	ctx = context.TODO()
	client = newUserClient("127.0.0.1:9001")

)

func TestGetUserByName(t *testing.T) {
	resp, err := client.GetUserByName(ctx, &userV1.GetUserByNameRequest{
		Name: "Zhangsan",
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Get user by name reply: %v", resp.User)
	}
}

func TestGetUserById(t *testing.T) {
	resp, err := client.GetUserById(ctx, &userV1.GetUserByIdRequest{
		Id: []int64{1, 2},
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Get user by Id reply: %v", resp.User)
	}
}

func TestSearchUserByName(t *testing.T) {
	resp, err := client.SearchUserByName(ctx, &userV1.SearchUserByNameRequest{
		Name: "Shelly",
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Search user by name reply: %v", resp.User)
	}
}

func TestRemoveUser(t *testing.T) {
	resp, err := client.RemoveUser(ctx, &userV1.RemoveUserRequest{
		Id: []int64{1, 2},
	})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Remove user by Id reply: %v", resp)
}

func TestSaveUser(t *testing.T) {
	resp, err := client.SaveUser(ctx, &userV1.SaveUserRequest{
		User: &userV1.User{
			Name: "Lisi",
			Password: "123456",
			Phone: "13223567901",
		},
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Save user reply: %v", resp.User)
	}
}

func newUserClient(addr string) userV1.UserServiceClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return userV1.NewUserServiceClient(conn)

}