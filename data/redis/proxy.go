package redis

import (
	Common "github.com/HDDDZ/test/chatApp/data/common"
)

type RedisService struct {
}

type UserRedisService interface {
	GetUserByToken(token string) Common.User
	GetUserByUid(uid int) Common.User
	AddUserList(requests []Common.User) error
}

type FriendRedisService interface {
	UpdateRequest(requestId int) error
	AddRequestList(requests Common.ReuqestOfAddingFriend) error
	GetAllRequests(uid int) []Common.ReuqestOfAddingFriend
	GetAllFriends(uid int) []Common.User
	GetAllFriendsUid(uid int) []int
	DeleteFriend(uid int, friendUid int) error
	AddFriends(uid int, friendUids []int)
	QueryRequestById(uid int) Common.ReuqestOfAddingFriend
	QueryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend
}

type GroupRedisService interface {
	UpdateGroupInfo(group Common.Group) error
	GetGroupByGid(gid int) Common.Group
	GetAllGroupsByUid(uid int) []Common.Group
	AddGroups(group []Common.Group) error
	AddMembers(members []Common.GroupMember) error
	RemoveMembers(gid int, uids ...int) error
	GetAllMembersInfo(gid int) []Common.GroupMember
	GetAllMembersUid(gid int) []int
	GetMemberInfo(gid int, uid int) Common.GroupMember
}

type MsgRedisService interface {
	PushMessages(message []Common.DBMessage) error
	PushWaitMessages(message []Common.DBMessage) error
	GetWaitMessages(uid int) []Common.DBMessage
	QueryMessagesByMsgIds(msgIds string) []Common.DBMessage
}

func (service *RedisService) GetUserByToken(token string) Common.User {
	panic("not implement")
}

func (service *RedisService) GetUserByUid(uid int) Common.User {
	panic("not implement")

}
func (service *RedisService) AddUserList(users []Common.User) error {
	panic("not implement")

}
func (service *RedisService) UpdateRequest(requestId int) error {
	panic("not implement")

}
func (service *RedisService) AddRequestList(requests Common.ReuqestOfAddingFriend) error {
	panic("not implement")

}
func (service *RedisService) GetAllRequests(uid int) []Common.ReuqestOfAddingFriend {
	panic("not implement")

}
func (service *RedisService) GetAllFriends(uid int) []Common.User {
	panic("not implement")

}
func (service *RedisService) GetAllFriendsUid(uid int) []int {
	panic("not implement")

}
func (service *RedisService) DeleteFriend(uid int, friendUid int) error {
	panic("not implement")

}
func (service *RedisService) AddFriends(uid int, friendUids []int) {
	panic("not implement")

}
func (service *RedisService) QueryRequestById(uid int) Common.ReuqestOfAddingFriend {
	panic("not implement")

}
func (service *RedisService) QueryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	panic("not implement")

}
func (service *RedisService) UpdateGroupInfo(group Common.Group) error {
	panic("not implement")

}
func (service *RedisService) GetGroupByGid(gid int) Common.Group {
	panic("not implement")

}
func (service *RedisService) GetAllGroupsByUid(uid int) []Common.Group {
	panic("not implement")

}
func (service *RedisService) AddGroups(group []Common.Group) error {
	panic("not implement")

}
func (service *RedisService) AddMembers(members []Common.GroupMember) error {
	panic("not implement")

}
func (service *RedisService) RemoveMembers(gid int, uids ...int) error {
	panic("not implement")

}
func (service *RedisService) GetAllMembersInfo(gid int) []Common.GroupMember {
	panic("not implement")

}
func (service *RedisService) GetAllMembersUid(gid int) []int {
	panic("not implement")

}
func (service *RedisService) GetMemberInfo(gid int, uid int) Common.GroupMember {
	panic("not implement")

}
func (service *RedisService) PushMessages(message []Common.DBMessage) (int64, error) {
	panic("not implement")

}
func (service *RedisService) QueryMessagesByMsgIds(msgIds string) []Common.DBMessage {
	panic("not implement")

}
