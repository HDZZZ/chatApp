package group

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	DBCommon "github.com/HDDDZ/test/chatApp/data/common"

	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/data"
	"github.com/gin-gonic/gin"
)

type GroupService interface {
	createGroup(c *gin.Context)
	tranferOwner(c *gin.Context)
	leaveGroup(c *gin.Context)
	updateGroupInfo(c *gin.Context)
	getGroupInfoByGid(c *gin.Context)
	getAllGroupsMeIn(c *gin.Context)
}

type GroupServiceInstance struct {
}

func (instance *GroupServiceInstance) createGroup(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	mermbers := AizuArray(c.Query(Common.REQUEST_PARAMS_MEMBER_UID), ",")
	fmt.Println("service,createGroup,mermbers=", mermbers)

	id, err := DB.Create(user.Id, createGroupName(user.UserName), mermbers...)
	if err == nil {
		group := DB.GetGroupByGid(int(id))
		c.JSON(200, Common.CreateResultDataSuccess(group))
		return
	}
	c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3001, Common.ErrCode[Common.ERROR_CODE_3001]+err.Error()))
}

func (instance *GroupServiceInstance) tranferOwner(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	newOWnerId, err := getMustParamNumber(Common.REQUEST_PARAMS_NEW_OWNER_UID, c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	if checkOperationPermission(user.Id, gid, DBCommon.Owner) {
		err = DB.TransferOwner(gid, newOWnerId, user.Id)
		if err != nil {
			c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3001, Common.ErrCode[Common.ERROR_CODE_3001]+err.Error()))
			return
		}
		c.JSON(200, Common.CreateResultDataSuccess("success"))
		return
	}
	c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_104, Common.ErrCode[Common.ERROR_CODE_104]))
}

func (instance *GroupServiceInstance) leaveGroup(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	group := DB.GetGroupByGid(gid)
	fmt.Println("group=", group)
	if group.OwnerId == user.Id && group.MemberCount != 1 {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3003, Common.ErrCode[Common.ERROR_CODE_3003]))
		return
	}
	err = DB.LeaveGroup(gid, user.Id)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3002, Common.ErrCode[Common.ERROR_CODE_3002]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}

func (instance *GroupServiceInstance) updateGroupInfo(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	groupName := c.Query(Common.REQUEST_PARAMS_GROUP_NAME)
	description := c.Query(Common.REQUEST_PARAMS_GROUP_DESCIPTION)
	if groupName == "" && description == "" {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3004, Common.ErrCode[Common.ERROR_CODE_3004]))
		return
	}
	if checkOperationPermission(user.Id, gid, DBCommon.Owner) {
		err = DB.UpdateGroupInfo(gid, groupName, description)
		if err != nil {
			c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3005, Common.ErrCode[Common.ERROR_CODE_3005]+err.Error()))
			return
		}
		c.JSON(200, Common.CreateResultDataSuccess("success"))
		return
	}
	c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_104, Common.ErrCode[Common.ERROR_CODE_104]))

}

func (instance *GroupServiceInstance) getGroupInfoByGid(c *gin.Context) {
	_, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	group := DB.GetGroupByGid(gid)
	if group == (DBCommon.Group{}) {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3006, Common.ErrCode[Common.ERROR_CODE_3006]))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess(group))
}

func (instance *GroupServiceInstance) getAllGroupsMeIn(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	groups := DB.GetAllGroupsByUid(user.Id)
	c.JSON(200, Common.CreateResultDataSuccess(groups))
}

/*
**

	通过gin.contxt获取用户, 完全处理token, 如果异常会直接调用c.json
*/
func getTokenByGin(c *gin.Context) (DBCommon.User, error) {
	tokens := c.Request.Header["Token"]
	if len(tokens) == 0 {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_103, Common.ErrCode[Common.ERROR_CODE_103]))
		return DBCommon.User{}, errors.New("have no token")
	}

	user := getUserByToken(tokens[0])
	if user == (DBCommon.User{}) {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_103, Common.ErrCode[Common.ERROR_CODE_103]))
		return DBCommon.User{}, errors.New("get get user by token")
	}
	return user, nil
}

func checkOperationPermission(uid, gid int, memberIdentity DBCommon.MemberIdentity) bool {
	group := DB.GetMemberInfo(gid, uid)
	return group.Identity == memberIdentity
}

func getUserByToken(token string) DBCommon.User {
	users := DB.QueryUserByToken(token)

	return users
}

/*
**
 */
func getMustParamNumber(paramKey string, c *gin.Context) (int, error) {
	param, err := strconv.Atoi(c.Query(paramKey))
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_102, Common.ErrCode[Common.ERROR_CODE_102]))
		return 0, errors.New("params not correct")
	}
	return param, nil
}

/*
**
 */
func getMustParam(paramKey string, c *gin.Context) (string, error) {
	param := c.Query(paramKey)
	if param == "" {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_102, Common.ErrCode[Common.ERROR_CODE_102]))
		return "", errors.New("params not correct")
	}
	return param, nil
}

func createGroupName(ownerName string) string {
	return ownerName + "的群聊"
}

func AizuArray(A string, N string) []int {
	if A == "" {
		return []int{}
	}
	a := strings.Split(A, N)
	b := make([]int, len(a))
	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}
	return b
}
