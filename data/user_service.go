package data

import (
	"errors"

	Common "github.com/HDDDZ/test/chatApp/data/common"
	DB "github.com/HDDDZ/test/chatApp/data/sql"
	Util "github.com/HDDDZ/test/chatApp/util"

	_ "github.com/go-sql-driver/mysql"
)

var userService DB.UserSQLService = &DB.SQLService{}

func QueryUserByToken(token string) Common.User {
	return userService.GetUserByToken(token)
}

func QueryUserByUserName(userName string) []Common.User {
	return userService.GetUserByUsername(userName)
}

func queryUserByUserNameAndPwd(userName string, passwrod string) (Common.User, error) {
	users := userService.GetUserByUsername(userName)
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
	return userService.AddUser(userName, password, Util.GenerateSecureToken(64))
}
