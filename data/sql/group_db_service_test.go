package db

import (
	"testing"

	"github.com/HDDDZ/test/chatApp/data/common"
)

func TestGroupCreate(t *testing.T) {
	// create(100000, "streep club3", 100010, 100011, 100012, 100013)
}
func TestGroupTransferOwner(t *testing.T) {
	transferOwner(100000019, 100011, 100000)
}

func TestGroupLeaveGroup(t *testing.T) {
	leaveGroup(100000019, 100000)
}

func TestGroupUpdateGroupInfo(t *testing.T) {
	updateGroupInfo(100000019, "streep club4", "welcome to streep club")
}

func TestGroupGetGroupByGid(t *testing.T) {
	group := getGroupByGid(100000019)
	log("TestGroupGetGroupByGid,group=", group)
}

func TestGroupGetAllGroupsByUid(t *testing.T) {
	groups := getAllGroupsByUid(100003)
	log("TestGroupGetAllGroupsByUid,group=", groups)
}

func TestGroupAddMember(t *testing.T) {
	addMember(100000019, 100018)
}

func TestGroupRemoveMember(t *testing.T) {
	removeMember(100000019, 100018)
}

func TestGroupGetAllMembersInfo(t *testing.T) {
	groups := getAllMembersInfo(100000019)
	log("TestGroupGetAllMembersInfo,group=", groups)
}

func TestGroupGetAllMembersUid(t *testing.T) {
	groups := getAllMembersUid(100000019)
	log("TestGroupGetAllMembersUid,group=", groups)
}

func TestGroupGetMemberInfo(t *testing.T) {
	groups := getMemberInfo(100000019, 100017)
	log("TestGroupGetMemberInfo,group=", groups)
}

func TestGroupUpdateMemberInfo(t *testing.T) {
	groups := updateMemberInfo(100000019, 100017, "小当家", common.Manager)
	log("TestGroupGetMemberInfo,group=", groups)
}
