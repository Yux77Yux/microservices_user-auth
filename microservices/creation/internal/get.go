package internal

import (
	"context"
	"log"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"

	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	"github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
)

func GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	// 取数据
	creationId := req.GetCreationId()
	creation, err := cache.GetCreationInfo(ctx, creationId)
	if err != nil {
		return &generated.GetCreationResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Internal Server Error",
				Details: err.Error(),
			},
		}, nil
	}
	if creation != nil {
		return &generated.GetCreationResponse{
			CreationInfo: creation,
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_SUCCESS,
				Code:   "200",
			},
		}, nil
	}

	creation, err = db.GetDetailInTransaction(ctx, creationId)
	if err != nil {
		return &generated.GetCreationResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Internal Server Error",
				Details: err.Error(),
			},
		}, nil
	}

	// 存作品至redis
	go func(creation *generated.CreationInfo) {
		err := messaging.SendMessage(messaging.StoreCreationInfo, messaging.StoreCreationInfo, creation)
		if err != nil {
			log.Printf("error: GetCreation SendMessage %v", err)
		}
	}(creation)

	return &generated.GetCreationResponse{
		CreationInfo: creation,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func GetCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)

	ids := req.GetIds()
	creations, err := db.GetCreationCardInTransaction(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.CreationInfoGroup = creations
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetPublicCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)

	ids := req.GetIds()
	creations, err := db.GetCreationCardInTransaction(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	length := len(ids)
	filteredCreations := make([]*generated.CreationInfo, 0, length)
	for _, info := range creations {
		if info.GetCreation().GetBaseInfo().GetStatus() == generated.CreationStatus_PUBLISHED {
			filteredCreations = append(filteredCreations, info)
		}
	}

	response.CreationInfoGroup = filteredCreations
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 拿用户id，然后author_id = id 作为筛选条件
func GetSpaceCreations(ctx context.Context, req *generated.GetSpaceCreationsRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)
	id := req.GetUserId()

	ids, err := cache.GetSpaceCreationList(ctx, id)
	if err != nil {
		log.Printf("error: cache GetSpaceCreationList %v", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	infos, err := db.GetCreationCardInTransaction(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.CreationInfoGroup = infos
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}
