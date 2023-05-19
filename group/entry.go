package group

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
	var groupService = GroupServiceInstance{}

	ginInstance.POST(http_path_create, groupService.createGroup)
	ginInstance.POST(http_path_transferOwner, groupService.tranferOwner)
	ginInstance.POST(http_path_leave, groupService.leaveGroup)
	ginInstance.POST(http_path_updateInfo, groupService.updateGroupInfo)
	ginInstance.GET(http_path_getGroupInfo, groupService.getGroupInfoByGid)
	ginInstance.GET(http_path_getAllGroups, groupService.getAllGroupsMeIn)

	var groupMemberService = GroupMemberServiceInstance{}

	ginInstance.POST(http_path_group_member_add, groupMemberService.addMember)
	ginInstance.POST(http_path_group_member_remove, groupMemberService.removeMember)
	ginInstance.GET(http_path_group_member_getAllMembers, groupMemberService.getAllMembersByGid)
	ginInstance.GET(http_path_group_member_getAllMembersId, groupMemberService.getAllMembersUidByGid)
	ginInstance.GET(http_path_group_member_getMember, groupMemberService.getMemberByGidAndUid)
	ginInstance.POST(http_path_group_member_updateMyInfo, groupMemberService.updateMyInfoInGroup)
}

func GroupMain() {

}
