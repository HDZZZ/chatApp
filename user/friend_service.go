package user

import (
	"errors"
	"fmt"
	"strconv"

	DBCommon "github.com/HDDDZ/test/chatApp/data/common"

	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/data"
	"github.com/gin-gonic/gin"
)

type FriendService interface {
	sendRequestOfAddingFriend(c *gin.Context)
	agreeRequestOfAddingFriend(c *gin.Context)
	refuseRequestOfAddingFriend(c *gin.Context)
	getRequestOfAddingFriendList(c *gin.Context)
	getAllFriendsInfo(c *gin.Context)
	getAllFriendsUid(c *gin.Context)
	deleteFriend(c *gin.Context)
}

type FriendServiceInstance struct {
}

func (instance *FriendServiceInstance) sendRequestOfAddingFriend(c *gin.Context) {
	receiverUid, err := getMustParamNumber(Common.REQUEST_PARAMS_RECEIVERUID, c)
	if err != nil {
		return
	}
	msg := c.Query(Common.REQUEST_PARAMS_REQUEST_MSG)
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	//防重与删了后再次添加
	olderReuest := DB.QueryRequestByUids(user.Id, receiverUid)
	if olderReuest != (DBCommon.ReuqestOfAddingFriend{}) {
		if olderReuest.Requst_state == DBCommon.Defualt {
			c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2007, Common.ErrCode[Common.ERROR_CODE_2007]))
		} else {
			err = DB.MakeRequestState(user.Id, receiverUid, DBCommon.Defualt)
			if err == nil {
				c.JSON(200, Common.CreateResultDataSuccess("success"))
			} else {
				c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2008, Common.ErrCode[Common.ERROR_CODE_2008]))
			}
		}
		return
	}

	_, err = DB.SendRequest(user.Id, receiverUid, msg)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2001, Common.ErrCode[Common.ERROR_CODE_2001]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}
func (instance *FriendServiceInstance) agreeRequestOfAddingFriend(c *gin.Context) {
	requestID, err := getMustParamNumber(Common.REQUEST_PARAMS_REQUEST_ID, c)
	if err != nil {
		return
	}
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	err = checkHaveOperation(requestID, user.Id, c)
	if err != nil {
		return
	}
	fmt.Println("agree,", "checkHaveOperation passed")
	err = DB.AgreeRequest(requestID)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2002, Common.ErrCode[Common.ERROR_CODE_2002]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}

func (instance *FriendServiceInstance) refuseRequestOfAddingFriend(c *gin.Context) {
	requestID, err := getMustParamNumber(Common.REQUEST_PARAMS_REQUEST_ID, c)
	if err != nil {
		return
	}
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	err = checkHaveOperation(requestID, user.Id, c)
	if err != nil {
		return
	}
	err = DB.RefuseRequest(requestID)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2003, Common.ErrCode[Common.ERROR_CODE_2003]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}

func (instance *FriendServiceInstance) getRequestOfAddingFriendList(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	requests := DB.GetAllRequestOfSomebody(user.Id)
	c.JSON(200, Common.CreateResultDataSuccess(requests))
}

func (instance *FriendServiceInstance) getAllFriendsInfo(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	users := DB.GetAllFriends(user.Id)
	c.JSON(200, Common.CreateResultDataSuccess(users))
}

func (instance *FriendServiceInstance) getAllFriendsUid(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	users := DB.GetAllFriendsUid(user.Id)
	c.JSON(200, Common.CreateResultDataSuccess(users))
}

func (instance *FriendServiceInstance) queryRquest(c *gin.Context) {
	_, err := getTokenByGin(c)
	if err != nil {
		return
	}
	requestID, err := getMustParamNumber(Common.REQUEST_PARAMS_REQUEST_ID, c)
	if err != nil {
		return
	}
	users := DB.QueryRequestById(requestID)
	c.JSON(200, Common.CreateResultDataSuccess(users))
}

func (instance *FriendServiceInstance) deleteFriend(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	friendId, err := getMustParamNumber(Common.REQUEST_PARAMS_FRIEND_UID, c)
	if err != nil {
		return
	}
	err = DB.DeleteFriend(user.Id, friendId)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2006, Common.ErrCode[Common.ERROR_CODE_2006]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
	err = DB.MakeRequestState(user.Id, friendId, DBCommon.NotWork)
	if err != nil {
		return
	}
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
*
检查本人是否有权限行使对该条请求的状态修改
*/
func checkHaveOperation(requestID int, userId int, c *gin.Context) error {
	request := DB.QueryRequestById(requestID)
	if request == (DBCommon.ReuqestOfAddingFriend{}) {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2004, Common.ErrCode[Common.ERROR_CODE_2004]))
		return errors.New(Common.ErrCode[Common.ERROR_CODE_2004])
	}
	if userId != request.Receiver_id {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_104, Common.ErrCode[Common.ERROR_CODE_104]))
		return errors.New(Common.ErrCode[Common.ERROR_CODE_104])
	}
	if request.Requst_state != DBCommon.Defualt {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_2005, Common.ErrCode[Common.ERROR_CODE_2005]))
		return errors.New(Common.ErrCode[Common.ERROR_CODE_104])
	}
	return nil
}
