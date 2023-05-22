package data

import (
	Common "github.com/HDDDZ/test/chatApp/data/common"
	DB "github.com/HDDDZ/test/chatApp/data/sql"
)

var messageDBService DB.MsgSQLService = &DB.SQLService{}

func PushMessage(message Common.DBMessage) (int64, error) {
	return messageDBService.PushMessage(message)
}

func QueryMessagesByUid(uid int) []Common.DBMessage {
	return messageDBService.QueryMessagesByUid(uid)
}

/*
*

	msgIds: msgid使用 | 拼接起来,如 "116812 | 148161 |151985"
*/
func QueryMessagesByMsgIds(msgIds string) []Common.DBMessage {
	return messageDBService.QueryMessagesByMsgIds(msgIds)
}
