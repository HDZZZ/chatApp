package db

import (
	"errors"
	"fmt"
	"reflect"

	Common "github.com/HDDDZ/test/chatApp/data/common"
	_ "github.com/go-sql-driver/mysql"
)

const _table_users = "users"
const _table_users_token = "users_token"
const _field_user_name = "user_name"
const _field_pass_word = "pass_word"
const _field_token = "token"

func queryUserByToken(token string) []Common.User {
	return _queryUserByAny("users_token.token", token)
}

// 必须要有token才能查到
func queryUserByUserName(userName string) []Common.User {
	return _queryUserByAny("users.user_name", userName)
}

func queryUserByUserNameAndPwd(userName string, passwrod string) (Common.User, error) {
	users := _queryUserByAny("users.user_name", userName)
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
	return Common.User{}, errors.New(fmt.Sprint(errorCode))
}

func _queryUserByAny(queryKey string, value any) []Common.User {
	var newValue any = value
	if reflect.TypeOf(value).Name() == "string" {
		newValue = "\"" + fmt.Sprint(newValue) + "\""
	}
	user, _ := queryStruct[Common.User](fmt.Sprintf(query_User_By_Users, queryKey, newValue))

	return user
}

func addUser(userName string, password string, token string) (Common.User, error) {
	id, err := insertRows(_table_users, []string{_field_user_name, _field_pass_word},
		[]any{userName, password})
	if err != nil {
		fmt.Println("insert into users error", err)
		return Common.User{}, err
	}

	id, err = insertRows(_table_users_token, []string{_field_token, _field_uid},
		[]any{token, id})

	if err != nil {
		fmt.Println("insert into users_token error", err)
		return Common.User{}, err
	}
	users := _queryUserByAny("users.uid", id)
	return users[0], err
}
