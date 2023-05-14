package common

type EventType int8

const (
	UserConnection EventType = 0
)

var eventSubscribers = [1][]func(params ...any){}

func CallUserConnectionSubscribers(uid int) {
	for _, callBack := range eventSubscribers[UserConnection] {
		callBack(uid)
	}
}

func RegisterSubscriber(eventType EventType, callback func(params ...any)) {
	eventSubscribers[eventType] = append(eventSubscribers[eventType], callback)
}
