package channel

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/db"
	"github.com/gorilla/websocket"
)

func receiveMessage(message Message, conn *websocket.Conn) {
	err := checkMessageLegal(message, conn)
	if err != nil {
		sendErrorMessage(err.Error(), conn)
		conn.Close()
		return
	}
	storeMessage(&message)
	tranferToReceiver(message, conn)
}

func tranferToReceiver(message Message, conn *websocket.Conn) {
	fmt.Println("tranferToReceiver")
	userCon := userChannals[message.ReceiverId]
	if userCon == nil {
		addWaitSendList(message)
		return
	}

	if userCon == ((*websocket.Conn)(nil)) {
		addWaitSendList(message)
		return
	}
	userCon.WriteJSON(message)
}

func addWaitSendList(message Message) {
	//todo, 走离线时间(断开连接存储),获取消息
	// Common.Set("uid", "sdajsfi")
	err := Common.AppendValue(Common.RedisKeyUnSendMsgIds+
		fmt.Sprint(message.ReceiverId), "|"+fmt.Sprint(message.MsgId))
	if err != nil {
		fmt.Println("addWaitSendList fiald", err, message)
		return
	}
}

func storeMessage(message *Message) {
	fmt.Println("storeMessage")

	id, err := DB.PushMessage(transferMessageToDBMessage(*message))
	if err != nil {
		fmt.Println("存储消息失败,message=", message)
		return
	}
	message.MsgId = id
}

func sendErrorMessage(content string, conn *websocket.Conn) {
	sendId, _ := mapkey(conn)
	body := MessageBody{
		MessageBodyType: Notification, Ext: ""}
	body.createNotificationContent(NotificationMessage{
		-1, content})
	conn.WriteJSON(Message{
		int64(rand.Intn(89999999) + 10000000), 000000, sendId, System, body})
}

func checkMessageLegal(message Message, conn *websocket.Conn) error {
	if message.SenderId == 0 {
		return errors.New("senderId can't be null")
	}
	if message.ReceiverId == 0 {
		return errors.New("receverId can't be null")
	}
	userConnection := userChannals[message.SenderId]
	if userConnection == nil {
		return errors.New("don't send verification message")
	}
	if conn != userConnection {
		return errors.New("aren't same person with senderId")
	}
	return nil
}

func mapkey(value *websocket.Conn) (key int, ok bool) {
	for k, v := range userChannals {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func transferMessageToDBMessage(message Message) DB.DBMessage {
	body, _ := json.Marshal(message.MessageBody)
	return DB.DBMessage{
		Msg_id: int(message.MsgId), Sender_id: message.SenderId, Receiver_id: message.ReceiverId,
		Conversatio_type: int(message.ConversationType), Message_body_type: int(message.MessageBody.MessageBodyType),
		Content: string(body)}
}

func transferDBMessageToMessage(message DB.DBMessage) Message {
	var body MessageBody
	json.Unmarshal([]byte(message.Content), &body)
	return Message{
		int64(message.Msg_id), message.Sender_id, message.Receiver_id,
		ConversationType(message.Conversatio_type), body}
}

func sendUnsendMessages(uid int) {
	// userChannals[iAreaId]
	value, _ := Common.Get(Common.RedisKeyUnSendMsgIds + fmt.Sprint(uid))
	if value == "" {
		return
	}
	msgs := DB.QueryMessagesByMsgIds(value[1:])
	for _, msg := range msgs {
		userChannals[uid].WriteJSON(msg)
	}
	Common.Delete(Common.RedisKeyUnSendMsgIds + fmt.Sprint(uid))
}

func init() {
	Common.RegisterSubscriber(Common.UserConnection, func(params ...any) {
		for _, uid := range params {
			iAreaId := uid.(int)
			sendUnsendMessages(iAreaId)
		}
	})
}
