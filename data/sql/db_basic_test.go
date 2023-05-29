package db

import (
	"testing"
)

func TestInsertRows(t *testing.T) {
	// insertRows("request_add_friend", []string{"msg", "sender_uid", "receiver_uid"}, []any{"msgs", 100006, 100001})
}

func TestUpdateRows(t *testing.T) {
	// updateRows("request_add_friend", fmt.Sprintf("request_state = %v", Common.AlreadyAgree), fmt.Sprintf("id=%v", 28))
}

func TestQuery(t *testing.T) {
	// users, _ := queryStruct[Common.User]("select users.uid, users.user_name, users.pass_word from users INNER JOIN friend_relation ON (users.uid=friend_relation.user_id_1 or users.uid=friend_relation.user_id_2) and users.uid != 100018 where friend_relation.user_id_1 = 100018 OR friend_relation.user_id_2 = 100018")
	// fmt.Println("TestQuery,users=", users)
}

func TestDeleteRows(t *testing.T) {
	// delRows("friend_relation", "(user_id_1 = 100018 AND user_id_2 = 100016) OR (user_id_1 = 100018 AND user_id_2 = 100016)")
}
