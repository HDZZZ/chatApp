package data

import (
	Common "github.com/HDDDZ/test/chatApp/data/common"
	Redis "github.com/HDDDZ/test/chatApp/data/redis"
	DB "github.com/HDDDZ/test/chatApp/data/sql"
)

var messageDBService DB.MsgSQLService = &DB.SQLService{}
var msgRedisService Redis.MsgRedisService = &Redis.RedisService{}

func PushMessage(message Common.DBMessage) (int64, error) {
	asyncExcute(func() {
		messageDBService.PushMessage(message)
		msgRedisService.PushMessages(message)
	})
	return 0, nil
}

func QueryMessagesByUid(uid int) []Common.DBMessage {
	msgs := messageDBService.QueryMessagesByUid(uid)
	if len(msgs) != 0 {
		go msgRedisService.PushMessages(msgs...)
	}
	return msgs
}

/*
*

	msgIds: msgid使用 | 拼接起来,如 "116812 | 148161 |151985"
*/
func QueryMessagesByMsgIds(msgIds string) []Common.DBMessage {
	msg := msgRedisService.GetMessageFromMsgID(msgIds)
	if msg == (Common.DBMessage{}) {
		msgs := messageDBService.QueryMessagesByMsgIds(msgIds)
		msgRedisService.PushMessages(msgs...)
		return msgs
	}
	return []Common.DBMessage{msg}
}

func asyncExcute(exucte func()) {
	go exucte()
}
