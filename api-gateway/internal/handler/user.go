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

// 用户登录
func UserShow(ginCtx *gin.Context) {
	var userReq service.DouyinUserRequest
	PanicIfUserError(ginCtx.BindQuery(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserShow(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.UserResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  e.GetMsg(userResp.StatusCode),
		Data:       userResp.User,
	}
	ginCtx.JSON(http.StatusOK, r)
}
