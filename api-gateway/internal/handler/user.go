package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq service.DouyinUserRegisterRequest
	PanicIfUserError(ginCtx.BindQuery(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken(userResp.UserId)
	r := res.UserResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  e.GetMsg(userResp.StatusCode),
		UserId:     userResp.UserId,
		Token:      token,
	}
	ginCtx.JSON(http.StatusOK, r)
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq service.DouyinUserLoginRequest
	PanicIfUserError(ginCtx.BindQuery(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken(userResp.UserId)
	r := res.UserResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  e.GetMsg(userResp.StatusCode),
		UserId:     userResp.UserId,
		Token:      token,
	}
	ginCtx.JSON(http.StatusOK, r)
}

// 用户信息
func UserShow(ginCtx *gin.Context) {
	var userReq service.DouyinUserRequest
	PanicIfUserError(ginCtx.BindQuery(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(userReq.Token)
	userResp, err := userService.UserShow(context.Background(), &userReq)
	relationReq := &service.UserId_Request{
		UserId:      userResp.User.Id,
		RequesterId: claims.UserID,
	}
	relationInfo, err := relationService.UserRelationInfoById(context.Background(), relationReq)
	PanicIfRelationError(err)
	userResp.User.FollowCount = relationInfo.FollowCount
	userResp.User.FollowerCount = relationInfo.FollowerCount
	userResp.User.IsFollow = relationInfo.IsFollow
	PanicIfUserError(err)
	r := res.UserResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  e.GetMsg(userResp.StatusCode),
		Data:       userResp.User,
	}
	ginCtx.JSON(http.StatusOK, r)
}
