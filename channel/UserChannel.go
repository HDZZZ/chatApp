package channel

import "github.com/gorilla/websocket"

var userChannals = make(map[int]*websocket.Conn)
