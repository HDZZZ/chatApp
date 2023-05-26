package data

import (
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
	Redis "github.com/HDDDZ/test/chatApp/data/redis"
	DB "github.com/HDDDZ/test/chatApp/data/sql"
)

var groupDBService DB.GroupSQLService = &DB.SQLService{}
var groupRedisService Redis.GroupRedisService = &Redis.RedisService{}

func Create(ownerUid int, groupName string, memberUids ...int) (int64, error) {
	log("group,Create,params,memberUids=", memberUids)
	gid, err := groupDBService.Create(ownerUid, groupName, memberUids...)
	log("group,Create,gid=", gid, "err=", err)
	go synchroniseGroupAndMembersInfo(int(gid))
	return gid, err
}

func TransferOwner(gid int, newOwnerId int, OlderOwnerId int) error {
	err := groupDBService.TransferOwner(gid, newOwnerId, OlderOwnerId)
	log("group,TransferOwner,err=", err)

	if err == nil {
		go synchroniseGroupInfo(gid)
	}
	return err
}

func LeaveGroup(gid int, uid int) error {
	err := groupDBService.LeaveGroup(gid, uid)
	log("group,LeaveGroup,err=", err)

	if err == nil {
		go synchroniseGroupAndMembersInfo(gid)
	}
	return err
}

func UpdateGroupInfo(gid int, groupName string, description string) error {
	err := groupDBService.UpdateGroupInfo(gid, groupName, description)
	log("group,UpdateGroupInfo,err=", err)

	if err == nil {
		go synchroniseGroupInfo(gid)
	}
	return err
}

func GetGroupByGid(gid int) Common.Group {
	group := groupRedisService.GetGroupByGid(fmt.Sprint(gid))
	log("group,GetGroupByGid,redis,group=", group)

	if group == (Common.Group{}) {
		group = groupDBService.GetGroupByGid(gid)
		log("group,GetGroupByGid,mysql,group=", group)
		err := groupRedisService.RefreshGroups(group)
		log("group,GetGroupByGid,redis,err=", err)
	}
	return group
}

func GetAllGroupsByUid(uid int) []Common.Group {
	groups := groupDBService.GetAllGroupsByUid(uid)
	log("group,GetAllGroupsByUid,mysql,groups=", groups)

	if len(groups) != 0 {
		err := groupRedisService.RefreshGroups(groups...)
		log("group,GetAllGroupsByUid,redis,err=", err)

	}
	return groups
}

func AddMember(gid int, uids ...int) error {
	err := groupDBService.AddMember(gid, uids...)
	log("group,AddMember,mysql,err=", err)
	if err == nil {
		go synchroniseGroupMembersInfo(gid)
	}
	return err
}

func RemoveMember(gid int, uids ...int) error {
	err := groupDBService.RemoveMember(gid, uids...)
	log("group,RemoveMember,mysql,err=", err)
	if err == nil {
		go synchroniseGroupMembersInfo(gid)
	}
	return err
}

func GetAllMembersInfo(gid int) []Common.GroupMember {
	members := groupRedisService.GetAllMembers(fmt.Sprint(gid))
	log("group,GetAllMembersInfo,redis,members=", members)

	if len(members) == 0 {
		members = groupDBService.GetAllMembersInfo(gid)
		log("group,GetAllMembersInfo,mysql,members=", members)

		err := groupRedisService.RefreshGroupAllMembers(fmt.Sprint(gid), members...)
		log("group,GetAllMembersInfo,redis,err=", err)

	}
	return members
}

func GetAllMembersUid(gid int) []int {
	var members []int
	membersId := groupRedisService.GetAllMembers(fmt.Sprint(gid))
	log("group,GetAllMembersUid,redis,membersId=", membersId)

	if len(members) == 0 {
		members = groupDBService.GetAllMembersUid(gid)
		log("group,GetAllMembersUid,mysql,members=", members)
	} else {
		members = make([]int, len(membersId))
		for index, id := range membersId {
			members[index] = id.Uid
		}
	}
	return members
}

func GetMemberInfo(gid int, uid int) Common.GroupMember {
	member := groupRedisService.GetMember(fmt.Sprint(gid), fmt.Sprint(uid))
	log("group,GetMemberInfo,redis,member=", member)
	if member == (Common.GroupMember{}) {
		member = groupDBService.GetMemberInfo(gid, uid)
		log("group,GetMemberInfo,mysql,member=", member)
		err := groupRedisService.AddGroupMembers(fmt.Sprint(gid), member)
		log("group,GetMemberInfo,redis,err=", err)

	}
	return member
}

func UpdateMemberInfo(gid int, uid int, alias string, identity Common.MemberIdentity) error {
	err := groupDBService.UpdateMemberInfo(gid, uid, alias, identity)
	log("group,UpdateMemberInfo,mysql,err=", err)
	if err != nil {
		go synchroniseGroupMembersInfo(gid)
	}
	return err
}

func synchroniseGroupAndMembersInfo(gid int) {
	synchroniseGroupInfo(gid)
	synchroniseGroupMembersInfo(gid)
}

func synchroniseGroupMembersInfo(gid int) {
	members := groupDBService.GetAllMembersInfo(gid)
	log("group,synchroniseGroupMembersInfo,mysql,members=", members)
	err := groupRedisService.RefreshGroupAllMembers(fmt.Sprint(gid), members...)
	log("group,synchroniseGroupMembersInfo,redis,err=", err)
}

func synchroniseGroupInfo(gid int) {
	group := groupDBService.GetGroupByGid(gid)
	log("group,synchroniseGroupInfo,mysql,group=", group)
	err := groupRedisService.RefreshGroups(group)
	log("group,synchroniseGroupInfo,redis,err=", err)
}
