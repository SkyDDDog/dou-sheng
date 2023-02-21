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

func FavoriteAction(ginCtx *gin.Context) {
	var ufReq service.DouyinFavoriteActionRequest
	PanicIfInteractError(ginCtx.ShouldBindWith(&ufReq, binding.Query))
	// 从gin.Key中取出服务实例
	interactService := ginCtx.Keys["interact"].(service.InteractServiceClient)
	claims, _ := util.ParseToken(ufReq.Token)
	ufReq.UserId = claims.UserID

	ufResp, err := interactService.FavoriteAction(context.Background(), &ufReq)
	PanicIfInteractError(err)
	r := res.Response{
		StatusCode: ufResp.StatusCode,
		StatusMsg:  e.GetMsg(ufResp.GetStatusCode()),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func FavoriteList(ginCtx *gin.Context) {
	var flReq service.DouyinFavoriteListRequest
	PanicIfInteractError(ginCtx.ShouldBindWith(&flReq, binding.Query))
	// 从gin.Key中取出服务实例
	interactService := ginCtx.Keys["interact"].(service.InteractServiceClient)
	videoService := ginCtx.Keys["video"].(service.VideoServiceClient)
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(flReq.Token)
	resp, err := interactService.FavoriteList(context.Background(), &flReq)
	PanicIfInteractError(err)
	for i, _ := range resp.VideoList {
		videoReq := &service.VideoId_Request{
			VideoId:     resp.VideoList[i].Id,
			RequesterId: claims.UserID,
		}
		video, err := videoService.VideoById(context.Background(), videoReq)
		PanicIfVideoError(err)

		var userReq service.UserIdRequest
		userReq.UserId = resp.VideoList[i].Author.Id
		user, err := userService.UserById(context.Background(), &userReq)
		PanicIfUserError(err)
		video.Author = user
		videoInfo, err := interactService.VideoInteractInfoById(context.Background(), videoReq)
		PanicIfInteractError(err)
		video.FavoriteCount = videoInfo.FavoriteCount
		video.CommentCount = videoInfo.CommentCount
		video.IsFavorite = videoInfo.IsFavorite

		relationReq := &service.UserIdRequest{
			UserId:      video.Author.Id,
			RequesterId: claims.UserID,
		}
		relationInfo, err := relationService.UserRelationInfoById(context.Background(), relationReq)
		PanicIfRelationError(err)
		video.Author.FollowCount = relationInfo.FollowCount
		video.Author.FollowerCount = relationInfo.FollowerCount
		video.Author.IsFollow = relationInfo.IsFollow

		resp.VideoList[i] = video
	}
	r := res.FavoriteResponse{
		StatusCode: resp.StatusCode,
		Data:       resp.VideoList,
		StatusMsg:  e.GetMsg(resp.StatusCode),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CommentAction(ginCtx *gin.Context) {
	var caReq service.DouyinCommentActionRequest
	PanicIfInteractError(ginCtx.ShouldBindWith(&caReq, binding.Query))
	// 从gin.Key中取出服务实例
	interactService := ginCtx.Keys["interact"].(service.InteractServiceClient)
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(caReq.Token)
	caReq.UserId = claims.UserID
	resp, err := interactService.CommentAction(context.Background(), &caReq)
	PanicIfInteractError(err)
	userReq := &service.UserIdRequest{UserId: caReq.UserId}
	user, err := userService.UserById(context.Background(), userReq)
	PanicIfUserError(err)

	relationReq := &service.UserIdRequest{
		UserId:      user.Id,
		RequesterId: claims.UserID,
	}
	relationInfo, err := relationService.UserRelationInfoById(context.Background(), relationReq)
	PanicIfRelationError(err)
	user.FollowCount = relationInfo.FollowCount
	user.FollowerCount = relationInfo.FollowerCount
	user.IsFollow = relationInfo.IsFollow

	resp.Comment.User = user

	r := res.CommentResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  e.GetMsg(resp.GetStatusCode()),
		Error:      "",
		Data:       resp.Comment,
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CommentList(ginCtx *gin.Context) {
	var clReq service.DouyinCommentListRequest
	PanicIfInteractError(ginCtx.ShouldBindWith(&clReq, binding.Query))
	// 从gin.Key中取出服务实例
	interactService := ginCtx.Keys["interact"].(service.InteractServiceClient)
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	relationService := ginCtx.Keys["relation"].(service.RelationServiceClient)
	claims, _ := util.ParseToken(clReq.Token)
	resp, err := interactService.CommentList(context.Background(), &clReq)
	PanicIfInteractError(err)
	for i, _ := range resp.CommentList {
		userReq := &service.UserIdRequest{UserId: resp.CommentList[i].User.Id}
		user, err := userService.UserById(context.Background(), userReq)
		PanicIfUserError(err)

		relationReq := &service.UserIdRequest{
			UserId:      user.Id,
			RequesterId: claims.UserID,
		}
		relationInfo, err := relationService.UserRelationInfoById(context.Background(), relationReq)
		PanicIfRelationError(err)
		user.FollowCount = relationInfo.FollowCount
		user.FollowerCount = relationInfo.FollowerCount
		user.IsFollow = relationInfo.IsFollow

		resp.CommentList[i].User = user
	}
	r := res.CommentListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  e.GetMsg(resp.GetStatusCode()),
		Error:      "",
		Data:       resp.CommentList,
	}
	ginCtx.JSON(http.StatusOK, r)
}
