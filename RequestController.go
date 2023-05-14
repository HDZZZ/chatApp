package main

import (
	"net/http"

	channel "github.com/HDDDZ/test/chatApp/channel"
	"github.com/gin-gonic/gin"
)

type Request struct {
}

func init() {
	http.HandleFunc("/v3", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	// var loginService = LoginServiceInstance{}

	// loginService.login()
	ginInstance := setupRouter()
	channel.InitWebSokcet(ginInstance)
	ginInstance.Run(":9002")
}

func setupRouter() *gin.Engine {
	var loginService = LoginServiceInstance{}
	ginInstance := gin.Default()
	ginInstance.Use(Cors())
	ginInstance.POST("/user/login", loginService.login)
	ginInstance.POST("/user/register", loginService.register)

	var apppService = AppServiceInstance{}

	ginInstance.GET("/appConfig", apppService.getConfigInfo)
	return ginInstance
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func registerHTTPApi(url string, callback func()) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		callback()
	})
}
