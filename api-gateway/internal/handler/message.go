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

func UserChat(ginCtx *gin.Context) {
	var msgReq service.DouyinMessageChatRequest
	PanicIfMessageError(ginCtx.ShouldBindWith(&msgReq, binding.Query))
	// 从gin.Key中取出服务实例
	messageService := ginCtx.Keys["message"].(service.MessageServiceClient)
	claims, _ := util.ParseToken(msgReq.Token)
	msgReq.FromUserId = claims.UserID
	msgResp, err := messageService.MessageChat(context.Background(), &msgReq)
	PanicIfMessageError(err)
	r := res.MessageResponse{
		StatusCode: msgResp.StatusCode,
		Data:       msgResp.MessageList,
		StatusMsg:  e.GetMsg(msgResp.GetStatusCode()),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}

func MessageAction(ginCtx *gin.Context) {
	var msgReq service.DouyinMessageActionRequest
	PanicIfMessageError(ginCtx.ShouldBindWith(&msgReq, binding.Query))
	// 从gin.Key中取出服务实例
	messageService := ginCtx.Keys["message"].(service.MessageServiceClient)
	claims, _ := util.ParseToken(msgReq.Token)
	msgReq.FromUserId = claims.UserID
	msgResp, err := messageService.MessageAction(context.Background(), &msgReq)
	PanicIfMessageError(err)
	r := res.Response{
		StatusCode: msgResp.StatusCode,
		StatusMsg:  e.GetMsg(msgResp.GetStatusCode()),
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}
