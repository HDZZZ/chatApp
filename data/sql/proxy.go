package db

import (
	"errors"
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

type SQLService struct {
}

type UserSQLService interface {
	GetUserByToken(token string) Common.User
	GetUserByUid(uid int) Common.User
	GetUserByUids(uids string) []Common.User
	GetUserByUsername(username string) []Common.User
	AddUser(username string, pwd string, token string) (Common.User, error)
}

type FriendSQLService interface {
	SendRequest(sendUid int, receiverUid int, msg string) (int64, error)
	AgreeRequest(requestId int) error
	RefuseRequest(requestId int) error
	UpdateRequestState(uid_1 int, uid_2 int, state Common.RequestState) error
	UpdateRequestStateByRId(requestId int, state Common.RequestState) error
	GetAllRequests(uid int) []Common.ReuqestOfAddingFriend
	GetAllFriends(uid int) []Common.User
	GetAllFriendsUid(uid int) []int
	DeleteFriend(uid int, friendUid int) error
	QueryRequestById(uid int) Common.ReuqestOfAddingFriend
	QueryRequestBySuidAndRuid(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend
	QueryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend
}

type GroupSQLService interface {
	Create(ownerUid int, groupName string, memberUids ...int) (int64, error)
	TransferOwner(gid int, newOwnerId int, OlderOwnerId int) error
	LeaveGroup(gid int, uid int) error
	UpdateGroupInfo(gid int, groupName string, description string) error
	GetGroupByGid(gid int) Common.Group
	GetAllGroupsByUid(uid int) []Common.Group
	AddMember(gid int, uids ...int) error
	RemoveMember(gid int, uids ...int) error
	GetAllMembersInfo(gid int) []Common.GroupMember
	GetAllMembersUid(gid int) []int
	GetMemberInfo(gid int, uid int) Common.GroupMember
	UpdateMemberInfo(gid int, uid int, alias string, identity Common.MemberIdentity) error
}

type MsgSQLService interface {
	PushMessage(message Common.DBMessage) (int64, error)
	QueryMessagesByUid(uid int) []Common.DBMessage
	QueryMessagesByMsgIds(msgIds string) []Common.DBMessage
}

func (service *SQLService) GetUserByToken(token string) Common.User {
	users := _queryUserByAny(_user_where_token, token)
	if len(users) == 0 {
		return Common.User{}
	}
	return users[0]
}

func (service *SQLService) GetUserByUid(uid int) Common.User {
	users := _queryUserByAny(_user_where_uid, uid)
	if len(users) == 0 {
		return Common.User{}
	}
	return users[0]
}
func (service *SQLService) GetUserByUids(uids string) []Common.User {
	users := _queryUserByAny(_user_where_uid, uids)
	if len(users) == 0 {
		return []Common.User{}
	}
	return users
}

func (service *SQLService) GetUserByUsername(username string) []Common.User {
	users := _queryUserByAny(_user_where_user_name, username)
	return users
}

func (service *SQLService) AddUser(username string, pwd string, token string) (Common.User, error) {
	user, err := addUser(username, pwd, token)
	return user, err
}

func (service *SQLService) SendRequest(sendUid int, receiverUid int, msg string) (int64, error) {
	return sendRequest(sendUid, receiverUid, msg)
}

/**	同意请求(更改该条请求数据, 添加好友记录)
*	error:requstId is correct
 */
func (service *SQLService) AgreeRequest(requestId int) error {

	err := agreeRequest(requestId)
	if err != nil {
		return err
	}
	request := queryRequestById(requestId)
	if request == (Common.ReuqestOfAddingFriend{}) {
		return errors.New("requstId is correct")
	}
	err = friendWithSomeone(request.Sender_id, request.Receiver_id)
	if err != nil {
		return err
	}
	return nil
}

func (service *SQLService) RefuseRequest(requestId int) error {
	return refuseRequest(requestId)
}

func (service *SQLService) UpdateRequestState(uid_1 int, uid_2 int, state Common.RequestState) error {
	return makeRequestState(uid_1, uid_2, state)
}
func (service *SQLService) UpdateRequestStateByRId(requestId int, state Common.RequestState) error {
	panic("not implement")
}

func (service *SQLService) GetAllRequests(uid int) []Common.ReuqestOfAddingFriend {
	return getAllRequestOfSomebody(uid)
}

func (service *SQLService) GetAllFriends(uid int) []Common.User {
	return getAllFriends(uid)
}

func (service *SQLService) GetAllFriendsUid(uid int) []int {
	return getAllFriendsUid(uid)
}

func (service *SQLService) DeleteFriend(uid int, friendUid int) error {
	return deleteFriend(uid, friendUid)
}

func (service *SQLService) QueryRequestById(uid int) Common.ReuqestOfAddingFriend {
	return queryRequestById(uid)
}
func (service *SQLService) QueryRequestBySuidAndRuid(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	return queryRequestBySuidAndRuid(uid_1, uid_2)
}
func (service *SQLService) QueryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	return queryRequestByUids(uid_1, uid_2)
}

/*
*
  - 创建群聊
    return int: 群id, 如果失败,这是0
    err
*/
func (service *SQLService) Create(ownerUid int, groupName string, memberUids ...int) (int64, error) {
	id, err := create(ownerUid, groupName, memberUids...)
	if err != nil {
		return 0, err
	}
	var insertValues = make(map[int]Common.MemberIdentity, len(memberUids)+1)
	for _, memberUid := range memberUids {
		insertValues[memberUid] = Common.Member
	}
	insertValues[ownerUid] = Common.Owner
	err = addMember(int(id), insertValues)

	if err != nil {
		fmt.Println("insert into group_members error", err)
		return 0, err
	}
	return id, nil
}
func (service *SQLService) TransferOwner(gid int, newOwnerId int, OlderOwnerId int) error {
	err := transferGroupOwner(gid, newOwnerId)
	if err != nil {
		return err
	}
	err = service.UpdateMemberInfo(gid, newOwnerId, "", Common.Owner)
	if err != nil {
		return err
	}
	err = service.UpdateMemberInfo(gid, OlderOwnerId, "", Common.Member)
	if err != nil {
		return err
	}
	return nil
}
func (service *SQLService) LeaveGroup(gid int, uid int) error {
	return leaveGroup(gid, uid)
}
func (service *SQLService) UpdateGroupInfo(gid int, groupName string, description string) error {
	return updateGroupInfo(gid, groupName, description)
}
func (service *SQLService) GetGroupByGid(gid int) Common.Group {
	return getGroupByGid(gid)
}
func (service *SQLService) GetAllGroupsByUid(uid int) []Common.Group {
	return getAllGroupsByUid(uid)
}
func (service *SQLService) AddMember(gid int, uids ...int) error {
	var insertValues = make(map[int]Common.MemberIdentity, len(uids))
	for _, memberUid := range uids {
		insertValues[memberUid] = Common.Member
	}
	return addMember(gid, insertValues)
}
func (service *SQLService) RemoveMember(gid int, uids ...int) error {
	return removeMember(gid, uids...)
}
func (service *SQLService) GetAllMembersInfo(gid int) []Common.GroupMember {
	return getAllMembersInfo(gid)
}

func (service *SQLService) GetAllMembersUid(gid int) []int {
	return getAllMembersUid(gid)
}
func (service *SQLService) GetMemberInfo(gid int, uid int) Common.GroupMember {
	return getMemberInfo(gid, uid)
}
func (service *SQLService) UpdateMemberInfo(gid int, uid int, alias string, identity Common.MemberIdentity) error {
	return updateMemberInfo(gid, uid, alias, identity)
}
func (service *SQLService) PushMessage(message Common.DBMessage) (int64, error) {
	return pushMessage(message)
}
func (service *SQLService) QueryMessagesByUid(uid int) []Common.DBMessage {
	return queryMessagesByUid(uid)
}
func (service *SQLService) QueryMessagesByMsgIds(msgIds string) []Common.DBMessage {
	return queryMessagesByMsgIds(msgIds)
}
