package db

const query_User_By_Users = `select users.uid, users.user_name,users.pass_word,
	users_token.token from users INNER JOIN users_token ON users.uid=users_token.uid 
	where %s = ?`

const query_Messages_By_Uid = `select msg_id, sender_id,receiver_id,
conversatio_type,message_body_type,content from messages 
	where sender_id = ? OR receiver_id = ?`

const query_Messages_By_MsgId_Rege = `select msg_id, sender_id,receiver_id,
conversatio_type,message_body_type,content from messages 
	where msg_id REGEXP '%s'`

const query_Request_By_Id = `select id, msg,sender_uid,
receiver_uid,request_state from request_add_friend 
	where id = ?`

const query_Request_By_Uids = `select id, msg,sender_uid,
receiver_uid,request_state from request_add_friend 
	where (sender_uid=? AND receiver_uid=?) OR (sender_uid=? AND receiver_uid=?)`

const query_All_Request_By_Uid = `select id, msg,sender_uid,
receiver_uid,request_state from request_add_friend 
	where sender_uid = ? OR receiver_uid = ?`

const query_All_Friends_By_Uid = `select users.uid, users.user_name from users
INNER JOIN friend_relation ON (users.uid=friend_relation.user_id_1 or users.uid=friend_relation.user_id_2) and users.uid != ? 
where friend_relation.user_id_1 = ? OR friend_relation.user_id_2 = ?`

const query_All_Friends_Uid_By_Uid = `select users.uid from users
INNER JOIN friend_relation ON (users.uid=friend_relation.user_id_1 or users.uid=friend_relation.user_id_2) and users.uid != ?
where friend_relation.user_id_1 = ? OR friend_relation.user_id_2 = ?`
