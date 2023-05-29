package db

import (
	"errors"
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const _table_reuest_add_friend = "request_add_friend"
const _table_friend_relation = "friend_relation"
const _field_msg = "msg"
const _field_sender_uid = "sender_uid"
const _field_receiver_uid = "receiver_uid"
const _field_user_id_1 = "user_id_1"
const _field_user_id_2 = "user_id_2"

func sendRequest(sendUid int, receiverUid int, msg string) (int64, error) {
	id, err := insertRows(_table_reuest_add_friend, []string{_field_msg, _field_sender_uid,
		_field_receiver_uid}, []any{msg, sendUid, receiverUid})

	return id, err
}

func agreeRequest(requestId int) error {
	err := updateRows(_table_reuest_add_friend, fmt.Sprintf("request_state = %v",
		Common.AlreadyAgree), fmt.Sprintf("id = %v", requestId))
	if err != nil {
		fmt.Println("insert into request_add_friend error", err)
		return err
	}
	request := queryRequestById(requestId)
	if request == (Common.ReuqestOfAddingFriend{}) {
		return errors.New("无此请求")
	}

	err = friendWithSomeone(request.Sender_id, request.Receiver_id)
	if err != nil {
		fmt.Println("add friend error", err)
		return err
	}
	return err
}

func refuseRequest(requestId int) error {
	err := updateRows(_table_reuest_add_friend, fmt.Sprintf("request_state = %v",
		Common.AlreadyRefuse), fmt.Sprintf("id = %v", requestId))
	if err != nil {
		fmt.Println("insert into request_add_friend error", err)
		return err
	}
	return err
}

func makeRequestState(uid_1 int, uid_2 int, state Common.RequestState) error {
	err := updateRows(_table_reuest_add_friend, fmt.Sprintf("request_state = %v",
		state), fmt.Sprintf(`(sender_uid=%v AND receiver_uid=%v) OR 
		(sender_uid=%v AND receiver_uid=%v)`, uid_1, uid_2, uid_2, uid_1))
	if err != nil {
		fmt.Println("UPDATE into request_add_friend error", err)
		return err
	}
	return err
}

func getAllRequestOfSomebody(uid int) []Common.ReuqestOfAddingFriend {
	messages, _ := queryStruct[Common.ReuqestOfAddingFriend](fmt.Sprintf(query_All_Request_By_Uid, uid, uid))
	return messages
}

func getAllFriends(uid int) []Common.User {
	users, _ := queryStruct[Common.User](fmt.Sprintf(query_All_Friends_By_Uid, uid, uid, uid))
	return users
}

func getAllFriendsUid(uid int) []int {
	var users = []int{}
	var friendUid int

	queryRows(fmt.Sprintf(query_All_Friends_Uid_By_Uid, uid, uid, uid), func() {
		users = append(users, friendUid)
	}, &friendUid)
	return users
}

func deleteFriend(uid int, friendUid int) error {
	err := delRows(_table_friend_relation, fmt.Sprintf(`(user_id_1 = %v AND user_id_2 = %v) 
	OR (user_id_1 = %v AND user_id_2 = %v)`, uid, friendUid, friendUid, uid))
	if err != nil {
		fmt.Println("DELETE friend error", err)
		return err
	}
	return nil
}

func queryRequestById(requestId int) Common.ReuqestOfAddingFriend {
	requests, _ := queryStruct[Common.ReuqestOfAddingFriend](
		fmt.Sprintf(query_Request_By_Id, requestId))
	if len(requests) == 0 {
		return Common.ReuqestOfAddingFriend{}
	}
	return requests[0]
}

func queryRequestBySuidAndRuid(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	requests, _ := queryStruct[Common.ReuqestOfAddingFriend](fmt.Sprintf(
		query_Request_By_Uids, uid_1, uid_2, uid_2, uid_1))
	if len(requests) == 0 {
		return Common.ReuqestOfAddingFriend{}
	}
	return requests[0]
}

func queryRequestByUids(uid_1 int, uid_2 int) Common.ReuqestOfAddingFriend {
	requests, _ := queryStruct[Common.ReuqestOfAddingFriend](fmt.Sprintf(
		query_Request_By_Uids, uid_1, uid_2, uid_2, uid_1))
	if len(requests) == 0 {
		return Common.ReuqestOfAddingFriend{}
	}
	return requests[0]
}

func friendWithSomeone(uid_1 int, uid_2 int) error {
	_, err := insertRows(_table_friend_relation, []string{_field_user_id_1, _field_user_id_2},
		[]any{uid_1, uid_2})

	if err != nil {
		fmt.Println("insert into request_add_friend error", err)
		return err
	}
	return nil
}
