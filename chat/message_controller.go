package channel

import (
	"errors"
	"fmt"
	"math/rand"

	Channel "github.com/HDDDZ/test/chatApp/channel"
	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/db"
	"github.com/gorilla/websocket"
)

func receiveMessage(message Channel.Message, conn *websocket.Conn) {
	err := checkMessageLegal(message, conn)
	if err != nil {
		sendErrorMessage(err.Error(), conn)
		conn.Close()
		return
	}
	storeMessage(&message)
	tranferToReceiver(message, conn)
}

func tranferToReceiver(message Channel.Message, conn *websocket.Conn) {
	fmt.Println("tranferToReceiver")
	userCon := Channel.GetChannel(message.ReceiverId)
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

func addWaitSendList(message Channel.Message) {
	//todo, 走离线时间(断开连接存储),获取消息
	// Common.Set("uid", "sdajsfi")
	err := Common.AppendValue(Common.RedisKeyUnSendMsgIds+
		fmt.Sprint(message.ReceiverId), "|"+fmt.Sprint(message.MsgId))
	if err != nil {
		fmt.Println("addWaitSendList fiald", err, message)
		return
	}
}

func storeMessage(message *Channel.Message) {
	fmt.Println("storeMessage")

	id, err := DB.PushMessage(Channel.TransferMessageToDBMessage(*message))
	if err != nil {
		fmt.Println("存储消息失败,message=", message)
		return
	}
	message.MsgId = id
}

func sendErrorMessage(content string, conn *websocket.Conn) {
	sendId := Channel.GetUidByChannel(conn)
	body := Channel.MessageBody{
		MessageBodyType: Channel.Notification, Ext: ""}
	body.CreateNotificationContent(Channel.NotificationMessage{
		NotifiType: -1, Text: content})
	conn.WriteJSON(Channel.Message{
		MsgId: int64(rand.Intn(89999999) + 10000000), SenderId: 000000, ReceiverId: sendId, ConversationType: Channel.System, MessageBody: body})
}

func checkMessageLegal(message Channel.Message, conn *websocket.Conn) error {
	if message.SenderId == 0 {
		return errors.New("senderId can't be null")
	}
	if message.ReceiverId == 0 {
		return errors.New("receverId can't be null")
	}
	userConnection := Channel.GetChannel(message.SenderId)
	if userConnection == nil {
		return errors.New("don't send verification message")
	}
	if conn != userConnection {
		return errors.New("aren't same person with senderId")
	}
	return nil
}

func sendUnsendMessages(uid int) {
	// userChannals[iAreaId]
	value, _ := Common.Get(Common.RedisKeyUnSendMsgIds + fmt.Sprint(uid))
	if value == "" {
		return
	}
	msgs := DB.QueryMessagesByMsgIds(value[1:])
	for _, msg := range msgs {
		Channel.GetChannel(uid).WriteJSON(Channel.TransferDBMessageToMessage(msg))
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