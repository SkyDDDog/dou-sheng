package repository

import "message/internal/service"

type Message struct {
	MessageId int64 `gorm:"primarykey"`
	Content   string
	FromId    int64
	ToId      int64
	Model
}

func (*Message) UserChat(req *service.DouyinMessageChatRequest) (mList []Message, err error) {
	err = DB.Model(&Message{}).
		Where("from_user_id = ? and to_user_id = ?", req.FromUserId, req.ToUserId).
		Find(&mList).Error
	return mList, err
}

func (*Message) CreateMessage(req *service.DouyinMessageActionRequest) error {
	message := &Message{
		Content: req.Content,
		FromId:  req.FromUserId,
		ToId:    req.ToUserId,
		Model:   Model{},
	}
	return DB.Create(&message).Error
}

func BuildMessage(item Message) *service.Message {
	return &service.Message{
		Id:         item.MessageId,
		Content:    item.Content,
		CreateTime: item.Create_date.Format("2006-01-02 03:04:05"),
	}
}

func BuildMessages(item []Message) (mList []*service.Message) {
	for _, message := range item {
		mList = append(mList, BuildMessage(message))
	}
	return mList
}
