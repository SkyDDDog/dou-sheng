package handler

import (
	"context"
	"errors"
	"fmt"
	"message/internal/repository"
	"message/internal/service"
	"message/pkg/e"
)

type MessageService struct {
	service.UnimplementedMessageServiceServer
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (*MessageService) MessageChat(ctx context.Context, req *service.DouyinMessageChatRequest) (resp *service.DouyinMessageChatResponse, err error) {
	var message repository.Message
	resp = new(service.DouyinMessageChatResponse)
	resp.StatusCode = e.SUCCESS
	mList, err := message.UserChat(req)
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	fmt.Println(mList)
	resp.MessageList = repository.BuildMessages(mList)
	fmt.Println(resp.MessageList)
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}

func (*MessageService) MessageAction(ctx context.Context, req *service.DouyinMessageActionRequest) (resp *service.DouyinMessageActionResponse, err error) {
	var message repository.Message
	resp = new(service.DouyinMessageActionResponse)
	resp.StatusCode = e.SUCCESS
	if req.ActionType == 1 {
		err = message.CreateMessage(req)
	} else {
		err = errors.New("action_type参数错误")
	}
	if err != nil {
		resp.StatusCode = e.ERROR
	}
	resp.StatusMsg = e.GetMsg(resp.GetStatusCode())
	return resp, err
}
