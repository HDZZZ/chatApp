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
	log("QueryUserByToken,redis,user=", user)
	if user == (Common.User{}) {
		user = userDBService.GetUserByToken(token)
		log("QueryUserByToken,mysql,user=", user)
		err := userRedisService.AddUserList(user)
		log("QueryUserByToken,redis,err=", err)
	}
	return user
}

func QueryUserByUserName(userName string) []Common.User {
	users := userDBService.GetUserByUsername(userName)
	log("QueryUserByUserName,mysql,user=", users)
	return users
}

func queryUserByUserNameAndPwd(userName string, passwrod string) (Common.User, error) {
	users := userDBService.GetUserByUsername(userName)
	log("queryUserByUserNameAndPwd,mysql,user=", users)
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
	user, err := userDBService.AddUser(userName, password, Util.GenerateSecureToken(32))
	log("AddUser,mysql,user=", user)
	err = userRedisService.AddUserList(user)
	log("AddUser,redis,err=", err)
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
	log("QueryUsersByUids,redis,users=", users)
	sqlUsers := userDBService.GetUserByUids(sqlQueryIds)
	for _, user := range sqlUsers {
		users[len(users)-1] = user
	}
	log("QueryUsersByUids,mysql,users=", users)
	err := userRedisService.AddUserList(sqlUsers...)
	log("QueryUsersByUids,redis,AddUserList=", err)
	return users
}

func log(logs ...any) {
	fmt.Println("data,", logs)
}
