package db

import (
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const _table_messages = "messages"
const _field_sender_id = "sender_id"
const _field_receiver_id = "receiver_id"
const _field_conversatio_type = "conversatio_type"
const _field_message_body_type = "message_body_type"
const _field_content = "content"

func pushMessage(message Common.DBMessage) (int64, error) {
	id, err := insertRows(_table_messages, []string{_field_sender_id, _field_receiver_id,
		_field_conversatio_type, _field_message_body_type, _field_content},
		[]any{message.Sender_id, message.Receiver_id, message.Conversatio_type,
			message.Message_body_type, message.Content})
	if err != nil {
		fmt.Println("insert into users error", err)
		return 0, err
	}
	return id, err
}

func queryMessagesByUid(uid int) []Common.DBMessage {
	messages, _ := queryStruct[Common.DBMessage](fmt.Sprintf(query_Messages_By_Uid, uid, uid))
	return messages
}

/*
*

	msgIds: msgid使用 | 拼接起来,如 "116812 | 148161 |151985"
*/
func queryMessagesByMsgIds(msgIds string) []Common.DBMessage {
	messages, _ := queryStruct[Common.DBMessage](fmt.Sprintf(
		query_Messages_By_MsgId_Rege, msgIds))
	return messages
}
