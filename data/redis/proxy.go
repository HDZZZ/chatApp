package redis

import (
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

type RedisService struct {
}

type UserRedisService interface {
	GetUserByToken(token string) Common.User
	GetUserByUid(uid string) Common.User
	AddUserList(requests ...Common.User) error
}

type FriendRedisService interface {
	GetRequestsByUid(uid string) []Common.ReuqestOfAddingFriend
	GetRequestsByRid(rid string) []Common.ReuqestOfAddingFriend
	RefreshRequests(requests ...Common.ReuqestOfAddingFriend) error
	SpecifyUserRefreshRequests(uid string, requests []Common.ReuqestOfAddingFriend) error
	CheckRequestsExist(requests []Common.ReuqestOfAddingFriend) []bool
	GetFriendsUidByUid(uid int) ([]string, error)
	UpdateFriendsUidByUid(uid int, friendsUid []int) error
}

type GroupRedisService interface {
	RefreshGroups(groups ...Common.Group) error
	GetGroupByGid(gid string) Common.Group
	AddGroupMembers(gid string, members ...Common.GroupMember) error
	RefreshGroupAllMembers(gid string, members ...Common.GroupMember) error
	RemoveGroupMembers(gid string, members ...Common.GroupMember) error
	UpdateGroupMemberInfo(members ...Common.GroupMember) error
	GetMember(gid string, uid string) Common.GroupMember
	GetAllMembers(gid string) []Common.GroupMember
}

type MsgRedisService interface {
	PushMessages(msgs ...Common.DBMessage) error
	PushMsgIdToWaitList(uid string, msgId int) error
	ClearWaitListOfSomebody(uid string) error
	GetAllWaitMsgIdOfSomeBody(uid string) []string
	GetMessageFromMsgID(uid string) Common.DBMessage
}

func (service *RedisService) GetUserByToken(token string) Common.User {
	return getUserByToken(token)
}

func (service *RedisService) GetUserByUid(uid string) Common.User {
	return getUserByUid(uid)

}
func (service *RedisService) AddUserList(users ...Common.User) error {
	return addUsers(users)
}

func (service *RedisService) GetRequestsByUid(uid string) []Common.ReuqestOfAddingFriend {
	return getRequestsByUid(uid)
}
func (service *RedisService) GetRequestsByRid(rid string) []Common.ReuqestOfAddingFriend {
	return getRequestsByRids(rid)
}
func (service *RedisService) RefreshRequests(requests ...Common.ReuqestOfAddingFriend) error {
	return refreshRequests(requests)
}
func (service *RedisService) SpecifyUserRefreshRequests(uid string, requests []Common.ReuqestOfAddingFriend) error {
	return specifyUserRefreshRequests(uid, requests)
}
func (service *RedisService) CheckRequestsExist(requests []Common.ReuqestOfAddingFriend) []bool {
	return checkRequestsExist(requests)
}
func (service *RedisService) GetFriendsUidByUid(uid int) ([]string, error) {
	return getFriendsUidByUid(uid)
}
func (service *RedisService) UpdateFriendsUidByUid(uid int, friendsUid []int) error {
	return updateFriendsUidByUid(uid, friendsUid)
}
func (service *RedisService) RefreshGroups(groups ...Common.Group) error {
	return refreshGroups(groups...)
}
func (service *RedisService) GetGroupByGid(gid string) Common.Group {
	return getGroupByGid(gid)
}
func (service *RedisService) AddGroupMembers(gid string, members ...Common.GroupMember) error {
	return addGroupMembers(gid, members...)
}

func (service *RedisService) RefreshGroupAllMembers(gid string, members ...Common.GroupMember) error {
	return refreshGroupAllMembers(gid, members...)
}
func (service *RedisService) RemoveGroupMembers(gid string, members ...Common.GroupMember) error {
	return removeGroupMembers(gid, members...)
}
func (service *RedisService) UpdateGroupMemberInfo(members ...Common.GroupMember) error {
	return updateGroupMemberInfo(members...)
}
func (service *RedisService) GetMember(gid string, uid string) Common.GroupMember {
	return getMember(gid, uid)
}
func (service *RedisService) GetAllMembers(gid string) []Common.GroupMember {
	return getAllMembers(gid)
}
func (service *RedisService) PushMessages(msgs ...Common.DBMessage) error {
	return pushMessages(msgs...)
}
func (service *RedisService) PushMsgIdToWaitList(uid string, msgId int) error {
	return pushMsgIdToWaitList(uid, msgId)
}
func (service *RedisService) ClearWaitListOfSomebody(uid string) error {
	return clearWaitListOfSomebody(uid)
}
func (service *RedisService) GetAllWaitMsgIdOfSomeBody(uid string) []string {
	return getAllWaitMsgIdOfSomeBody(uid)
}
func (service *RedisService) GetMessageFromMsgID(uid string) Common.DBMessage {
	return getMessageFromMsgID(uid)
}

func log(logs ...string) {
	fmt.Println("redis,", logs)
}
