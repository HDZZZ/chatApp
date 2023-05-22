package channel

import (
	"encoding/json"

	DBCommon "github.com/HDDDZ/test/chatApp/data/common"
	"github.com/gorilla/websocket"
)

var userChannels = make(map[int]*websocket.Conn)

// userId string, con *websocket.Conn
func CloseChannel(channel interface{}) {
	var iAreaId int
	var con *websocket.Conn
	switch channel.(type) {
	case int:
		iAreaId = channel.(int)
		con = userChannels[iAreaId]

	case *websocket.Conn:
		con = channel.(*websocket.Conn)
		iAreaId, _ = mapkey(con)

	case websocket.Conn:
		conn := (channel.(websocket.Conn))
		con = &conn
		iAreaId, _ = mapkey(con)
	default:
		return
	}
	con.Close()
	delete(userChannels, iAreaId)
}

func GetChannel(uid int) *websocket.Conn {
	return userChannels[uid]
}
func GetUidByChannel(con *websocket.Conn) int {
	uid, _ := mapkey(con)
	return uid
}

var messageListener = []func(uid int, msg Message){}

func RegisterReceiveMessageListener(listener func(uid int, msg Message)) {
	messageListener = append(messageListener, listener)
}

func notifyReceiverMsgListener(uid int, msg Message) {
	for _, listener := range messageListener {
		listener(uid, msg)
	}
}

func mapkey(value *websocket.Conn) (key int, ok bool) {
	for k, v := range userChannels {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func TransferMessageToDBMessage(message Message) DBCommon.DBMessage {
	body, _ := json.Marshal(message.MessageBody)
	return DBCommon.DBMessage{
		Msg_id: int(message.MsgId), Sender_id: message.SenderId, Receiver_id: message.ReceiverId,
		Conversatio_type: int(message.ConversationType), Message_body_type: int(message.MessageBody.MessageBodyType),
		Content: string(body)}
}

func TransferDBMessageToMessage(message DBCommon.DBMessage) Message {
	var body MessageBody
	json.Unmarshal([]byte(message.Content), &body)
	return Message{
		int64(message.Msg_id), message.Sender_id, message.Receiver_id,
		ConversationType(message.Conversatio_type), body}
}
