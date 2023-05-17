package db

import (
	"errors"
	"log"
)

func SendRequest(sendUid int, receiverUid int, msg string) (int64, error) {
	// var value = fmt.Sprintf("%d,%d", userName, password);
	id, err := _exec("INSERT INTO request_add_friend(msg,sender_uid,receiver_uid) VALUES(?,?,?)",
		msg, sendUid, receiverUid)
	if err != nil {
		log.Fatal("insert into request_add_friend error", err)
		return 0, err
	}
	return id, err
}

func AgreeRequest(requestId int) error {
	// var value = fmt.Sprintf("%d,%d", userName, password);
	_, err := _exec("UPDATE request_add_friend set request_state = ? where id=?", AlreadyAgree, requestId)
	if err != nil {
		log.Fatal("insert into request_add_friend error", err)
		return err
	}

	request := QueryRequestById(requestId)
	if request == (ReuqestOfAddingFriend{}) {
		return errors.New("无此请求")
	}

	err = friendWithSomeone(request.Sender_id, request.Receiver_id)
	if err != nil {
		log.Fatal("add friend error", err)
		return err
	}
	return err
}

func RefuseRequest(requestId int) error {
	// var value = fmt.Sprintf("%d,%d", userName, password);
	_, err := _exec("UPDATE request_add_friend set request_state = ? where id=?", AlreadyRefuse, requestId)
	if err != nil {
		log.Fatal("insert into request_add_friend error", err)
		return err
	}
	return err
}

func MakeRequestState(uid_1 int, uid_2 int, state RequestState) error {
	// var value = fmt.Sprintf("%d,%d", userName, password);
	_, err := _exec("UPDATE request_add_friend set request_state = ? where (sender_uid=? AND receiver_uid=?) OR (sender_uid=? AND receiver_uid=?)", state, uid_1, uid_2, uid_2, uid_1)
	if err != nil {
		log.Fatal("UPDATE into request_add_friend error", err)
		return err
	}
	return err
}

func GetAllRequestOfSomebody(uid int) []ReuqestOfAddingFriend {
	var messages = []ReuqestOfAddingFriend{}
	inputUser := ReuqestOfAddingFriend{}

	_query(query_All_Request_By_Uid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{uid, uid}, &inputUser.Id, &inputUser.Msg, &inputUser.Sender_id, &inputUser.Receiver_id, &inputUser.Requst_state)
	return messages
}

func GetAllFriends(uid int) []User {
	var users = []User{}
	user := User{}

	_query(query_All_Friends_By_Uid, func(a ...any) {
		user := user
		users = append(users, user)
	}, []any{uid, uid, uid}, &user.Id, &user.UserName)
	return users
}

func GetAllFriendsUid(uid int) []int {
	var users = []int{}
	var friendUid int

	_query(query_All_Friends_Uid_By_Uid, func(a ...any) {
		users = append(users, friendUid)
	}, []any{uid, uid, uid}, &friendUid)
	return users
}

func DeleteFriend(uid int, friendUid int) error {
	_, err := _exec("DELETE FROM friend_relation WHERE (user_id_1 = ? AND user_id_2 = ?) OR (user_id_1 = ? AND user_id_2 = ?)", uid, friendUid, friendUid, uid)
	if err != nil {
		log.Fatal("DELETE friend error", err)
		return err
	}
	return nil
}

func QueryRequestById(requestId int) ReuqestOfAddingFriend {
	var messages = []ReuqestOfAddingFriend{}
	inputUser := ReuqestOfAddingFriend{}

	_query(query_Request_By_Id, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{requestId}, &inputUser.Id, &inputUser.Msg, &inputUser.Sender_id, &inputUser.Receiver_id, &inputUser.Requst_state)
	return messages[0]
}

func QueryRequestByUids(uid_1 int, uid_2 int) ReuqestOfAddingFriend {
	var messages = []ReuqestOfAddingFriend{}
	inputUser := ReuqestOfAddingFriend{}

	_query(query_Request_By_Uids, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{uid_1, uid_2, uid_2, uid_1}, &inputUser.Id, &inputUser.Msg, &inputUser.Sender_id, &inputUser.Receiver_id, &inputUser.Requst_state)
	return messages[0]
}

func friendWithSomeone(uid_1 int, uid_2 int) error {
	_, err := _exec("INSERT INTO friend_relation(user_id_1,user_id_2) VALUES(?,?)",
		uid_1, uid_2)
	if err != nil {
		log.Fatal("insert into request_add_friend error", err)
		return err
	}
	return nil
}
