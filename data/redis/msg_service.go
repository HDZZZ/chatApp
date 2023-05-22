package redis

import (
	"encoding/json"
	"fmt"
	"strings"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const MESSAGE_MID = "message_msgid"     //
const MESSAGE_UNSEND = "message_unsend" //

func pushMessages(msgs ...Common.DBMessage) error {
	requestParis := make(map[string]string, len(msgs))
	for _, msg := range msgs {
		stream, _ := json.Marshal(msg)
		requestParis[createReidsrKey(MESSAGE_MID, msg.Msg_id)] = string(stream)
	}
	return setPairs(requestParis)
}

func pushMsgIdToWaitList(uid string, msgId int) error {
	return appendValue(createReidsrKey(MESSAGE_UNSEND, uid), "|"+fmt.Sprint(msgId))
}

func clearWaitListOfSomebody(uid string) error {
	_, err := delete(createReidsrKey(MESSAGE_UNSEND, uid))
	return err
}

func getAllWaitMsgIdOfSomeBody(uid string) []string {
	friendsUid, _ := get(createReidsrKey(MESSAGE_UNSEND, uid))
	if friendsUid == "" {
		return []string{}
	}

	uids := removenullvalue(strings.Split(friendsUid, "|"))
	return uids
}

func getMessageFromMsgID(uid string) Common.DBMessage {
	user, err := get(createReidsrKey(MESSAGE_MID, uid))
	if err != nil {
		return Common.DBMessage{}
	}
	var newUser Common.DBMessage
	json.Unmarshal([]byte(user), &newUser)
	return newUser
}
