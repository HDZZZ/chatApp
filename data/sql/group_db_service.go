package db

import (
	"fmt"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

const _table_chat_group = "chat_group"
const _table_group_members = "group_members"
const _field_group_name = "group_name"
const _field_ownerId = "ownerId"
const _field_gid = "gid"
const _field_uid = "uid"
const _field_identity = "identity"

func create(ownerUid int, groupName string, memberUids ...int) (int64, error) {
	id, err := insertRows(_table_chat_group, []string{_field_group_name, _field_ownerId},
		[]any{groupName, ownerUid})
	if err != nil {
		fmt.Println("insert into chat_group error", err)
		return 0, err
	}

	var insertValues = make([][]any, len(memberUids)+1)
	for index, memberUid := range memberUids {
		insertValues[index] = []any{id, memberUid, Common.Member}
	}
	insertValues[len(insertValues)-1] = []any{id,
		ownerUid, Common.Owner}

	_, err = insertRows(_table_group_members, []string{_field_gid,
		_field_uid, _field_identity}, insertValues...)

	if err != nil {
		fmt.Println("insert into group_members error", err)
		return 0, err
	}
	return id, err
}

func transferOwner(gid int, newOwnerId int, OlderOwnerId int) error {

	err := updateRows(_table_chat_group+" JOIN group_members ON chat_group.gid = group_members.gid",
		fmt.Sprintf("chat_group.ownerId = %v , group_members.identity = %v", newOwnerId,
			Common.Owner),
		fmt.Sprintf("group_members.uid=%v AND chat_group.gid = %v", newOwnerId, gid))
	if err != nil {
		fmt.Println("update chat_group owner error", err)
		return err
	}

	err = updateRows(_table_group_members, fmt.Sprintf("identity = %v", Common.Member),
		fmt.Sprintf("gid=%v AND uid=%v", gid, OlderOwnerId))
	if err != nil {
		fmt.Println("update chat_group owner error", err)
		return err
	}
	return err
}

func leaveGroup(gid int, uid int) error {
	err := delRows(_table_group_members, fmt.Sprintf("gid = %v AND uid= %v", gid, uid))
	if err != nil {
		fmt.Println("leave group error", err)
		return err
	}
	return nil
}

func updateGroupInfo(gid int, groupName string, description string) error {
	var selectUpdate string
	if groupName != "" {
		selectUpdate = "group_name = '" + groupName + "'"
	}
	if description != "" {
		selectUpdate = selectUpdate + ", description = '" + description + "'"
	}

	err := updateRows(_table_chat_group, selectUpdate, fmt.Sprintf("gid=%v", gid))
	if err != nil {
		fmt.Println("update chat_group info error", err)
		return err
	}
	return err
}

func getGroupByGid(gid int) Common.Group {
	groups, _ := queryStruct[Common.Group](fmt.Sprintf(query_Group_By_Gid, gid))
	if len(groups) == 0 {
		return Common.Group{}
	}
	return groups[0]
}

func getAllGroupsByUid(uid int) []Common.Group {
	groups, _ := queryStruct[Common.Group](fmt.Sprintf(query_Groups_By_Uid, uid))
	return groups
}

func addMember(gid int, uids ...int) error {

	var insertValues = make([][]any, len(uids))
	for index, memberUid := range uids {
		insertValues[index] = []any{gid, memberUid, Common.Member}
	}

	_, err := insertRows(_table_group_members, []string{_field_gid,
		_field_uid, _field_identity}, insertValues...)

	if err != nil {
		fmt.Println("insert into group_members error", err)
		return err
	}
	return err
}

func removeMember(gid int, uids ...int) error {
	var sqlSentence string
	for index, uid := range uids {
		if index == 0 {
			sqlSentence = sqlSentence + fmt.Sprintf("(gid = %v AND uid= %v)", gid, uid)
		} else {
			sqlSentence = sqlSentence + fmt.Sprintf("(gid = %v AND uid= %v)", gid, uid)
		}
	}
	err := delRows(_table_group_members, sqlSentence)
	if err != nil {
		fmt.Println("DELETE friend error", err)
		return err
	}
	return nil
}

func getAllMembersInfo(gid int) []Common.GroupMember {
	members, _ := queryStruct[Common.GroupMember](fmt.Sprintf(query_GroupMembers_By_Gid, gid))
	return members
}

func getAllMembersUid(gid int) []int {
	var messages = []int{}
	var inputUser int

	queryRows(fmt.Sprintf(query_GroupMembersUid_By_Gid, gid), func() {
		newUser := inputUser
		messages = append(messages, newUser)
	}, &inputUser)
	return messages
}

func getMemberInfo(gid int, uid int) Common.GroupMember {
	members, _ := queryStruct[Common.GroupMember](fmt.Sprintf(query_GroupMember, gid, uid))
	if len(members) == 0 {
		return Common.GroupMember{}
	}
	return members[0]

}

func updateMemberInfo(gid int, uid int, alias string, identity Common.MemberIdentity) error {

	var selectUpdate string
	if alias != "" {
		selectUpdate = "alias = '" + alias + "'"
	}
	if identity != Common.NONE {
		selectUpdate = selectUpdate + ", identity = " + fmt.Sprint(identity)
	}

	err := updateRows(_table_group_members, selectUpdate,
		fmt.Sprintf("gid=%v AND uid=%v", gid, uid))
	if err != nil {
		fmt.Println("update chat_group info error", err)
		return err
	}
	return nil
}
