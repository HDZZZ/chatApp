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
