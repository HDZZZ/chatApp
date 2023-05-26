package db

import (
	"fmt"
	"log"

	Common "github.com/HDDDZ/test/chatApp/data/common"
)

func create(ownerUid int, groupName string, memberUids ...int) (int64, error) {
	fmt.Println("sql,create,memberUids=", memberUids)
	id, err := _exec("INSERT INTO chat_group(group_name,ownerId) VALUES(?,?)",
		groupName, ownerUid)
	if err != nil {
		log.Println("insert into chat_group error", err)
		return 0, err
	}
	var selectSQL string
	for _, memberUid := range memberUids {
		selectSQL = selectSQL + fmt.Sprintf("(%v,%v,%v),", id, memberUid, Common.Member)
	}
	selectSQL = selectSQL + fmt.Sprintf("(%v,%v,%v)", id, ownerUid, Common.Owner)

	syntax := "INSERT INTO group_members(gid,uid,identity) VALUES" + selectSQL
	_, err = _exec(syntax)

	if err != nil {
		log.Println("insert into group_members error", err)
		return 0, err
	}
	return id, err
}

func transferOwner(gid int, newOwnerId int, OlderOwnerId int) error {
	// _, err := _exec("UPDATE chat_group set ownerId = ? where gid=?", newOwnerId, gid)
	// if err != nil {
	// 	log.Println("update chat_group owner error", err)
	// 	return err
	// }
	fmt.Println("gid=", gid)
	fmt.Println("newOwnerId=", newOwnerId)
	fmt.Println("OlderOwnerId=", OlderOwnerId)
	_, err := _exec("UPDATE chat_group JOIN group_members ON chat_group.gid = group_members.gid  set chat_group.ownerId = ? , group_members.identity = ? where group_members.uid=? AND chat_group.gid = ?", newOwnerId, Common.Owner, newOwnerId, gid)
	if err != nil {
		log.Println("update chat_group owner error", err)
		return err
	}

	_, err = _exec("UPDATE group_members set identity = ? where gid=? AND uid=?", Common.Member, gid, OlderOwnerId)
	if err != nil {
		log.Println("update chat_group owner error", err)
		return err
	}
	//todo
	return err
}

func leaveGroup(gid int, uid int) error {
	_, err := _exec(fmt.Sprintf("DELETE FROM group_members WHERE %s", fmt.Sprintf("gid = %v AND uid= %v", gid, uid)))
	if err != nil {
		log.Println("leave group error", err)
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
	sqlSentence := fmt.Sprintf("UPDATE chat_group set %s where gid=?", selectUpdate)
	fmt.Println("sqlSentence=", sqlSentence)
	_, err := _exec(sqlSentence, gid)
	if err != nil {
		log.Println("update chat_group info error", err)
		return err
	}
	return err
}

func getGroupByGid(gid int) Common.Group {
	var messages = []Common.Group{}
	inputUser := Common.Group{}

	_query(query_Group_By_Gid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid}, &inputUser.Gid, &inputUser.GroupName, &inputUser.OwnerId, &inputUser.Description)
	return messages[0]
}

func getAllGroupsByUid(uid int) []Common.Group {
	var messages = []Common.Group{}
	inputUser := Common.Group{}

	_query(query_Groups_By_Uid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{uid}, &inputUser.Gid, &inputUser.GroupName, &inputUser.OwnerId, &inputUser.Description)
	return messages
}

func addMember(gid int, uids ...int) error {

	sqlSentence := "INSERT INTO group_members(gid,uid,identity)"

	for index, uid := range uids {
		if index == 0 {
			sqlSentence = sqlSentence + fmt.Sprintf(" VALUES(%v,%v,%v)", gid, uid, Common.Member)
		} else {
			sqlSentence = sqlSentence + fmt.Sprintf(",(%v,%v,%v)", gid, uid, Common.Member)
		}
	}
	fmt.Println("sqlSentence=", sqlSentence)
	_, err := _exec(sqlSentence)

	if err != nil {
		log.Println("insert into group_members error", err)
		return err
	}
	return err
}

func removeMember(gid int, uids ...int) error {
	sqlSentence := "DELETE FROM group_members WHERE"
	for index, uid := range uids {
		if index == 0 {
			sqlSentence = sqlSentence + fmt.Sprintf("(gid = %v AND uid= %v)", gid, uid)
		} else {
			sqlSentence = sqlSentence + fmt.Sprintf("(gid = %v AND uid= %v)", gid, uid)
		}
	}
	_, err := _exec(sqlSentence)
	if err != nil {
		log.Println("DELETE friend error", err)
		return err
	}
	return nil
}

func getAllMembersInfo(gid int) []Common.GroupMember {
	var messages = []Common.GroupMember{}
	inputUser := Common.GroupMember{}

	_query(query_GroupMembers_By_Gid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid}, &inputUser.Gid, &inputUser.Uid, &inputUser.Alias, &inputUser.Identity)
	return messages
}

func getAllMembersUid(gid int) []int {
	var messages = []int{}
	var inputUser int

	_query(query_GroupMembersUid_By_Gid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid}, &inputUser)
	return messages
}

func getMemberInfo(gid int, uid int) Common.GroupMember {
	var messages = []Common.GroupMember{}
	inputUser := Common.GroupMember{}

	_query(query_GroupMember, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid, uid}, &inputUser.Gid, &inputUser.Uid, &inputUser.Alias, &inputUser.Identity)
	return messages[0]
}

func updateMemberInfo(gid int, uid int, alias string, identity Common.MemberIdentity) error {
	var selectUpdate string
	if alias != "" {
		selectUpdate = "alias = '" + alias + "'"
	}
	if identity != Common.NONE {
		selectUpdate = selectUpdate + ", identity = " + fmt.Sprint(identity)
	}

	_, err := _exec(fmt.Sprintf("UPDATE group_members set %s where gid=? AND uid=?", selectUpdate), gid, uid)
	if err != nil {
		log.Println("update chat_group info error", err)
		return err
	}
	return nil
}
