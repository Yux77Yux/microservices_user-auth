package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	auth "github.com/Yux77Yux/platform_backend/generated/auth" // 你生成的 package 名字
	user "github.com/Yux77Yux/platform_backend/generated/user" // 你生成的 package 名字
)

type AuthClient struct {
	connection *grpc.ClientConn
}

func NewAuthClient() (*AuthClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &AuthClient{
		connection: conn,
	}

	return client, nil
}

func (c *AuthClient) Login(ctx context.Context, userLogin *user.UserLogin) (*auth.LoginResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := auth.NewAuthServiceClient(c.connection)

	// 创建请求
	req := &auth.LoginRequest{UserAuth: &user.UserAuth{
		UserId:   userLogin.GetUserDefault().GetUserId(),
		UserRole: userLogin.GetUserRole(),
	}}

	response, err := client.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
