package group

import (
	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/data"
	DBCommon "github.com/HDDDZ/test/chatApp/data/common"
	"github.com/gin-gonic/gin"
)

type GroupMemberService interface {
	addMember(c *gin.Context)
	removeMember(c *gin.Context)
	getAllMembersByGid(c *gin.Context)
	getAllMembersUidByGid(c *gin.Context)
	getMemberByGidAndUid(c *gin.Context)
	updateMyInfoInGroup(c *gin.Context)
}

type GroupMemberServiceInstance struct {
}

func (instance *GroupMemberServiceInstance) addMember(c *gin.Context) {
	_, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	mermbers := AizuArray(c.Query(Common.REQUEST_PARAMS_GROUP_MEMBER_ID), ",")
	if len(mermbers) == 0 {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_102, Common.ErrCode[Common.ERROR_CODE_102]))
		return
	}
	err = DB.AddMember(gid, mermbers...)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3007, Common.ErrCode[Common.ERROR_CODE_3007]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}
func (instance *GroupMemberServiceInstance) removeMember(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	mermbers := AizuArray(c.Query(Common.REQUEST_PARAMS_GROUP_MEMBER_ID), ",")
	if len(mermbers) == 0 {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_102, Common.ErrCode[Common.ERROR_CODE_102]))
		return
	}
	if !checkOperationPermission(user.Id, gid, DBCommon.Owner) {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_104, Common.ErrCode[Common.ERROR_CODE_104]))
		return
	}
	err = DB.RemoveMember(gid, mermbers...)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3001, Common.ErrCode[Common.ERROR_CODE_3001]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))

}
func (instance *GroupMemberServiceInstance) getAllMembersByGid(c *gin.Context) {
	_, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	groupMembers := DB.GetAllMembersInfo(gid)
	c.JSON(200, Common.CreateResultDataSuccess(groupMembers))

}
func (instance *GroupMemberServiceInstance) getAllMembersUidByGid(c *gin.Context) {
	_, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	groupMembers := DB.GetAllMembersUid(gid)
	c.JSON(200, Common.CreateResultDataSuccess(groupMembers))
}
func (instance *GroupMemberServiceInstance) getMemberByGidAndUid(c *gin.Context) {
	_, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	memberUid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_MEMBER_ID, c)
	if err != nil {
		return
	}
	groupMembers := DB.GetMemberInfo(gid, memberUid)
	c.JSON(200, Common.CreateResultDataSuccess(groupMembers))
}
func (instance *GroupMemberServiceInstance) updateMyInfoInGroup(c *gin.Context) {
	user, err := getTokenByGin(c)
	if err != nil {
		return
	}
	gid, err := getMustParamNumber(Common.REQUEST_PARAMS_GROUP_GID, c)
	if err != nil {
		return
	}
	alias, err := getMustParam(Common.REQUEST_PARAMS_GROUP_MEMBER_ALIAS, c)
	if err != nil {
		return
	}
	err = DB.UpdateMemberInfo(gid, user.Id, alias, DBCommon.NONE)
	if err != nil {
		c.JSON(200, Common.CreateResultDataError(Common.ERROR_CODE_3008, Common.ErrCode[Common.ERROR_CODE_3008]+err.Error()))
		return
	}
	c.JSON(200, Common.CreateResultDataSuccess("success"))
}
