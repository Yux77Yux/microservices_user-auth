package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	log.Println("info: login service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.LoginResponse{}, err
	default:
		response, err := internal.Login(req)
		if err != nil {
			log.Println("error: login occur fail")
			return response, err
		}

		log.Println("info: login occur success")
		return response, nil
	}
}
