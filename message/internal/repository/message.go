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
		Where("(from_id = ? and to_id = ?) or (from_id = ? and to_id = ?)", req.FromUserId, req.ToUserId, req.ToUserId, req.GetFromUserId()).
		Order("create_date asc").
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
		ToUserId:   item.ToId,
		FromUserId: item.FromId,
		Content:    item.Content,
		CreateTime: item.Create_date.Unix(),
	}
}

func BuildMessages(item []Message) (mList []*service.Message) {
	for _, message := range item {
		mList = append(mList, BuildMessage(message))
	}
	return mList
}
