package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const FRIEND_KEY = "friend_uid"

const REQUEST_KEY = "request"
const REQUEST_RID_KEY = "request_rid"
const REQUEST_UID_KEY = "request_uid"

func getRequestByUid(uid string) Common.ReuqestOfAddingFriend {
	rid, err := get(createReidsrKey(REQUEST_UID_KEY, uid))
	if err != nil {
		return Common.ReuqestOfAddingFriend{}
	}
	return getRequestByRid(rid)
}

func getRequestByRid(rid string) Common.ReuqestOfAddingFriend {
	user, err := get(createReidsrKey(REQUEST_RID_KEY, rid))
	if err != nil {
		return Common.ReuqestOfAddingFriend{}
	}
	var newUser Common.ReuqestOfAddingFriend
	json.Unmarshal([]byte(user), &newUser)
	return newUser
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
	friendsUid, _ := get(createReidsrKey(FRIEND_KEY, uid))
	if friendsUid == "" {
		return []string{}, errors.New("don't have")
	}

	uids := removenullvalue(strings.Split(friendsUid, "|"))
	return uids, nil
}

/*
**
更新uid对应的所有好友的id数据
*
*/
func updateFriendsUidByUid(uid int, friendsUid []int) error {
	err := set(createReidsrKey(FRIEND_KEY, uid), friendsUid)
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
