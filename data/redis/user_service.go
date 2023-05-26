package redis

import (
	"encoding/json"
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const USER_KEY = "user"
const USER_TOKEN_KEY = "user_token"
const USER_UID_KEY = "user_uid"

func getUserByToken(token string) Common.User {
	uid, err := get(createReidsrKey(USER_TOKEN_KEY, token))
	if err != nil {
		return Common.User{}
	}
	return getUserByUid(uid)
}

func getUserByUid(uid string) Common.User {
	user, err := get(createReidsrKey(USER_UID_KEY, uid))
	if err != nil {
		return Common.User{}
	}
	var newUser Common.User
	json.Unmarshal([]byte(user), &newUser)
	return newUser
}

func addUsers(users []Common.User) error {
	mUsers := make(map[string]interface{}, len(users)*2)
	for _, user := range users {
		stream, _ := json.Marshal(user)
		mUsers[createReidsrKey(USER_UID_KEY, user.Id)] = string(stream)
		mUsers[createReidsrKey(USER_TOKEN_KEY, user.Token)] = user.Id
	}
	return setPairs(mUsers)
}

func createReidsrKey(key string, token any) string {
	return key + string(fmt.Sprint(token))
}
