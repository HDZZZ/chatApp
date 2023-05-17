package db

import (
	"fmt"
	"log"
)

func PushMessage(message DBMessage) (int64, error) {
	// var value = fmt.Sprintf("%d,%d", userName, password);
	id, err := _exec("INSERT INTO messages(sender_id,receiver_id,conversatio_type,message_body_type,content) VALUES(?,?,?,?,?)",
		message.Sender_id, message.Receiver_id, message.Conversatio_type, message.Message_body_type, message.Content)
	if err != nil {
		log.Fatal("insert into users error", err)
		return 0, err
	}
	return id, err
}

func QueryMessagesByUid(uid int) []DBMessage {
	var messages = []DBMessage{}
	inputUser := DBMessage{}
	_query(query_Messages_By_Uid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{uid}, &inputUser.Msg_id, &inputUser.Sender_id, &inputUser.Receiver_id, &inputUser.Conversatio_type, &inputUser.Message_body_type, &inputUser.Content)
	return messages
}

/*
*

	msgIds: msgid使用 | 拼接起来,如 "116812 | 148161 |151985"
*/
func QueryMessagesByMsgIds(msgIds string) []DBMessage {
	var messages = []DBMessage{}
	inputUser := DBMessage{}

	_query(fmt.Sprintf(query_Messages_By_MsgId_Rege, msgIds), func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{}, &inputUser.Msg_id, &inputUser.Sender_id, &inputUser.Receiver_id, &inputUser.Conversatio_type, &inputUser.Message_body_type, &inputUser.Content)
	return messages
}
