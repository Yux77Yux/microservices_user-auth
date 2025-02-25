package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) DelReviewer(ctx context.Context, req *generated.DelReviewerRequest) (*generated.DelReviewerResponse, error) {
	log.Println("info: update user service start")

	response, err := internal.DelReviewer(req)
	if err != nil {
		log.Println("error: update user occur fail")
		return response, err
	}

	log.Println("info: update user occur success")
	return response, nil
}

func (s *Server) UpdateUserSpace(ctx context.Context, req *generated.UpdateUserSpaceRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user service start")

	response, err := internal.UpdateUserSpace(req)
	if err != nil {
		log.Println("error: update user occur fail")
		return response, err
	}

	log.Println("info: update user occur success")
	return response, nil
}

func (s *Server) UpdateUserAvatar(ctx context.Context, req *generated.UpdateUserAvatarRequest) (*generated.UpdateUserAvatarResponse, error) {
	log.Println("info: update user avatar service start")

	response, err := internal.UpdateUserAvatar(req)
	if err != nil {
		log.Println("error: update user avatar occur fail")
		return response, err
	}

	log.Printf("info: update user avatar occur success: %v", response)
	return response, nil
}

func (s *Server) UpdateUserStatus(ctx context.Context, req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user status service start")

	response, err := internal.UpdateUserStatus(req)
	if err != nil {
		log.Println("error: update user status occur fail")
		return response, nil
	}

	log.Println("info: update user status occur success")
	return response, nil
}

func (s *Server) UpdateUserBio(ctx context.Context, req *generated.UpdateUserBioRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user status service start")

	response, err := internal.UpdateUserBio(req)
	if err != nil {
		log.Println("error: update user status occur fail")
		return response, nil
	}

	log.Println("info: update user status occur success")
	return response, nil
}
