package common

type EventType int8

const (
	UserConnection EventType = 0
	AppClose       EventType = 1
)

var eventSubscribers = [2][]func(params ...any){}

func CallUserConnectionSubscribers(uid int) {
	for _, callBack := range eventSubscribers[UserConnection] {
		callBack(uid)
	}
}

func CallAppCloseSubscribers() {
	for _, callBack := range eventSubscribers[AppClose] {
		callBack()
	}
}

func RegisterSubscriber(eventType EventType, callback func(params ...any)) {
	eventSubscribers[eventType] = append(eventSubscribers[eventType], callback)
}
