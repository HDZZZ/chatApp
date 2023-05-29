package user

import (
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/common"
	"github.com/gin-gonic/gin"
)

func init() {
	Common.RegisterHTTPService(initUserService)
}

func initUserService(ginInstance *gin.Engine) {
	fmt.Println("initUserService")
	var loginService = LoginServiceInstance{}

	ginInstance.POST("/user/login", loginService.login)
	ginInstance.POST("/user/register", loginService.register)

	var apppService = AppServiceInstance{}

	ginInstance.GET("/appConfig", apppService.getConfigInfo)

	var friendService = FriendServiceInstance{}

	ginInstance.POST(http_path_friend_send_requst, friendService.sendRequestOfAddingFriend)
	ginInstance.POST(http_path_friend_agrees_request, friendService.agreeRequestOfAddingFriend)
	ginInstance.POST(http_path_friend_refuse_request, friendService.refuseRequestOfAddingFriend)
	ginInstance.GET(http_path_friend_all_request, friendService.getRequestOfAddingFriendList)
	ginInstance.GET(http_path_friend_all, friendService.getAllFriendsInfo)
	ginInstance.GET(http_path_friend_all_uid, friendService.getAllFriendsUid)
	ginInstance.GET(http_path_friend_query, friendService.queryRquest)
	ginInstance.POST(http_path_friend_delete, friendService.deleteFriend)

	var dev = DevelopeServiceInstance{}
	ginInstance.POST("/testDB", dev.testDB)

}

func UserMain() {

}
