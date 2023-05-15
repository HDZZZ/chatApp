package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var httpServices []func(gin *gin.Engine)

/**
*	调用该方法注册请求接口, 必需放在init方法中调用, 否则不会生效
 */
func RegisterHTTPService(excute func(gin *gin.Engine)) {
	httpServices = append(httpServices, excute)
}

/*
*

	该方法需要放在main中调用
*/
func InitHTTPService() {

	ginInstance := setupRouter()
	for _, excute := range httpServices {
		excute(ginInstance)
	}
	ginInstance.Run(":9002")
}

func setupRouter() *gin.Engine {
	ginInstance := gin.Default()
	ginInstance.Use(Cors())

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
