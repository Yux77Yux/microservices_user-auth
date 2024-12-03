package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	auth "github.com/Yux77Yux/platform_backend/generated/auth" // 你生成的 package 名字
)

type AuthClient struct {
	connection *grpc.ClientConn
}

func NewAuthClient() (*AuthClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(auth_service_address)
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &AuthClient{
		connection: conn,
	}

	return client, nil
}

func (c *AuthClient) Login(user_id int64) (*auth.LoginResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := auth.NewAuthServiceClient(c.connection)

	// 创建请求
	req := &auth.LoginRequest{UserId: user_id}

	// 调用 gRPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := client.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
