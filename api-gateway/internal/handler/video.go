package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"path/filepath"
)

// 视频投稿
func ActionVideo(ginCtx *gin.Context) {
	var videoReq service.DouyinPublishActionRequest
	PanicIfVideoError(ginCtx.ShouldBindWith(&videoReq, binding.Form))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["video"].(service.VideoServiceClient)
	claims, _ := util.ParseToken(videoReq.Token)
	videoReq.UserId = claims.UserID

	videoResp, err := videoService.ActionVideo(context.Background(), &videoReq)
	PanicIfVideoError(err)
	videoId := videoResp.VideoId
	data, err := ginCtx.FormFile("data")
	if err != nil {
		videoResp.StatusCode = e.ErrorUploadFail
	}
	//filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d.%s", videoId, "mp4")
	saveFile := filepath.Join("./public/video/", finalName)
	err = ginCtx.SaveUploadedFile(data, saveFile)
	if err != nil {
		videoResp.StatusCode = e.ErrorDownloadFail
	}
	r := res.Response{
		StatusCode: videoResp.GetStatusCode(),
		Data:       nil,
		StatusMsg:  e.GetMsg(videoResp.GetStatusCode()),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func VideoList(ginCtx *gin.Context) {
	var videoReq service.DouyinPublishListRequest
	PanicIfVideoError(ginCtx.BindQuery(&videoReq))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["video"].(service.VideoServiceClient)
	videoResp, err := videoService.VideoList(context.Background(), &videoReq)
	PanicIfVideoError(err)
	r := res.VideoResponse{
		StatusCode: videoResp.GetStatusCode(),
		Data:       videoResp.GetVideoList(),
		StatusMsg:  e.GetMsg(videoResp.GetStatusCode()),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func VideoFeed(ginCtx *gin.Context) {
	var videoReq service.DouyinFeedRequest
	PanicIfVideoError(ginCtx.BindQuery(&videoReq))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["video"].(service.VideoServiceClient)
	videoResp, err := videoService.VideoFeed(context.Background(), &videoReq)
	PanicIfVideoError(err)
	r := res.FeedResponse{
		StatusCode: videoResp.GetStatusCode(),
		Data:       videoResp.GetVideoList(),
		StatusMsg:  e.GetMsg(videoResp.GetStatusCode()),
		NextTime:   videoResp.NextTime,
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}
