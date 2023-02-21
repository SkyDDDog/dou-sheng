package repository

import (
	"message/internal/service"
	"strconv"
)

type Message struct {
	MessageId int64 `gorm:"primarykey"`
	Content   string
	FromId    int64
	ToId      int64
	Model
}

func (*Message) UserChat(req *service.DouyinMessageChatRequest) (mList []Message, err error) {
	//oldMsg, err := RDB.HGetAll(BuildMessageKey(req.ToUserId)).Result()
	//oldMsgSlice := make([]string, 0)
	//for k, v := range oldMsg {
	//	if v == "1" {
	//		oldMsgSlice = append(oldMsgSlice, k)
	//	}
	//}
	//if len(oldMsgSlice) == 0 {
	//	oldMsgSlice = append(oldMsgSlice, "")
	//}

	err = DB.Model(&Message{}).
		Where("create_date > ?", req.PreMsgTime).
		//Where("(message_id NOT IN ?)", oldMsgSlice).
		//Where("((from_id = ? and to_id = ?) or (from_id = ? and to_id = ?))", req.FromUserId, req.ToUserId, req.ToUserId, req.GetFromUserId()).
		Order("create_date asc").
		Find(&mList).Error
	//m := make(map[string]interface{}, len(mList))
	//for _, msg := range mList {
	//	m[strconv.FormatInt(msg.MessageId, 10)] = true
	//}
	//RDB.HMSet(BuildMessageKey(req.ToUserId), m)
	//RDB.Expire(BuildMessageKey(req.ToUserId), 5*time.Minute)
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
		CreateTime: strconv.FormatInt(item.Create_date.Unix(), 10),
		//CreateTime: item.Create_date.Format("2006-01-02 03:04:05"),
	}
}

func BuildMessages(item []Message) (mList []*service.Message) {
	for _, message := range item {
		mList = append(mList, BuildMessage(message))
	}
	return mList
}
