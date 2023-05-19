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

CREATE TABLE request_add_friend (  id int primary key auto_increment,     msg VARCHAR(64),     sender_uid int,     receiver_uid int,  request_state int DEFAULT 0, CONSTRAINT FK_request_add_friend_send FOREIGN KEY (sender_uid)     REFERENCES users(uid) on delete cascade on update cascade,     CONSTRAINT FK_request_add_friend_receive FOREIGN KEY (receiver_uid) REFERENCES users(uid) on delete cascade on update cascade );

CREATE TABLE chat_group (
	gid int NOT NULL AUTO_INCREMENT,
    group_name VARCHAR(64),
    ownerId int,
    member_count int,
    description VARCHAR(256),
    foreign key(ownerId) references users(uid),
    PRIMARY KEY (gid)
);
ALTER TABLE chat_group AUTO_INCREMENT=100000000;

CREATE TABLE group_members (
	gid int,
    uid int,
	alias VARCHAR(64),
    identity int,
    foreign key(gid) references chat_group(gid),
    foreign key(uid) references users(uid),
    PRIMARY KEY (gid,uid)
);