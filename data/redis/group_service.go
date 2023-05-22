package redis

import (
	"encoding/json"
	"fmt"
	"strings"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

/*
*无法根据用户id获取所有其参与的群聊列表
group_id165161, group_id165161_member:[123150,16511],group_member_gid2156123_uid1056161,
*/
const GROUP_KEY = "group_gid"                           //群表的数据
const GROUP_MEMBER_KEY = "group_member_gid%s_uid"       //群成员表的数据
const GROUP_MEMBER_GROUP_KEY = "group_member_group_gid" //关系表的概念

/*
*
添加或者刷新群聊
*
*/
func refreshGroups(groups ...Common.Group) error {
	requestParis := make(map[string]string, len(groups))
	for _, group := range groups {
		stream, _ := json.Marshal(group)
		requestParis[createReidsrKey(GROUP_KEY, group.Gid)] = string(stream)
	}
	return setPairs(requestParis)
}

func getGroupByGid(gid string) Common.Group {
	group, err := get(createReidsrKey(GROUP_KEY, gid))
	if err != nil {
		return Common.Group{}
	}
	var newUser Common.Group
	json.Unmarshal([]byte(group), &newUser)
	return newUser
}

func refreshGroupAllMembers(gid string, members ...Common.GroupMember) error {
	requestParis := make(map[string]string, len(members)+1)
	var MemberIdsOfGroup string

	for _, group := range members {
		stream, _ := json.Marshal(group)
		requestParis[createReidsrKey(fmt.Sprintf(GROUP_MEMBER_KEY, gid), group.Uid)] = string(stream)
		MemberIdsOfGroup = MemberIdsOfGroup + "|" + fmt.Sprint(group.Uid)
	}
	requestParis[createReidsrKey(GROUP_MEMBER_GROUP_KEY, gid)] = MemberIdsOfGroup
	return setPairs(requestParis)
}

/**
*
 */
func addGroupMembers(gid string, members ...Common.GroupMember) error {
	requestParis := make(map[string]string, len(members))
	var MemberIdsOfGroup string

	for _, group := range members {
		stream, _ := json.Marshal(group)
		requestParis[createReidsrKey(fmt.Sprintf(GROUP_MEMBER_KEY, gid), group.Uid)] = string(stream)
		MemberIdsOfGroup = MemberIdsOfGroup + "|" + fmt.Sprint(group.Uid)
	}
	appendValue(createReidsrKey(GROUP_MEMBER_GROUP_KEY, gid), MemberIdsOfGroup)
	return setPairs(requestParis)
}

/**
*	删除群成员(会删除群成员数据并且清理群映射的所有群成员id)
 */
func removeGroupMembers(gid string, members ...Common.GroupMember) error {
	memberIds := make([]string, len(members))
	groupMemberIds, _ := get(createReidsrKey(GROUP_MEMBER_GROUP_KEY, gid))

	for index, group := range members {
		memberIds[index] = createReidsrKey(fmt.Sprintf(GROUP_MEMBER_KEY, gid), group.Uid)
		groupMemberIds = strings.Replace(groupMemberIds, "|"+fmt.Sprint(group.Uid), "", -1)
	}
	set(createReidsrKey(GROUP_MEMBER_GROUP_KEY, gid), groupMemberIds)
	_, err := delete(memberIds...)
	return err
}

/**
*	更新群成员数据
 */
func updateGroupMemberInfo(members ...Common.GroupMember) error {
	requestParis := make(map[string]string, len(members))
	for _, group := range members {
		stream, _ := json.Marshal(group)
		requestParis[createReidsrKey(fmt.Sprintf(GROUP_MEMBER_KEY, group.Gid), group.Uid)] = string(stream)
	}
	return setPairs(requestParis)
}

func getMember(gid string, uid string) Common.GroupMember {
	stream, err := get(createReidsrKey(fmt.Sprintf(GROUP_MEMBER_KEY, gid), uid))
	if err != nil {
		return Common.GroupMember{}
	}
	var newUser Common.GroupMember
	json.Unmarshal([]byte(stream), &newUser)
	return newUser
}

func getAllMembers(gid string) []Common.GroupMember {
	stream, err := get(createReidsrKey(GROUP_MEMBER_GROUP_KEY, gid))
	if err != nil {
		return []Common.GroupMember{}
	}
	uids := removenullvalue(strings.Split(stream, "|"))
	membersStream, err := mget(uids...)
	members := make([]Common.GroupMember, len(membersStream))

	var newUser Common.GroupMember
	for index, stream := range membersStream {
		json.Unmarshal([]byte(stream.(string)), &newUser)
		members[index] = newUser
	}
	return members
}
