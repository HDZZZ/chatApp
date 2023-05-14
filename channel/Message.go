package channel

type ConversationType int8

const (
	Chat   ConversationType = 0
	Group  ConversationType = 2
	System ConversationType = 3
)

type Message struct {
	MsgId            int64            `json:"msgId"`
	SenderId         int              `json:"senderId"`
	ReceiverId       int              `json:"receiverId"`
	ConversationType ConversationType `json:"conversationType"`
	MessageBody      MessageBody      `json:"messageBody"`
}

type MessageBodyType int8

const (
	Text         MessageBodyType = 0
	Image        MessageBodyType = 1
	Custom       MessageBodyType = 6
	Notification MessageBodyType = 10
)

type MessageBody struct {
	MessageBodyType MessageBodyType `json:"messageBodyType"`
	Content         any             `json:"content"`
	Ext             any             `json:"ext"`
}

type ImageMessage struct {
	Url      string `json:"url"`
	Size     []int  `json:"size"`
	ThumbUrl string `json:"thumbUrl"`
}

type NotificationMessage struct {
	NotifiType int    `json:"notifiType"`
	Text       string `json:"text"`
}

func (msgBody *MessageBody) createTextContent(content string) {
	msgBody.Content = content
}

func (msgBody *MessageBody) createImageContent(message ImageMessage) {
	msgBody.Content = message
}

func (msgBody *MessageBody) createNotificationContent(message NotificationMessage) {
	msgBody.Content = message
}
