package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// UserLogin	用户登录
func (*UserService) UserLogin(ctx context.Context, req *service.DouyinUserLoginRequest) (resp *service.DouyinUserLoginResponse, err error) {
	var user repository.User
	resp = new(service.DouyinUserLoginResponse)
	resp.StatusCode = e.SUCCESS
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.UserId = user.UserId
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*UserService) UserRegister(ctx context.Context, req *service.DouyinUserRegisterRequest) (resp *service.DouyinUserRegisterResponse, err error) {
	var user repository.User
	resp = new(service.DouyinUserRegisterResponse)
	resp.StatusCode = e.SUCCESS
	user, err = user.UserCreate(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.UserId = user.UserId
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*UserService) UserShow(ctx context.Context, req *service.DouyinUserRequest) (resp *service.DouyinUserResponse, err error) {
	var user repository.User
	resp = new(service.DouyinUserResponse)
	resp.StatusCode = e.SUCCESS
	user, err = user.UserShow(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	resp.User = repository.BuildUser(user)
	return resp, err
}
