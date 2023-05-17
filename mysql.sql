CREATE TABLE users(
uid int NOT NULL AUTO_INCREMENT,
user_name VARCHAR(32),
pass_word VARCHAR(32),
PRIMARY KEY (uid)
)
ALTER TABLE users AUTO_INCREMENT=100000;

CREATE TABLE users_token (
    token VARCHAR(64),
    uid int,
	CONSTRAINT FK_users_token FOREIGN KEY (uid)
    REFERENCES users(uid)
);

CREATE TABLE messages(
msg_id int NOT NULL AUTO_INCREMENT,
sender_id int,
receiver_id int,
conversatio_type TINYINT,
message_body_type TINYINT,
content VARCHAR(1024),
PRIMARY KEY (msg_id)
);

ALTER TABLE messages AUTO_INCREMENT=10000000;

create table friend_relation(
id int primary key auto_increment,
user_id_1 int not null,
user_id_2 int not null,unique(user_id_1,user_id_2),
foreign key(user_id_1) references users(uid) on delete cascade on update cascade,
foreign key(user_id_2) references users(uid) on delete cascade on update cascade
);

CREATE TABLE request_add_friend (  id int primary key auto_increment,     msg VARCHAR(64),     sender_uid int,     receiver_uid int,  request_state int DEFAULT 0, CONSTRAINT FK_request_add_friend_send FOREIGN KEY (sender_uid)     REFERENCES users(uid) on delete cascade on update cascade,     CONSTRAINT FK_request_add_friend_receive FOREIGN KEY (receiver_uid) REFERENCES users(uid) on delete cascade on update cascade )