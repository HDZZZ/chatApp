package db

const query_User_By_Users = `select users.uid, users.user_name,users.pass_word,
	users_token.token from users INNER JOIN users_token ON users.uid=users_token.uid 
	where %s = %v`

const query_Messages_By_Uid = `select msg_id, sender_id,receiver_id,
conversatio_type,message_body_type,content from messages 
	where sender_id = %v OR receiver_id = %v`

const query_Messages_By_MsgId_Rege = `select msg_id, sender_id,receiver_id,
conversatio_type,message_body_type,content from messages 
	where msg_id REGEXP '%s'`

const query_Request_By_Id = `select id, msg,sender_uid,
receiver_uid,request_state from request_add_friend 
	where id = %v`

const query_Request_By_Uids = `select id, msg,sender_uid,
receiver_uid,request_state from request_add_friend 
	where (sender_uid=%v AND receiver_uid=%v) OR (sender_uid=%v AND receiver_uid=%v)`

const query_All_Request_By_Uid = `select id, msg,sender_uid,
receiver_uid,request_state from request_add_friend 
	where sender_uid = %v OR receiver_uid = %v`

const query_All_Friends_By_Uid = `select users.uid, users.user_name from users
INNER JOIN friend_relation ON (users.uid=friend_relation.user_id_1 or users.uid=friend_relation.user_id_2) and users.uid != %v 
where friend_relation.user_id_1 = %v OR friend_relation.user_id_2 = %v`

const query_All_Friends_Uid_By_Uid = `select users.uid from users
INNER JOIN friend_relation ON (users.uid=friend_relation.user_id_1 or users.uid=friend_relation.user_id_2) and users.uid != %v
where friend_relation.user_id_1 = %v OR friend_relation.user_id_2 = %v`

const query_Group_By_Gid = `select gid, group_name,ownerId,
 COALESCE(description, '') as description from chat_group 
	where gid = %v`

const query_Groups_By_Uid = `select chat_group.gid, group_name,ownerId,
COALESCE(description, '') as description from chat_group JOIN group_members ON chat_group.gid = group_members.gid
	where group_members.uid = %v`

const query_GroupMembers_By_Gid = `select gid, uid,COALESCE(alias, '') as alias,
identity from group_members where gid = %v`

const query_GroupMembersUid_By_Gid = `select uid from group_members where gid = %v`

const query_GroupMember = `select gid, uid,COALESCE(alias, '') as alias,
identity from group_members where gid = %v AND uid = %v`
