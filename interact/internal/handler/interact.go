package handler

import (
	"context"
	"errors"
	"interact/internal/repository"
	"interact/internal/service"
	"interact/pkg/e"
)

type InteractService struct {
	service.UnimplementedInteractServiceServer
}

func NewInteractService() *InteractService {
	return &InteractService{}
}

func (*InteractService) FavoriteAction(ctx context.Context, req *service.DouyinFavoriteActionRequest) (resp *service.DouyinFavoriteActionResponse, err error) {
	var userFavorite repository.UserFavorite
	resp = new(service.DouyinFavoriteActionResponse)
	resp.StatusCode = e.SUCCESS

	if req.ActionType == 1 {
		err = userFavorite.CreateUf(req)
	} else if req.ActionType == 2 {
		err = userFavorite.DeleteUf(req)
	} else {
		err = errors.New("action_type 错误")
	}
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*InteractService) FavoriteList(ctx context.Context, req *service.DouyinFavoriteListRequest) (resp *service.DouyinFavoriteListResponse, err error) {
	var uf repository.UserFavorite
	resp = new(service.DouyinFavoriteListResponse)
	resp.StatusCode = e.SUCCESS
	ufList, err := uf.FavoriteList(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	var videoList []*service.Video
	for _, ufItem := range ufList {
		var video service.Video
		video.Author = &service.User{Id: ufItem.UserId}
		video.Id = ufItem.VideoId
		videoList = append(videoList, &video)
	}
	resp.VideoList = videoList
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*InteractService) CommentAction(ctx context.Context, req *service.DouyinCommentActionRequest) (resp *service.DouyinCommentActionResponse, err error) {
	var vc repository.VideoComment
	resp = new(service.DouyinCommentActionResponse)
	resp.StatusCode = e.SUCCESS
	if req.ActionType == 1 {
		vc, err = vc.CreateComment(req)
	} else if req.ActionType == 2 {
		vc, err = vc.DeleteComment(req)
	} else {
		err = errors.New("action_type 错误")
	}
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.Comment = repository.BuildVideoComment(vc)
	return resp, err
}

func (*InteractService) CommentList(ctx context.Context, req *service.DouyinCommentListRequest) (resp *service.DouyinCommentListResponse, err error) {
	var vc repository.VideoComment
	resp = new(service.DouyinCommentListResponse)
	resp.StatusCode = e.SUCCESS
	vcList, err := vc.CommentList(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.CommentList = repository.BuildVideoComments(vcList)
	return resp, err
}

func (*InteractService) VideoInteractInfoById(ctx context.Context, req *service.VideoId_Request) (resp *service.Video, err error) {
	var uf repository.UserFavorite
	var vc repository.VideoComment
	resp = new(service.Video)
	resp.Id = req.VideoId
	resp.FavoriteCount = uf.FavoriteCount(req.VideoId)
	resp.CommentCount = vc.CommentCount(req.VideoId)
	resp.IsFavorite = uf.IsFavorite(req.VideoId, req.RequesterId)
	return resp, nil
}
