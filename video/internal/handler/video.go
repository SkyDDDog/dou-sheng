package handler

import (
	"context"
	"video/internal/repository"
	"video/internal/service"
	"video/pkg/e"
)

type VideoService struct {
	service.UnimplementedVideoServiceServer
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (*VideoService) ActionVideo(ctx context.Context, req *service.DouyinPublishActionRequest) (resp *service.DouyinPublishActionResponse, err error) {
	var video repository.Video
	resp = new(service.DouyinPublishActionResponse)
	resp.StatusCode = e.SUCCESS
	video, err = video.CreateVideo(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.VideoId = video.VideoId
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*VideoService) VideoList(ctx context.Context, req *service.DouyinPublishListRequest) (resp *service.DouyinPublishListResponse, err error) {
	var video repository.Video
	resp = new(service.DouyinPublishListResponse)
	resp.StatusCode = e.SUCCESS
	videoList, err := video.VideoShow(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.VideoList = repository.BuildVideos(videoList)
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*VideoService) VideoFeed(ctx context.Context, req *service.DouyinFeedRequest) (resp *service.DouyinFeedResponse, err error) {
	var video repository.Video
	resp = new(service.DouyinFeedResponse)
	resp.StatusCode = e.SUCCESS
	videoList, err := video.VideoFeed(req)
	resp.VideoList = repository.BuildVideos(videoList)
	if len(videoList) != 0 {
		resp.NextTime = videoList[0].Create_date.Unix()
	}
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*VideoService) VideoById(ctx context.Context, req *service.VideoId_Request) (resp *service.Video, err error) {
	var video repository.Video
	video, err = video.VideoById(req.VideoId)
	if err != nil {
		return nil, err
	}
	resp = repository.BuildVideo(video)
	return resp, nil
}
