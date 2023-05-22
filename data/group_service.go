package data

import (
	Common "github.com/HDDDZ/test/chatApp/data/common"
	DB "github.com/HDDDZ/test/chatApp/data/sql"
)

var groupDBService DB.GroupSQLService = &DB.SQLService{}

func Create(ownerUid int, groupName string, memberUids ...int) (int64, error) {
	return groupDBService.Create(ownerUid, groupName, memberUids...)
}

func TransferOwner(gid int, newOwnerId int, OlderOwnerId int) error {
	return groupDBService.TransferOwner(gid, newOwnerId, OlderOwnerId)

}

func LeaveGroup(gid int, uid int) error {
	return groupDBService.LeaveGroup(gid, uid)
}

func UpdateGroupInfo(gid int, groupName string, description string) error {
	return groupDBService.UpdateGroupInfo(gid, groupName, description)
}

func GetGroupByGid(gid int) Common.Group {
	return groupDBService.GetGroupByGid(gid)
}

func GetAllGroupsByUid(uid int) []Common.Group {
	return groupDBService.GetAllGroupsByUid(uid)
}

func AddMember(gid int, uids ...int) error {
	return groupDBService.AddMember(gid, uids...)
}

func RemoveMember(gid int, uids ...int) error {
	return groupDBService.RemoveMember(gid, uids...)
}

func GetAllMembersInfo(gid int) []Common.GroupMember {
	return groupDBService.GetAllMembersInfo(gid)
}

func GetAllMembersUid(gid int) []int {
	return groupDBService.GetAllMembersUid(gid)
}

func GetMemberInfo(gid int, uid int) Common.GroupMember {
	return groupDBService.GetMemberInfo(gid, uid)
}

func UpdateMemberInfo(gid int, uid int, alias string, identity Common.MemberIdentity) error {
	return groupDBService.UpdateMemberInfo(gid, uid, alias, identity)
}
