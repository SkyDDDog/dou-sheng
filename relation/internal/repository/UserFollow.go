package repository

import "relation/internal/service"

type UserFollow struct {
	Id         int64 `gorm:"primarykey"`
	FromUserId int64
	ToUserId   int64
	Model
}

func (*UserFollow) CreateRelation(req *service.DouyinRelationActionRequest) error {
	uf := &UserFollow{
		FromUserId: req.UserId,
		ToUserId:   req.ToUserId,
		Model:      Model{},
	}
	return DB.Create(&uf).Error
}

func (*UserFollow) DeleteRelation(req *service.DouyinRelationActionRequest) error {
	return DB.Model(UserFollow{}).
		Unscoped().
		Where("from_user_id = ? and to_user_id = ?", req.UserId, req.ToUserId).
		Delete(&UserFollow{}).Error
}

func (*UserFollow) FollowList(req *service.DouyinRelationFollowListRequest) (ufList []UserFollow, err error) {
	err = DB.Model(&UserFollow{}).
		Where("from_user_id = ?", req.UserId).
		Find(&ufList).Error
	return ufList, err
}

func (*UserFollow) FollowerList(req *service.DouyinRelationFollowerListRequest) (ufList []UserFollow, err error) {
	err = DB.Model(&UserFollow{}).
		Where("to_user_id = ?", req.UserId).
		Find(&ufList).Error
	return ufList, err
}

func (*UserFollow) FriendList(req *service.DouyinRelationFriendListRequest) (ufList []UserFollow, err error) {
	// 查找我关注的人
	err = DB.Model(&UserFollow{}).
		Where("from_user_id = ?", req.UserId).
		Find(&ufList).Error
	uf := &UserFollow{
		FromUserId: 0,
		ToUserId:   req.UserId,
	}
	for i := 0; i < len(ufList); i++ {
		uf.FromUserId = ufList[i].ToUserId
		// 其中未关注我的除外
		if !uf.IsFollow() {
			// 删除切片元素
			ufList = append(ufList[:i], ufList[i+1:]...)
			i--
		}
	}
	return ufList, err
}

func BuildUser(item UserFollow) *service.User {
	return &service.User{
		Id:            item.ToUserId,
		Name:          "",
		FollowCount:   item.FollowerCount(),
		FollowerCount: item.FollowCount(),
		IsFollow:      item.IsFollow(),
	}
}

func BuildUsers(item []UserFollow) (userList []*service.User) {
	for _, i := range item {
		userList = append(userList, BuildUser(i))
	}
	return userList
}

func (uf *UserFollow) FollowCount() int64 {
	var count int64
	DB.Model(&UserFollow{}).Where("from_user_id = ?", uf.FromUserId).Count(&count)
	return count
}

func (uf *UserFollow) FollowerCount() int64 {
	var count int64
	DB.Model(&UserFollow{}).Where("to_user_id = ?", uf.ToUserId).Count(&count)
	return count
}

func (uf *UserFollow) IsFollow() bool {
	var count int64
	DB.Model(&UserFollow{}).
		Where("from_user_id = ? and to_user_id = ?", uf.FromUserId, uf.ToUserId).
		Count(&count)
	return 0 < count
}
