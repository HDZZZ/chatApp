package data

import (
	Common "github.com/HDDDZ/test/chatApp/data/common"

	DB "github.com/HDDDZ/test/chatApp/data/sql"
)

var friendDBService DB.FriendSQLService = &DB.SQLService{}

func SendRequest(sendUid int, receiverUid int, msg string) (int64, error) {
	return friendDBService.SendRequest(sendUid, receiverUid, msg)
}

func AgreeRequest(requestId int) error {
	return friendDBService.AgreeRequest(requestId)
}

func RefuseRequest(requestId int) error {
	return friendDBService.RefuseRequest(requestId)
}

func MakeRequestState(uid_1 int, uid_2 int, state Common.RequestState) error {
	return friendDBService.UpdateRequestState(uid_1, uid_2, state)
}

func GetAllRequestOfSomebody(uid int) []Common.ReuqestOfAddingFriend {
	return friendDBService.GetAllRequests(uid)
}

func GetAllFriends(uid int) []Common.User {
	return friendDBService.GetAllFriends(uid)
}

func GetAllFriendsUid(uid int) []int {
	return friendDBService.GetAllFriendsUid(uid)
}

func DeleteFriend(uid int, friendUid int) error {
	return friendDBService.DeleteFriend(uid, friendUid)
}

func QueryRequestById(requestId int) Common.ReuqestOfAddingFriend {
	return friendDBService.QueryRequestById(requestId)
}

func QueryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	return friendDBService.QueryRequestByUids(uid_1, uid_2)
}
