package internal

import (
	"context"
	"fmt"
	"log"
	"strings"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	comment "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	event "github.com/Yux77Yux/platform_backend/generated/common/event"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	interaction "github.com/Yux77Yux/platform_backend/generated/interaction"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	messaging "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
	"google.golang.org/grpc/metadata"
)

func WatchCreation(ctx context.Context, req *generated.WatchCreationRequest) (*generated.WatchCreationResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	ipv4 := ""
	if ok {
		ipv4s := md.Get("x-forwarded-for")
		if len(ipv4s) > 0 {
			ipv4 = strings.Split(ipv4s[0], ",")[0]
			ipv4 = strings.TrimSpace(ipv4) // 清理首尾空格
			log.Printf("ipv4: %s", ipv4)
		} else {
			log.Println("warning: x-forwarded-for not found in metadata")
		}
	}

	response := new(generated.WatchCreationResponse)
	id := req.GetCreationId()

	creation_client, err := client.GetCreationClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	creationResponse, err := creation_client.GetCreation(ctx, &creation.GetCreationRequest{
		CreationId: id,
	})
	if err != nil {
		var msg *common.ApiResponse
		if creationResponse != nil {
			msg = creationResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if creationResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = creationResponse.Msg
		return response, err
	}

	creationInfo := creationResponse.GetCreationInfo()
	userId := creationInfo.GetCreation().GetBaseInfo().GetAuthorId()
	user_client, err := client.GetUserClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	userResponse, err := user_client.GetUser(ctx, userId)
	if err != nil {
		var msg *common.ApiResponse
		if userResponse != nil {
			msg = userResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if userResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = userResponse.Msg
		return response, err
	}

	// 事件发布
	// 播放数
	if ipv4 == "" {
		log.Println("info: event not sent because ipv4 is empty")
	} else {
		go func(id int64, ipv4 string) {
			err := messaging.SendMessage(
				event.Exchange_EXCHANGE_ADD_VIEW.String(),
				event.RoutingKey_KEY_ADD_VIEW.String(),
				&common.ViewCreation{
					Id:   id,
					Ipv4: ipv4,
				},
			)
			if err != nil {
				log.Printf("error: SendMessage ADD_VIEW %v", err)
			}

		}(id, ipv4)
	}

	// 组装开始
	user := userResponse.GetUser()

	response.CreationInfo = creationInfo
	response.CreationUser = &common.UserCreationComment{
		UserDefault: &common.UserDefault{
			UserId:   user.GetUserDefault().GetUserId(),
			UserName: user.GetUserDefault().GetUserName(),
		},
		UserAvatar: user.GetUserAvatar(),
		UserBio:    user.GetUserBio(),
	}
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	// 组装完成返回至前端
	return response, nil
}

func SimilarCreations(ctx context.Context, req *generated.SimilarCreationsRequest) (*generated.GetCardsResponse, error) {
	response := new(generated.GetCardsResponse)
	id := req.GetCreationId()

	// 从 用户数据服务 调取相似列表
	interaction_client, err := client.GetInteractionClient()
	if err != nil {
		err = fmt.Errorf("error: interaction client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	interactionResponse, err := interaction_client.GetRecommendBaseCreation(ctx, &interaction.GetRecommendRequest{
		Id: id,
	})
	if err != nil {
		var msg *common.ApiResponse
		if interactionResponse != nil {
			msg = interactionResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if interactionResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = interactionResponse.Msg
		return response, err
	}

	creationIds := interactionResponse.GetCreations()
	if len(creationIds) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no member",
		}
		return response, nil
	}

	creation_client, err := client.GetCreationClient()
	if err != nil {
		err = fmt.Errorf("error: creation client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	creationResponse, err := creation_client.GetPublicCreationList(ctx, &creation.GetCreationListRequest{
		Ids: creationIds,
	})
	if err != nil {
		var msg *common.ApiResponse
		if creationResponse != nil {
			msg = creationResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if creationResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = creationResponse.Msg
		return response, err
	}

	creationInfos := creationResponse.GetCreationInfoGroup()
	length := len(creationInfos)
	if length <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no member",
		}
		return response, nil
	}

	userIds := make([]int64, length)
	for i, info := range creationInfos {
		userIds[i] = info.GetCreation().GetBaseInfo().GetAuthorId()
	}

	user_client, err := client.GetUserClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	userResponse, err := user_client.GetUsers(ctx, userIds)
	if err != nil {
		var msg *common.ApiResponse
		if userResponse != nil {
			msg = userResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if userResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = userResponse.Msg
		return response, err
	}

	// 构建 userId -> 用户信息的映射表
	users := userResponse.GetUsers()
	limit := len(users)
	if limit <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no member",
		}
		return response, nil
	}

	userMap := make(map[int64]*common.UserDefault, limit)
	for _, user := range users {
		userMap[user.GetUserDefault().GetUserId()] = user.GetUserDefault()
	}

	cards := make([]*generated.CreationCard, 0, length)
	for _, info := range creationInfos {
		creation := info.GetCreation()
		authorId := creation.GetBaseInfo().GetAuthorId()
		card := &generated.CreationCard{
			Creation:           creation,
			CreationEngagement: info.GetCreationEngagement(),
		}
		if user, exists := userMap[authorId]; exists {
			card.User = user
		}
		cards = append(cards, card)
	}

	response.Cards = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	response := new(generated.InitialCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	commentsResponse, err := comment_client.InitialComments(ctx, &comment.InitialCommentsRequest{
		CreationId: creationId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if commentsResponse != nil {
			msg = commentsResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if commentsResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = commentsResponse.Msg
		return response, err
	}
	comments := commentsResponse.GetComments()
	comments_len := len(comments)
	if comments_len <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no comment",
		}
		return response, nil
	}

	userIds := make([]int64, 0, comments_len)
	for i := 0; i < comments_len; i++ {
		userIds = append(userIds, comments[i].Comment.GetUserId())
	}

	userMap, err := getUsers(ctx, userIds)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	if len(userMap) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no user",
		}
		return response, nil
	}

	userMap[0] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	userMap[-1] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	topComments := make([]*generated.TopCommentInfo, comments_len)
	for i := 0; i < comments_len; i++ {
		comment := comments[i]
		topComments[i] = &generated.TopCommentInfo{
			CommentUser: userMap[comment.Comment.GetUserId()],
			TopComment:  comment,
		}
	}

	response.Comments = topComments
	response.Area = commentsResponse.GetCommentArea()
	response.PageCount = commentsResponse.GetPageCount()
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetTopCommentsResponse, error) {
	response := new(generated.GetTopCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	commentsResponse, err := comment_client.GetTopComments(ctx, &comment.GetTopCommentsRequest{
		CreationId: creationId,
		Page:       request.GetPage(),
	})
	if err != nil {
		var msg *common.ApiResponse
		if commentsResponse != nil {
			msg = commentsResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if commentsResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = commentsResponse.Msg
		return response, err
	}

	comments := commentsResponse.GetComments()
	comments_len := len(comments)
	if comments_len <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no comment",
		}
		return response, nil
	}

	userIds := make([]int64, 0, comments_len)
	for i := 0; i < comments_len; i++ {
		userIds = append(userIds, comments[i].Comment.GetUserId())
	}

	userMap, err := getUsers(ctx, userIds)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	if len(userMap) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no user",
		}
		return response, nil
	}

	userMap[0] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	userMap[-1] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	topComments := make([]*generated.TopCommentInfo, comments_len)
	for i := 0; i < comments_len; i++ {
		comment := comments[i]
		topComments[i] = &generated.TopCommentInfo{
			CommentUser: userMap[comment.Comment.GetUserId()],
			TopComment:  comment,
		}
	}

	response.Comments = topComments
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetSecondCommentsResponse, error) {
	response := new(generated.GetSecondCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	commentsResponse, err := comment_client.GetSecondComments(ctx, &comment.GetSecondCommentsRequest{
		CreationId: creationId,
		Root:       request.GetRoot(),
		Page:       request.GetPage(),
	})
	if err != nil {
		var msg *common.ApiResponse
		if commentsResponse != nil {
			msg = commentsResponse.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if commentsResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = commentsResponse.Msg
		return response, err
	}

	comments := commentsResponse.GetComments()
	comments_len := len(comments)
	if comments_len <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no comment",
		}
		return response, nil
	}

	userIds := make([]int64, 0, comments_len*2)
	for i := 0; i < comments_len; i++ {
		replyUserId := comments[i].GetReplyUserId()
		if replyUserId > 0 {
			userIds = append(userIds, replyUserId)
		}
		userIds = append(userIds, comments[i].Comment.GetUserId())
	}

	userMap, err := getUsers(ctx, userIds)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	if len(userMap) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no user",
		}
		return response, nil
	}

	userMap[0] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	userMap[-1] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	secondComments := make([]*generated.SecondCommentInfo, comments_len)
	for i := 0; i < comments_len; i++ {
		comment := comments[i]
		secondComments[i] = &generated.SecondCommentInfo{
			ReplyUser: userMap[comment.GetReplyUserId()],
		}
		secondComments[i].SecondComment = &generated.CommentInfo{
			CommentUser: userMap[comment.Comment.GetUserId()],
			Comment:     comment.GetComment(),
		}
	}

	response.Comments = secondComments
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func getUsers(ctx context.Context, userIds []int64) (map[int64]*common.UserCreationComment, error) {
	userMap := make(map[int64]*common.UserCreationComment)

	user_client, err := client.GetUserClient()
	if err != nil {
		return nil, err
	}
	userResponse, err := user_client.GetUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}
	if userResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		return nil, fmt.Errorf("error: %s", userResponse.Msg.GetDetails())
	}

	users := userResponse.GetUsers()
	if len(users) <= 0 {
		return nil, nil
	}

	for _, user := range users {
		userId := user.GetUserDefault().GetUserId()
		userMap[userId] = user
	}

	return userMap, nil
}
