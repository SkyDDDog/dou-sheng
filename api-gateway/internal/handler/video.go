package handler

import (
	"api-gateway/internal/service"
	"api-gateway/middleware"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
)

// 视频投稿
func ActionVideo(ginCtx *gin.Context) {
	var videoReq service.DouyinPublishActionRequest
	PanicIfVideoError(ginCtx.ShouldBindWith(&videoReq, binding.Form))
	// 从gin.Key中取出服务实例
	videoService := ginCtx.Keys["video"].(service.VideoServiceClient)
	cosService := ginCtx.Keys["cos"].(service.CosServiceClient)
	claims, _ := util.ParseToken(videoReq.Token)
	videoReq.UserId = claims.UserID

	videoId := int64(middleware.GenID())
	data, err := ginCtx.FormFile("data")
	//fmt.Println("data: ", data)
	PanicIfVideoError(err)
	dataContent, _ := data.Open()
	buff := new(bytes.Buffer)
	//fmt.Println("buff: ", buff)
	_, err = io.Copy(buff, dataContent)
	PanicIfVideoError(err)
	cosReq := &service.CosUploadRequest{
		Id:   videoId,
		Data: buff.Bytes(),
	}
	cosResp, err := cosService.UploadFile(context.Background(), cosReq)
	PanicIfCosError(err)

	videoReq.VideoId = videoId
	videoReq.VideoUrl = cosResp.GetVideoUrl()
	videoReq.CoverUrl = cosResp.GetCoverUrl()
	videoResp, err := videoService.ActionVideo(context.Background(), &videoReq)
	PanicIfVideoError(err)

	//filename := filepath.Base(data.Filename)
	//// 保存文件至本地
	//finalName := fmt.Sprintf("%d.%s", videoId, "mp4")
	//saveFile := filepath.Join("./public/video/", finalName)
	//err = ginCtx.SaveUploadedFile(data, saveFile)
	//if err != nil {
	//	videoResp.StatusCode = e.ErrorDownloadFail
	//}
	//// 生成缩略图
	//err = util.GetSnapshot(videoId, 1)
	//if err != nil {
	//	videoResp.StatusCode = e.ErrorSnapshotFail
	//}
	r := res.Response{
		StatusCode: videoResp.GetStatusCode(),
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
	interactService := ginCtx.Keys["interact"].(service.InteractServiceClient)
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(videoReq.Token)
	videoResp, err := videoService.VideoList(context.Background(), &videoReq)
	PanicIfVideoError(err)
	PanicIfVideoError(err)
	for i, _ := range videoResp.VideoList {
		interactReq := &service.VideoId_Request{
			VideoId:     videoResp.VideoList[i].Id,
			RequesterId: claims.UserID,
		}
		videoInfo, err := interactService.VideoInteractInfoById(context.Background(), interactReq)
		PanicIfInteractError(err)
		videoResp.VideoList[i].FavoriteCount = videoInfo.FavoriteCount
		videoResp.VideoList[i].CommentCount = videoInfo.CommentCount
		videoResp.VideoList[i].IsFavorite = videoInfo.IsFavorite

		relationReq := &service.UserId_Request{
			UserId:      videoResp.VideoList[i].Author.Id,
			RequesterId: claims.UserID,
		}
		relationInfo, err := relationService.UserRelationInfoById(context.Background(), relationReq)
		PanicIfRelationError(err)
		videoResp.VideoList[i].Author.FollowCount = relationInfo.FollowCount
		videoResp.VideoList[i].Author.FollowerCount = relationInfo.FollowerCount
		videoResp.VideoList[i].Author.IsFollow = relationInfo.IsFollow
	}

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
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	interactService := ginCtx.Keys["interact"].(service.InteractServiceClient)
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(videoReq.Token)
	videoResp, err := videoService.VideoFeed(context.Background(), &videoReq)
	PanicIfVideoError(err)
	for i, _ := range videoResp.VideoList {
		interactReq := &service.VideoId_Request{
			VideoId:     videoResp.VideoList[i].Id,
			RequesterId: claims.UserID,
		}
		videoInfo, err := interactService.VideoInteractInfoById(context.Background(), interactReq)
		PanicIfInteractError(err)
		videoResp.VideoList[i].FavoriteCount = videoInfo.FavoriteCount
		videoResp.VideoList[i].CommentCount = videoInfo.CommentCount
		videoResp.VideoList[i].IsFavorite = videoInfo.IsFavorite
		relationReq := &service.UserId_Request{
			UserId:      videoResp.VideoList[i].Author.Id,
			RequesterId: claims.UserID,
		}
		relationInfo, err := relationService.UserRelationInfoById(context.Background(), relationReq)
		PanicIfRelationError(err)
		videoResp.VideoList[i].Author.FollowCount = relationInfo.FollowCount
		videoResp.VideoList[i].Author.FollowerCount = relationInfo.FollowerCount
		videoResp.VideoList[i].Author.IsFollow = relationInfo.IsFollow
		user, err := userService.UserById(context.Background(), relationReq)
		PanicIfUserError(err)
		videoResp.VideoList[i].Author.Name = user.Name
	}

	r := res.FeedResponse{
		StatusCode: videoResp.GetStatusCode(),
		Data:       videoResp.GetVideoList(),
		StatusMsg:  e.GetMsg(videoResp.GetStatusCode()),
		NextTime:   videoResp.NextTime,
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}
