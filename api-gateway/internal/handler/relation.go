package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func RelationAction(ginCtx *gin.Context) {
	var relationReq service.DouyinRelationActionRequest
	PanicIfRelationError(ginCtx.ShouldBindWith(&relationReq, binding.Query))
	// 从gin.Key中取出服务实例
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(relationReq.Token)
	relationReq.UserId = claims.UserID
	resp, err := relationService.RelationAction(context.Background(), &relationReq)
	PanicIfRelationError(err)
	r := res.Response{
		StatusCode: resp.StatusCode,
		StatusMsg:  e.GetMsg(resp.GetStatusCode()),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func FollowList(ginCtx *gin.Context) {
	var relationReq service.DouyinRelationFollowListRequest
	PanicIfRelationError(ginCtx.ShouldBindWith(&relationReq, binding.Query))
	// 从gin.Key中取出服务实例
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	resp, err := relationService.FollowList(context.Background(), &relationReq)
	PanicIfRelationError(err)
	for i, _ := range resp.UserList {
		userReq := &service.UserIdRequest{UserId: resp.UserList[i].Id}
		user, err := userService.UserById(context.Background(), userReq)
		PanicIfUserError(err)
		resp.UserList[i].Name = user.Name
	}
	r := res.UserListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  e.GetMsg(resp.GetStatusCode()),
		Data:       resp.UserList,
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func FollowerList(ginCtx *gin.Context) {
	var relationReq service.DouyinRelationFollowerListRequest
	PanicIfRelationError(ginCtx.ShouldBindWith(&relationReq, binding.Query))
	// 从gin.Key中取出服务实例
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	resp, err := relationService.FollowerList(context.Background(), &relationReq)
	PanicIfRelationError(err)
	for i, _ := range resp.UserList {
		userReq := &service.UserIdRequest{UserId: resp.UserList[i].Id}
		user, err := userService.UserById(context.Background(), userReq)
		PanicIfUserError(err)
		resp.UserList[i].Name = user.Name
	}
	r := res.UserListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  e.GetMsg(resp.GetStatusCode()),
		Data:       resp.UserList,
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func FriendList(ginCtx *gin.Context) {
	var relationReq service.DouyinRelationFriendListRequest
	PanicIfRelationError(ginCtx.ShouldBindWith(&relationReq, binding.Query))
	// 从gin.Key中取出服务实例
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	resp, err := relationService.FriendList(context.Background(), &relationReq)
	PanicIfRelationError(err)
	for i, _ := range resp.UserList {
		userReq := &service.UserIdRequest{UserId: resp.UserList[i].Id}
		user, err := userService.UserById(context.Background(), userReq)
		PanicIfUserError(err)
		resp.UserList[i].Name = user.Name
	}
	r := res.UserListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  e.GetMsg(resp.GetStatusCode()),
		Data:       resp.UserList,
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}
