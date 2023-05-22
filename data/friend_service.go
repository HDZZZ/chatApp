package data

import (
	"fmt"
	"strconv"

	Common "github.com/HDDDZ/test/chatApp/data/common"
	Redis "github.com/HDDDZ/test/chatApp/data/redis"

	DB "github.com/HDDDZ/test/chatApp/data/sql"
)

var friendDBService DB.FriendSQLService = &DB.SQLService{}
var friendRedisService Redis.FriendRedisService = &Redis.RedisService{}

func SendRequest(sendUid int, receiverUid int, msg string) (int64, error) {
	result, err := friendDBService.SendRequest(sendUid, receiverUid, msg)
	go QueryRequestById(int(result))
	return result, err
}

func AgreeRequest(requestId int) error {
	err := friendDBService.AgreeRequest(requestId)
	go synchroniseRequestAndFriend(requestId)
	return err
}

func RefuseRequest(requestId int) error {
	err := friendDBService.RefuseRequest(requestId)
	go QueryRequestById(requestId)
	return err
}

func MakeRequestState(uid_1 int, uid_2 int, state Common.RequestState) error {
	err := friendDBService.UpdateRequestState(uid_1, uid_2, state)
	go QueryRequestByUids(uid_1, uid_2)
	return err
}

func GetAllRequestOfSomebody(uid int) []Common.ReuqestOfAddingFriend {
	requests := friendRedisService.GetRequestsByUid(fmt.Sprint(uid))
	if len(requests) == 0 {
		requests = friendDBService.GetAllRequests(uid)
	}
	friendRedisService.SpecifyUserRefreshRequests(fmt.Sprint(uid), requests)
	return requests
}

func GetAllFriends(uid int) []Common.User {
	uids, _ := friendRedisService.GetFriendsUidByUid(uid)
	if len(uids) == 0 {
		users := friendDBService.GetAllFriends(uid)
		userRedisService.AddUserList(users...)
		return users
	}
	users := QueryUsersByUids(uids)
	return users
}

func GetAllFriendsUid(uid int) []int {
	uids, _ := friendRedisService.GetFriendsUidByUid(uid)
	if len(uids) == 0 {
		users := friendDBService.GetAllFriendsUid(uid)
		friendRedisService.UpdateFriendsUidByUid(uid, users)
		return users
	}
	uidsInt := Map(uids, func(uid string) int {
		uidInt, _ := strconv.Atoi(uid)
		return uidInt
	})
	return uidsInt
}

func DeleteFriend(uid int, friendUid int) error {
	err := friendDBService.DeleteFriend(uid, friendUid)
	go synchroniseFriends(uid, friendUid)
	return err
}

func QueryRequestById(requestId int) Common.ReuqestOfAddingFriend {
	requests := friendRedisService.GetRequestsByRid(fmt.Sprint(requestId))
	if len(requests) == 0 {
		request := friendDBService.QueryRequestById(requestId)
		friendRedisService.RefreshRequests(request)
		return request
	}
	return requests[0]
}

func QueryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	requests := friendDBService.QueryRequestByUids(uid_1, uid_2)
	friendRedisService.RefreshRequests(requests)
	return requests
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func synchroniseRequestAndFriend(requestId int) {
	request := QueryRequestById(requestId)
	synchroniseFriends(request.Sender_id, request.Receiver_id)
}

func synchroniseFriends(uids ...int) {
	for _, uid := range uids {
		users := friendDBService.GetAllFriendsUid(uid)
		friendRedisService.UpdateFriendsUidByUid(uid, users)
	}
}
