package user

import (
	Common "github.com/HDDDZ/test/chatApp/common"
	"github.com/gin-gonic/gin"
)

func init() {
	Common.RegisterHTTPService(initUserService)
}

func initUserService(ginInstance *gin.Engine) {
	var loginService = LoginServiceInstance{}

	ginInstance.POST("/user/login", loginService.login)
	ginInstance.POST("/user/register", loginService.register)

	var apppService = AppServiceInstance{}

	ginInstance.GET("/appConfig", apppService.getConfigInfo)
}
