package db

import (
	"fmt"
	"log"
)

func Create(ownerUid int, groupName string, memberUids ...int) (int64, error) {
	id, err := _exec("INSERT INTO chat_group(group_name,ownerId,member_count) VALUES(?,?,?)",
		groupName, ownerUid, len(memberUids)+1)
	if err != nil {
		log.Println("insert into chat_group error", err)
		return 0, err
	}
	selectSQL := fmt.Sprintf("SELECT %v,%v,%v", id, ownerUid, Owner)
	for _, memberUid := range memberUids {
		selectSQL = selectSQL + fmt.Sprintf("UNION ALL SELECT %v,%v,%v", id, memberUid, Member)
	}

	_, err = _exec("INSERT INTO group_members(gid,uid,identity)" + selectSQL)

	if err != nil {
		log.Println("insert into group_members error", err)
		return 0, err
	}
	return id, err
}

func TransferOwner(gid int, newOwnerId int, OlderOwnerId int) error {
	// _, err := _exec("UPDATE chat_group set ownerId = ? where gid=?", newOwnerId, gid)
	// if err != nil {
	// 	log.Fatal("update chat_group owner error", err)
	// 	return err
	// }
	fmt.Println("gid=", gid)
	fmt.Println("newOwnerId=", newOwnerId)
	fmt.Println("OlderOwnerId=", OlderOwnerId)
	_, err := _exec("UPDATE chat_group JOIN group_members ON chat_group.gid = group_members.gid  set chat_group.ownerId = ? , group_members.identity = ? where group_members.uid=? AND chat_group.gid = ?", newOwnerId, Owner, newOwnerId, gid)
	if err != nil {
		log.Fatal("update chat_group owner error", err)
		return err
	}

	_, err = _exec("UPDATE group_members set identity = ? where gid=? AND uid=?", Member, gid, OlderOwnerId)
	if err != nil {
		log.Fatal("update chat_group owner error", err)
		return err
	}
	//todo
	return err
}

func LeaveGroup(gid int, uid int) error {
	_, err := _exec(fmt.Sprintf("DELETE FROM group_members WHERE %s", fmt.Sprintf("gid = %v AND uid= %v", gid, uid)))
	if err != nil {
		log.Println("leave group error", err)
		return err
	}
	_, err = _exec("UPDATE chat_group SET member_count = member_count - 1 WHERE gid = ?", gid)
	if err != nil {
		log.Println("leave group error", err)
		return err
	}
	return nil
}

func UpdateGroupInfo(gid int, groupName string, description string) error {
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
		log.Fatal("update chat_group info error", err)
		return err
	}
	return err
}

func GetGroupByGid(gid int) Group {
	var messages = []Group{}
	inputUser := Group{}

	_query(query_Group_By_Gid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid}, &inputUser.Gid, &inputUser.GroupName, &inputUser.OwnerId, &inputUser.Description, &inputUser.MemberCount)
	return messages[0]
}

func GetAllGroupsByUid(uid int) []Group {
	var messages = []Group{}
	inputUser := Group{}

	_query(query_Groups_By_Uid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{uid}, &inputUser.Gid, &inputUser.GroupName, &inputUser.OwnerId, &inputUser.Description, &inputUser.MemberCount)
	return messages
}

func AddMember(gid int, uids ...int) error {

	sqlSentence := "INSERT INTO group_members(gid,uid,identity)"

	for index, uid := range uids {
		if index == 0 {
			sqlSentence = sqlSentence + fmt.Sprintf(" VALUES(%v,%v,%v)", gid, uid, Member)
		} else {
			sqlSentence = sqlSentence + fmt.Sprintf(",(%v,%v,%v)", gid, uid, Member)
		}
	}
	fmt.Println("sqlSentence=", sqlSentence)
	_, err := _exec(sqlSentence)

	if err != nil {
		log.Println("insert into group_members error", err)
		return err
	}
	_, err = _exec("UPDATE chat_group SET member_count = member_count + ? WHERE gid = ?", len(uids), gid)
	if err != nil {
		log.Println("insert into group_members error", err)
		return err
	}
	return err
}

func RemoveMember(gid int, uids ...int) error {
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
		log.Fatal("DELETE friend error", err)
		return err
	}
	_, err = _exec("UPDATE chat_group SET member_count = member_count - 1 WHERE gid = ?", gid)
	if err != nil {
		log.Fatal("DELETE friend error", err)
		return err
	}
	return nil
}

func GetAllMembersInfo(gid int) []GroupMember {
	var messages = []GroupMember{}
	inputUser := GroupMember{}

	_query(query_GroupMembers_By_Gid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid}, &inputUser.Gid, &inputUser.Uid, &inputUser.Alias, &inputUser.Identity)
	return messages
}

func GetAllMembersUid(gid int) []int {
	var messages = []int{}
	var inputUser int

	_query(query_GroupMembersUid_By_Gid, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid}, &inputUser)
	return messages
}

func GetMemberInfo(gid int, uid int) GroupMember {
	var messages = []GroupMember{}
	inputUser := GroupMember{}

	_query(query_GroupMember, func(a ...any) {
		newUser := inputUser
		messages = append(messages, newUser)
	}, []any{gid, uid}, &inputUser.Gid, &inputUser.Uid, &inputUser.Alias, &inputUser.Identity)
	return messages[0]
}

func UpdateMemberInfo(gid int, uid int, alias string, identity MemberIdentity) error {
	var selectUpdate string
	if alias != "" {
		selectUpdate = "alias = '" + alias + "'"
	}
	if identity != NONE {
		selectUpdate = selectUpdate + ", identity = " + fmt.Sprint(identity)
	}

	_, err := _exec(fmt.Sprintf("UPDATE group_members set %s where gid=? AND uid=?", selectUpdate), gid, uid)
	if err != nil {
		log.Fatal("update chat_group info error", err)
		return err
	}
	return nil
}
