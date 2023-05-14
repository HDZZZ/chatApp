package main

import (
	DB "github.com/HDDDZ/test/chatApp/db"
	"github.com/gorilla/websocket"
)

type ClientChannel struct {
	connection *websocket.Conn
	userInfo   DB.User
}
