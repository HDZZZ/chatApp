package redis

import (
	"encoding/json"
	"fmt"
	"reflect"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const FRIEND_KEY = "friend_uid"

const REQUEST_KEY = "request"
const REQUEST_RID_KEY = "request_rid" //存储的key的 请求id
const REQUEST_UID_KEY = "request_uid"

func getRequestsByUid(uid string) []Common.ReuqestOfAddingFriend {
	rids, err := get(createReidsrKey(REQUEST_UID_KEY, uid))
	if err != nil {
		return []Common.ReuqestOfAddingFriend{}
	}

	return getRequestsByRids(rids)
}

func getRequestsByRids(rids ...string) []Common.ReuqestOfAddingFriend {
	for index, rid := range rids {
		rids[index] = createReidsrKey(REQUEST_RID_KEY, rid)
	}

	requestStreams, _ := mget(rids...)
	requests := make([]Common.ReuqestOfAddingFriend, len(requestStreams))
	var newUser Common.ReuqestOfAddingFriend
	for index, stream := range requestStreams {
		json.Unmarshal([]byte(stream.(string)), &newUser)
		requests[index] = newUser
	}
	return requests
}

/*
*
更新与添加都调用该方法,注: 如果用户对应的用户列表在redis为空的话,会忽略掉
*
*/
func refreshRequests(requests []Common.ReuqestOfAddingFriend) error {
	requestParis := make(map[string]interface{}, len(requests))
	uidParis := make(map[int]interface{}, len(requests)*2)

	isExists := checkRequestsExist(requests)

	for index, request := range requests {
		stream, _ := json.Marshal(request)
		requestParis[createReidsrKey(REQUEST_UID_KEY, request.Id)] = stream
		if isExists[index] {
			uidParis[request.Sender_id] = fmt.Sprint(uidParis[request.Sender_id]) + "|" + fmt.Sprint(request.Id)
			uidParis[request.Receiver_id] = fmt.Sprint(uidParis[request.Receiver_id]) + "|" + fmt.Sprint(request.Id)
		}
	}
	for k, v := range uidParis {
		//如果redis中没有该用户的请求列表,那么就忽略掉,(不然会导致问题)
		isExsit, _ := exists(fmt.Sprint(k))
		if isExsit == 1 {
			appendValue(createReidsrKey(REQUEST_UID_KEY, k), v)
		}
	}
	return setPairs(requestParis)
}

/*
*
指定具体用户更新对应的请求列表,如果当前用户对应的请求列表没有存在redis,需要调用该方法(初始化用户对应的请求列表)
*
*/
func specifyUserRefreshRequests(uid string, requests []Common.ReuqestOfAddingFriend) error {
	requestParis := make(map[string]interface{}, len(requests)+1)
	var uidRequests string

	for _, request := range requests {
		stream, _ := json.Marshal(request)
		requestParis[createReidsrKey(REQUEST_UID_KEY, request.Id)] = stream
		uidRequests = uidRequests + "|" + fmt.Sprint(request.Id)
	}
	requestParis[createReidsrKey(REQUEST_UID_KEY, uid)] = uidRequests
	return setPairs(requestParis)
}

/*
**
检查对应的请求是否存在
input: 需要检查的请求.
return: 存在的请求id
**
*/
func checkRequestsExist(requests []Common.ReuqestOfAddingFriend) []bool {
	requestIds := make([]bool, len(requests))
	for index, v := range requests {
		result, _ := exists(createReidsrKey(REQUEST_RID_KEY, v.Id))
		if result == 0 {
			requestIds[index] = false
		} else {
			requestIds[index] = true
		}
	}
	return requestIds
}

/*
**
通过uid获取所有的好友id
*
*/
func getFriendsUidByUid(uid int) ([]string, error) {
	friendsUid, err := LRange(createReidsrKey(FRIEND_KEY, uid))
	return friendsUid, err
}

/*
**
更新uid对应的所有好友的id数据
*
*/
func updateFriendsUidByUid(uid int, friendsUid []int) error {
	err := RPush(createReidsrKey(FRIEND_KEY, uid), friendsUid)
	return err
}

func removenullvalue[T string | int](slice []T) []T {
	var output []T
	for _, element := range slice {
		if reflect.ValueOf(element).IsValid() { //if condition satisfies add the elements in new slice
			output = append(output, element)
		}
	}
	return output //slice with no nil-values
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
