package handler

import (
	"context"
	"errors"
	"relation/internal/repository"
	"relation/internal/service"
	"relation/pkg/e"
)

type RelationService struct {
	service.UnimplementedRelationServiceServer
}

func NewRelationService() *RelationService {
	return &RelationService{}
}

func (*RelationService) RelationAction(ctx context.Context, req *service.DouyinRelationActionRequest) (resp *service.DouyinRelationActionResponse, err error) {
	var uf repository.UserFollow
	resp = new(service.DouyinRelationActionResponse)
	resp.StatusCode = e.SUCCESS
	if req.ActionType == 1 {
		err = uf.CreateRelation(req)
	} else if req.ActionType == 2 {
		err = uf.DeleteRelation(req)
	} else {
		err = errors.New("action_type 错误")
	}
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*RelationService) FollowList(ctx context.Context, req *service.DouyinRelationFollowListRequest) (resp *service.DouyinRelationFollowListResponse, err error) {
	var uf repository.UserFollow
	resp = new(service.DouyinRelationFollowListResponse)
	resp.StatusCode = e.SUCCESS
	ufList, err := uf.FollowList(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.UserList = repository.BuildUsers(ufList)
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*RelationService) FollowerList(ctx context.Context, req *service.DouyinRelationFollowerListRequest) (resp *service.DouyinRelationFollowerListResponse, err error) {
	var uf repository.UserFollow
	resp = new(service.DouyinRelationFollowerListResponse)
	resp.StatusCode = e.SUCCESS
	ufList, err := uf.FollowerList(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.UserList = repository.BuildUsers(ufList)
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*RelationService) FriendList(ctx context.Context, req *service.DouyinRelationFriendListRequest) (resp *service.DouyinRelationFriendListResponse, err error) {
	var uf repository.UserFollow
	resp = new(service.DouyinRelationFriendListResponse)
	resp.StatusCode = e.SUCCESS
	ufList, err := uf.FriendList(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.UserList = repository.BuildUsers(ufList)
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*RelationService) UserRelationInfoById(ctx context.Context, req *service.UserId_Request) (resp *service.User, err error) {
	uf := repository.UserFollow{
		FromUserId: req.RequesterId,
		ToUserId:   req.UserId,
	}
	resp = repository.BuildUser(uf)
	return resp, nil
}
