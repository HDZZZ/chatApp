package user

import (
	Common "github.com/HDDDZ/test/chatApp/common"
	SQL "github.com/HDDDZ/test/chatApp/data/sql"
	"github.com/gin-gonic/gin"
)

type DevelopeService interface {
	testDB(c *gin.Context)
}

type DevelopeServiceInstance struct {
}

func (instance *DevelopeServiceInstance) testDB(c *gin.Context) {
	SQL.Query("select * from users")
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}
