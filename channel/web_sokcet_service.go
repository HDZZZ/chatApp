package channel

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/data"
	DBCommon "github.com/HDDDZ/test/chatApp/data/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func init() {
	Common.RegisterHTTPService(InitWebSokcet)
}

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
	defer CloseChannel(conn)
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
	userChannels[user.Id] = conn
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
		notifyReceiverMsgListener(user.Id, userMessage)
	}
}

func checkUserInfo(token string) (DBCommon.User, error) {
	users := DB.QueryUserByToken(token)
	return users, nil
}

func ChannelMain() {

}
