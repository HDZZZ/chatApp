package db

import (
	"testing"
)

func TestMsgPush(t *testing.T) {
	// pushMessage(Common.DBMessage{
	// 	Sender_id:         100017,
	// 	Receiver_id:       100003,
	// 	Conversatio_type:  0,
	// 	Message_body_type: 1,
	// 	Content:           "这是文本消息",
	// })
}

func TestMsgQueryMessagesByUid(t *testing.T) {
	messages := queryMessagesByUid(100003)
	log("TestMsgQueryMessagesByUid,messages=", messages)
}

func TestMsgQueryMessagesByMsgIds(t *testing.T) {
	messages := queryMessagesByMsgIds("10000068")
	log("TestMsgQueryMessagesByMsgIds,messages=", messages)
}
