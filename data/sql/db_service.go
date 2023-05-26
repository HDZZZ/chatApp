package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	AppCommon "github.com/HDDDZ/test/chatApp/common"
	Common "github.com/HDDDZ/test/chatApp/data/common"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// dbc, err := sql.Open("mysql",
	// 	"root:hanzhi123@tcp(127.0.0.1:3306)/chat_app")
	dbc, err := sql.Open("mysql",
		"root:mysql@tcp(120.79.7.215:3306)/chat_app")
	if err != nil {
		log.Println(err)
	}
	db = dbc

	AppCommon.RegisterSubscriber(AppCommon.AppClose, func(params ...any) {
		appClosed()
	})
}

func queryUserByToken(token string) []Common.User {
	return _queryUserByAny("users_token.token", token)
}

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
	return Common.User{}, errors.New(string(errorCode))
}

func _queryUserByAny(queryKey string, value ...any) []Common.User {
	var user = []Common.User{}
	inputUser := Common.User{}
	_query(fmt.Sprintf(query_User_By_Users, queryKey), func(a ...any) {
		newUser := inputUser
		user = append(user, newUser)
	}, value, &inputUser.Id, &inputUser.UserName, &inputUser.Password, &inputUser.Token)
	return user
}

func addUser(userName string, password string, token string) (Common.User, error) {
	// var value = fmt.Sprintf("%d,%d", userName, password);
	id, err := _exec("INSERT INTO users(user_name,pass_word) VALUES(?,?)", userName, password)
	if err != nil {
		log.Println("insert into users error", err)
		return Common.User{}, err
	}
	_, err = _exec("INSERT INTO users_token(token,uid) VALUES(?,?)", token, id)
	if err != nil {
		log.Println("insert into users_token error", err)
		return Common.User{}, err
	}
	users := _queryUserByAny("users.uid", id)
	return users[0], err
}

/*
*
只有insert语句才会返回id
*/
func _exec(query string, args ...any) (lastId int64, err error) {
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	lastId, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("sql,_exec,lastId=", lastId)
	return
}

func _delete(query string, args ...any) (count int64, err error) {
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	count, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return
}

func _query(sqlQuery string, call func(...any), args []any, queryValue ...any) {

	rows, err := db.Query(sqlQuery, args...)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(queryValue...)
		if err != nil {
			log.Println(err)
		}
		call(queryValue...)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
}

func appClosed() {
	db.Close()
}
