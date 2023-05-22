package data

import (
	"errors"
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
	Redis "github.com/HDDDZ/test/chatApp/data/redis"
	DB "github.com/HDDDZ/test/chatApp/data/sql"
	Util "github.com/HDDDZ/test/chatApp/util"

	_ "github.com/go-sql-driver/mysql"
)

var userDBService DB.UserSQLService = &DB.SQLService{}
var userRedisService Redis.UserRedisService = &Redis.RedisService{}

func QueryUserByToken(token string) Common.User {
	user := userRedisService.GetUserByToken(token)
	if user == (Common.User{}) {
		user = userDBService.GetUserByToken(token)
	}
	return user
}

func QueryUserByUserName(userName string) []Common.User {
	return userDBService.GetUserByUsername(userName)
}

func queryUserByUserNameAndPwd(userName string, passwrod string) (Common.User, error) {
	users := userDBService.GetUserByUsername(userName)
	var getByUserName bool
	for _, user := range users {
		getByUserName = true
		if user.Password == passwrod {
			return user, nil
		}
	}
	var errorCode int
	if getByUserName {
		errorCode = 101
	} else {
		errorCode = 102
	}
	return Common.User{}, errors.New(string(errorCode))
}

func AddUser(userName string, password string) (Common.User, error) {
	user, err := userDBService.AddUser(userName, password, Util.GenerateSecureToken(64))
	userRedisService.AddUserList(user)
	return user, err
}

func QueryUsersByUids[T string | int](uids []T) []Common.User {
	users := make([]Common.User, len(uids))
	var sqlQueryIds string
	for index, uidT := range uids {
		users[index] = userRedisService.GetUserByUid(fmt.Sprint(uidT))
		if users[index] == (Common.User{}) {
			sqlQueryIds = sqlQueryIds + "|" + fmt.Sprint(uidT)
		}
	}
	sqlUsers := userDBService.GetUserByUids(sqlQueryIds)
	for _, user := range sqlUsers {
		users[len(users)-1] = user
	}
	userRedisService.AddUserList(sqlUsers...)
	return users
}
