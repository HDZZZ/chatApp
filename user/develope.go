package user

import (
	"github.com/gin-gonic/gin"
)

type DevelopeService interface {
	testDB(c *gin.Context)
}

type DevelopeServiceInstance struct {
}

func (instance *DevelopeServiceInstance) testDB(c *gin.Context) {
	// var user DBUser
	// users, _ := SQL.QueryStruct("select * from users", user)
	// // fmt.Println("")
	// c.JSON(200, Common.CreateResultDataSuccess(users))
}

type DBUser struct {
	Useruid  int    `mapstructure:"uid"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"pass_word"`
}
