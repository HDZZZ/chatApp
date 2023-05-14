package channel

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/db"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func InitWebSokcet(ginInstance *gin.Engine) {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ginInstance.GET("/websocket", func(ctx *gin.Context) {

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		go createNewConnection(conn)
	})
}

func createNewConnection(conn *websocket.Conn) {
	defer conn.Close()
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("get first message failed:", err)
		return
	}
	user, err := checkUserInfo(string(message))
	if err != nil {
		log.Println(err)
		return
	}
	userChannals[user.Id] = conn
	Common.CallUserConnectionSubscribers(user.Id)
	for {
		mt, message, err := conn.ReadMessage()
		fmt.Println("读到了消息,mt=", mt, "message=", string(message), "error=", err)
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		var userMessage Message
		json.Unmarshal(message, &userMessage)
		receiveMessage(userMessage, conn)
	}
}

func checkUserInfo(token string) (DB.User, error) {
	users := DB.QueryUserByToken(token)
	if len(users) == 0 {
		return DB.User{}, errors.New("user token verify error")
	}
	return users[0], nil
}
