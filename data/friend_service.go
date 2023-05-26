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
	log("SendRequest,mysql,result=", result, "err=", err)
	go synchroniseReuqest(int(result))
	return result, err
}

func AgreeRequest(requestId int) error {
	err := friendDBService.AgreeRequest(requestId)
	log("AgreeRequest,mysql,err=", err)
	go synchroniseRequestAndFriend(requestId)
	return err
}

func RefuseRequest(requestId int) error {
	err := friendDBService.RefuseRequest(requestId)
	log("RefuseRequest,mysql,err=", err)
	go synchroniseReuqest(requestId)
	return err
}

func MakeRequestState(uid_1 int, uid_2 int, state Common.RequestState) error {
	err := friendDBService.UpdateRequestState(uid_1, uid_2, state)
	log("MakeRequestState,mysql,err=", err)
	asyncExcute(func() {
		request := friendDBService.QueryRequestBySuidAndRuid(uid_1, uid_2)
		log("MakeRequestState,mysql,request=", request)
		err := friendRedisService.RefreshRequests(request)
		log("MakeRequestState,redis,err=", err)
	})
	return err
}

func GetAllRequestOfSomebody(uid int) []Common.ReuqestOfAddingFriend {
	requests := friendRedisService.GetRequestsByUid(fmt.Sprint(uid))
	log("GetAllRequestOfSomebody,redis,requests=", requests)
	if len(requests) == 0 {
		requests = friendDBService.GetAllRequests(uid)
		log("GetAllRequestOfSomebody,mysql,requests=", requests)
		err := friendRedisService.SpecifyUserRefreshRequests(fmt.Sprint(uid), requests)
		log("GetAllRequestOfSomebody,redis,SpecifyUserRefreshRequests=", err)
	}
	return requests
}

func GetAllFriends(uid int) []Common.User {
	uids, _ := friendRedisService.GetFriendsUidByUid(uid)
	log("GetAllFriends,redis,uids=", uids)
	if len(uids) == 0 {
		users := friendDBService.GetAllFriends(uid)
		log("GetAllFriends,mysql,users=", users)
		err := userRedisService.AddUserList(users...)
		log("GetAllFriends,redis,err=", err)
		return users
	}
	users := QueryUsersByUids(uids)
	return users
}

func GetAllFriendsUid(uid int) []int {
	uids, _ := friendRedisService.GetFriendsUidByUid(uid)
	log("GetAllFriendsUid,redis,uids=", uids)

	if len(uids) == 0 {
		users := friendDBService.GetAllFriendsUid(uid)
		log("GetAllFriendsUid,mysql,users=", users)

		err := friendRedisService.UpdateFriendsUidByUid(uid, users)
		log("GetAllFriendsUid,redis,err=", err)

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
	log("QueryRequestById,params,requestId=", requestId)
	requests := friendRedisService.GetRequestsByRid(fmt.Sprint(requestId))
	log("QueryRequestById,redis,requests=", requests)

	if len(requests) == 0 {
		request := friendDBService.QueryRequestById(requestId)
		log("QueryRequestById,mysql,request=", request)

		err := friendRedisService.RefreshRequests(request)
		log("QueryRequestById,redis,err=", err)
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
	request := synchroniseReuqest(requestId)
	synchroniseFriends(request.Sender_id, request.Receiver_id)
}

func synchroniseFriends(uids ...int) {
	for _, uid := range uids {
		users := friendDBService.GetAllFriendsUid(uid)
		log("synchroniseFriends,users=", users)
		err := friendRedisService.UpdateFriendsUidByUid(uid, users)
		if err != nil {
			log("synchroniseFriends,err=", err)
		}
	}
}

func synchroniseReuqest(requestId int) Common.ReuqestOfAddingFriend {
	log("synchroniseReuqest,params,requestId=", requestId)
	request := friendDBService.QueryRequestById(requestId)
	log("synchroniseReuqest,mysql,request=", request)
	err := friendRedisService.RefreshRequests(request)
	log("synchroniseReuqest,redis,err=", err)
	return request
}
