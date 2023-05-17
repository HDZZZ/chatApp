package user

import (
	"fmt"
	"strconv"

	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/db"
	"github.com/gin-gonic/gin"
)

type LoginService interface {
	login(c *gin.Context)
	register(c *gin.Context)
	getUserInfo(c *gin.Context)
}

type AppService interface {
	getConfigInfo(c *gin.Context)
}

type LoginServiceInstance struct {
}

type AppServiceInstance struct {
}

type AppConfig struct {
	Websocket string
}

func (instance *LoginServiceInstance) login(c *gin.Context) {

	users := DB.QueryUserByUserName(c.Query(Common.REQUEST_PARAMS_USERNAME))
	fmt.Println("QueryUserByUserName", users)
	for index, v := range users {
		if v.Password == c.Query(Common.REQUEST_PARAMS_PASSWORD) {
			c.JSON(200, Common.CreateResultDataSuccess(v))
			return
		}
		if (index + 1) == len(users) {
			c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_1004, Common.ErrCode[Common.ERROR_CODE_1004]))
			return
		}
	}
	c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_1003, Common.ErrCode[Common.ERROR_CODE_1003]))
}

func (instance *LoginServiceInstance) register(c *gin.Context) {
	users := DB.QueryUserByUserName(c.Query(Common.REQUEST_PARAMS_USERNAME))
	if len(users) > 0 {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_1002, Common.ErrCode[Common.ERROR_CODE_1002]))
		return
	}
	user, err := DB.AddUser(c.Query(Common.REQUEST_PARAMS_USERNAME), c.Query(Common.REQUEST_PARAMS_PASSWORD))
	if err != nil {
		value, _ := strconv.Atoi(err.Error())
		c.JSON(200, Common.CreateResultDataError(value, Common.ErrCode[value]))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess(user))
}

func (instance *AppServiceInstance) getConfigInfo(c *gin.Context) {
	// ip, err := Util.GetClientIp()
	// if err != nil {
	// 	c.JSON(200, createResultDataError(ERROR_CODE_101, errCode[ERROR_CODE_101]))
	// 	return
	// }
	c.JSON(200, Common.CreateResultDataSuccess(AppConfig{
		Websocket: "ws://120.79.7.215:9002/websocket",
		// Websocket: "ws://127.0.0.1:9002/websocket",
	}))
}

func getUserByToken(token string) DB.User {
	users := DB.QueryUserByToken(token)
	if len(users) > 0 {
		return users[0]
	}
	return DB.User{}
}
